package test4mimetypesicons
import (
    p85a4d026d "github.com/starter-go/mimetypes"
    p477725ed1 "github.com/starter-go/mimetypes-icons/app/classes/icons"
    p5117e698c "github.com/starter-go/mimetypes-icons/src/test/golang/unit"
    p51a2e1bfe "github.com/starter-go/mimetypes-icons/src/test/golang/webunits"
     "github.com/starter-go/application"
)

// type p5117e698c.DemoUnit in package:github.com/starter-go/mimetypes-icons/src/test/golang/unit
//
// id:com-5117e698cd678564-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p5117e698cd_unit_DemoUnit struct {
}

func (inst* p5117e698cd_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5117e698cd678564-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5117e698cd_unit_DemoUnit) new() any {
    return &p5117e698c.DemoUnit{}
}

func (inst* p5117e698cd_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p5117e698c.DemoUnit)
	nop(ie, com)

	


    return nil
}



// type p5117e698c.IconsServiceUnit in package:github.com/starter-go/mimetypes-icons/src/test/golang/unit
//
// id:com-5117e698cd678564-unit-IconsServiceUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p5117e698cd_unit_IconsServiceUnit struct {
}

func (inst* p5117e698cd_unit_IconsServiceUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5117e698cd678564-unit-IconsServiceUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5117e698cd_unit_IconsServiceUnit) new() any {
    return &p5117e698c.IconsServiceUnit{}
}

func (inst* p5117e698cd_unit_IconsServiceUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p5117e698c.IconsServiceUnit)
	nop(ie, com)

	
    com.Icons = inst.getIcons(ie)
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p5117e698cd_unit_IconsServiceUnit) getIcons(ie application.InjectionExt)p477725ed1.Service{
    return ie.GetComponent("#alias-477725ed1f3ba2156c23126bc5507aca-Service").(p477725ed1.Service)
}


func (inst*p5117e698cd_unit_IconsServiceUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type p51a2e1bfe.TestImageListController in package:github.com/starter-go/mimetypes-icons/src/test/golang/webunits
//
// id:com-51a2e1bfed0b0eb1-webunits-TestImageListController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p51a2e1bfed_webunits_TestImageListController struct {
}

func (inst* p51a2e1bfed_webunits_TestImageListController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-51a2e1bfed0b0eb1-webunits-TestImageListController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p51a2e1bfed_webunits_TestImageListController) new() any {
    return &p51a2e1bfe.TestImageListController{}
}

func (inst* p51a2e1bfed_webunits_TestImageListController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p51a2e1bfe.TestImageListController)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p51a2e1bfed_webunits_TestImageListController) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}


