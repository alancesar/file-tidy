package path

import (
	"bytes"
	"github.com/alancesar/tidy-file/mime"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	defaultSeparator = "/"
)

type Builder struct {
	separator string
}

func NewBuilder() *Builder {
	return &Builder{
		separator: defaultSeparator,
	}
}

// LookFor deep walks in a path and get all files that match with a provided mime.Type.
func LookFor(rootPath string, types ...mime.Type) []string {
	paths := make([]string, 0)
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		for index := range types {
			if mime.Is(path, types[index]) {
				paths = append(paths, path)
			}
		}

		return nil
	})

	return paths
}

// FromPattern creates a path from an interface{} based on a pattern using text/template engine.
func (b *Builder) FromPattern(pattern string, source interface{}) (string, error) {
	parsed, err := template.New("path").Parse(pattern)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := parsed.Execute(buf, source); err != nil {
		return "", err
	}

	elements := strings.Split(buf.String(), b.separator)
	return filepath.Join(elements...), nil
}
