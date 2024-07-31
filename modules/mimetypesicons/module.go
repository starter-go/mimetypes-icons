package mimetypesicons

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/mimetypes-common/modules/mimetypescommon"
	mimetypesicons "github.com/starter-go/mimetypes-icons"
	"github.com/starter-go/mimetypes-icons/gen/main4mimetypesicons"
	"github.com/starter-go/mimetypes-icons/gen/test4mimetypesicons"
	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/units/modules/units"
)

// Module  ...
func Module() application.Module {
	mb := mimetypesicons.NewMainModule()
	mb.Components(main4mimetypesicons.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(mimetypes.Module())
	mb.Depend(mimetypescommon.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := mimetypesicons.NewTestModule()
	mb.Components(test4mimetypesicons.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}
