package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"tgBot/lib/e"
)

var ErrNoSavedPages = errors.New("no saved pages")

type Storage interface {
	Save(p *Page) error
	Remove(p *Page) error
	PickRandom(userName string) (*Page, error)
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cant calculate hash", err)
	}

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cant calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
