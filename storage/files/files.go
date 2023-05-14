package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"tbot/lib/reqerr"
	"tbot/storage"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

var ErrNoSavesPage = errors.New("No saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = reqerr.Wrap("cant save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.Mkdir(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)

	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)

	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = reqerr.WrapIfErr("canrt pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavesPage
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fName, err := fileName(p)
	if err != nil {
		return reqerr.Wrap("cant remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("cant remove file %s", path)
		return reqerr.Wrap(msg, err)
	}

	return nil
}

func (s Storage) IsExist(p *storage.Page) (bool, error) {
	fName, err := fileName(p)
	if err != nil {
		return false, reqerr.Wrap("cant check if file exist", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("cant check if file %s exist", path)
		return false, reqerr.Wrap(msg, err)
	}

	return true, nil
}

func (s Storage) decodePage(fPath string) (*storage.Page, error) {
	file, err := os.Open(fPath)
	if err != nil {
		return nil, reqerr.Wrap("cant read file", err)
	}

	defer func() {
		_ = file.Close()
	}()

	var p storage.Page

	if err := gob.NewDecoder(file).Decode(&p); err != nil {
		return nil, reqerr.Wrap("cant decode page", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
