package files

import (
	"os"
	"path/filepath"
	"tbot/lib/reqerr"
	"tbot/storage"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

const defaultPerm = 0774

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = reqerr.Wrap("cant save", err) }()
	filepath := filepath.Join(s.basePath, page.UserName)

	if err := os.Mkdir(filepath, defaultPerm); err != nil {
		return err
	}

}
