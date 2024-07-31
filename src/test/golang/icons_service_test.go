package main

import (
	"testing"

	"github.com/starter-go/mimetypes-icons/modules/mimetypesicons"
	"github.com/starter-go/mimetypes-icons/src/test/golang/unit"
	"github.com/starter-go/units"
)

func TestIconsService(t *testing.T) {

	args := []string{"app_name", "stop"}
	mod := mimetypesicons.ModuleForTest()
	units.Run(&units.Config{
		Args:   args,
		Cases:  unit.UnitNameIconsService,
		Module: mod,
		T:      t,
	})
}
