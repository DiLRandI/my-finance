package pdf

import (
	"fmt"
	"io"
	"os"
)

type Pdf interface {
	NumPage() int
	io.ReadCloser
}

type pdf struct {
	f        io.ReadCloser
	path     string
	numPages int
}

func Open(path string) (Pdf, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF file %s: %w", path, err)
	}

	return &pdf{
		path: path,
		f:    f,
	}, nil
}

func (p *pdf) NumPage() int {
	return p.numPages
}

func (p *pdf) Read(data []byte) (n int, err error) {

	return 0, nil
}

func (p *pdf) Close() error {
	return p.f.Close()
}
