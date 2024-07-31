package vo

import "github.com/starter-go/mimetypes-icons/app/web/dto"

// MediaTypes ...
type MediaTypes struct {
	Base

	Items []*dto.MediaType `json:"mediatypes"`
}
