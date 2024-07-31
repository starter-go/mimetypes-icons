package icons

import (
	"context"

	"github.com/starter-go/mimetypes"
)

// Image ...
type Image struct {
	Name string
	Path string
	Type mimetypes.Type
	Data []byte
}

// Service ...
type Service interface {
	FindImage(c context.Context, path string) (*Image, error)
}
