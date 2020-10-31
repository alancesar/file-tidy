package mime

import (
	"github.com/gabriel-vasile/mimetype"
	"strings"
)

type Type string

const (
	AudioType Type = "audio/"
	ImageType Type = "image/"
)

// Is checks if a file is a provide Type.
func Is(path string, t Type) bool {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return false
	}

	return strings.Contains(mime.String(), string(t))
}
