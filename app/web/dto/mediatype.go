package dto

import (
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes-icons/app/web/dxo"
)

// MediaType ...
type MediaType struct {
	ID dxo.MediaTypeID `json:"id"`

	Base

	Type        mimetypes.Type `json:"type"`
	Label       string         `json:"label"`
	Description string         `json:"description"`
	Icon        string         `json:"icon"`

	// Suffixes    []string
}
