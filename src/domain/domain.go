package domain

import (
	"io"
)

type Domain interface {
	FromXML(reader io.Reader) error
	FromJson(reader io.Reader) error
	ToXML() ([]byte, error)
	ToJson() ([]byte, error)
}
