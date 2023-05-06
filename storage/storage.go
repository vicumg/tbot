package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"tbot/lib/reqerr"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", reqerr.Wrap("cant calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", reqerr.Wrap("cant generate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
