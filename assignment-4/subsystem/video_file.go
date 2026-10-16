/* Package subsystem has the parts that do the real video work
   They dont know about the facade package */
package subsystem

import (
	"fmt"
	"path/filepath"
	"strings"
)

// VideoFile is the input file. It only knows its name and the codec it was encoded with
type VideoFile struct {
	Name      string
	CodecType string
}

// NewVideoFile makes a VideoFile and finds its codec type from the extension
func NewVideoFile(name string) (*VideoFile, error) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(name), "."))
	if ext == "" {
		return nil, fmt.Errorf("file %q has no extension", name)
	}

	codecType := ""
	switch ext {
	case "mp4":
		codecType = "mpeg4"
	case "ogg":
		codecType = "ogg"
	default:
		return nil, fmt.Errorf("unknown file type %q", ext)
	}

	return &VideoFile{Name: name, CodecType: codecType}, nil
}
