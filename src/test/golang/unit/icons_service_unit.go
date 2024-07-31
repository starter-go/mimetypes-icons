package unit

import (
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes-icons/app/classes/icons"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// IconsServiceUnit ... 单元测试示例
type IconsServiceUnit struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	Icons icons.Service     //starter:inject("#")
	TM    mimetypes.Manager //starter:inject("#")

}

func (inst *IconsServiceUnit) _impl() units.Units { return inst }

// Units ...
func (inst *IconsServiceUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     UnitNameIconsService,
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

// Units ...
func (inst *IconsServiceUnit) test1() error {

	x, err := inst.TM.Find("text/html", nil)
	if err != nil {
		return err
	}
	vlog.Debug("find type: %s", x.Type.Pure())
	return nil
}
