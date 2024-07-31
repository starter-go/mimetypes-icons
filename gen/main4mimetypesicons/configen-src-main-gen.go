package main4mimetypesicons
import (
    p0ef6f2938 "github.com/starter-go/application"
    pd1a916a20 "github.com/starter-go/libgin"
    p85a4d026d "github.com/starter-go/mimetypes"
    p477725ed1 "github.com/starter-go/mimetypes-icons/app/classes/icons"
    p50e8d11cf "github.com/starter-go/mimetypes-icons/app/web/controllers"
     "github.com/starter-go/application"
)

// type p477725ed1.ConfigProviderImpl in package:github.com/starter-go/mimetypes-icons/app/classes/icons
//
// id:com-477725ed1f3ba215-icons-ConfigProviderImpl
// class:
// alias:alias-477725ed1f3ba2156c23126bc5507aca-ConfigProvider
// scope:singleton
//
type p477725ed1f_icons_ConfigProviderImpl struct {
}

func (inst* p477725ed1f_icons_ConfigProviderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-477725ed1f3ba215-icons-ConfigProviderImpl"
	r.Classes = ""
	r.Aliases = "alias-477725ed1f3ba2156c23126bc5507aca-ConfigProvider"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p477725ed1f_icons_ConfigProviderImpl) new() any {
    return &p477725ed1.ConfigProviderImpl{}
}

func (inst* p477725ed1f_icons_ConfigProviderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p477725ed1.ConfigProviderImpl)
	nop(ie, com)

	
    com.WebPathPrefix = inst.getWebPathPrefix(ie)
    com.ResPathPrefix = inst.getResPathPrefix(ie)
    com.RegPriority = inst.getRegPriority(ie)


    return nil
}


func (inst*p477725ed1f_icons_ConfigProviderImpl) getWebPathPrefix(ie application.InjectionExt)string{
    return ie.GetString("${mimetypes.icons.web-path-prefix}")
}


func (inst*p477725ed1f_icons_ConfigProviderImpl) getResPathPrefix(ie application.InjectionExt)string{
    return ie.GetString("${mimetypes.icons.res-path-prefix}")
}


func (inst*p477725ed1f_icons_ConfigProviderImpl) getRegPriority(ie application.InjectionExt)int{
    return ie.GetInt("${mimetypes.icons.registration-priority}")
}



// type p477725ed1.Loader in package:github.com/starter-go/mimetypes-icons/app/classes/icons
//
// id:com-477725ed1f3ba215-icons-Loader
// class:class-85a4d026daf77828ef49edb2adfd695e-Registry
// alias:alias-477725ed1f3ba2156c23126bc5507aca-Service
// scope:singleton
//
type p477725ed1f_icons_Loader struct {
}

func (inst* p477725ed1f_icons_Loader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-477725ed1f3ba215-icons-Loader"
	r.Classes = "class-85a4d026daf77828ef49edb2adfd695e-Registry"
	r.Aliases = "alias-477725ed1f3ba2156c23126bc5507aca-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p477725ed1f_icons_Loader) new() any {
    return &p477725ed1.Loader{}
}

func (inst* p477725ed1f_icons_Loader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p477725ed1.Loader)
	nop(ie, com)

	
    com.ConfigProvider = inst.getConfigProvider(ie)
    com.AppContext = inst.getAppContext(ie)


    return nil
}


func (inst*p477725ed1f_icons_Loader) getConfigProvider(ie application.InjectionExt)p477725ed1.ConfigProvider{
    return ie.GetComponent("#alias-477725ed1f3ba2156c23126bc5507aca-ConfigProvider").(p477725ed1.ConfigProvider)
}


func (inst*p477725ed1f_icons_Loader) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type p50e8d11cf.IconsController in package:github.com/starter-go/mimetypes-icons/app/web/controllers
//
// id:com-50e8d11cfb5c6e27-controllers-IconsController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p50e8d11cfb_controllers_IconsController struct {
}

func (inst* p50e8d11cfb_controllers_IconsController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-50e8d11cfb5c6e27-controllers-IconsController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p50e8d11cfb_controllers_IconsController) new() any {
    return &p50e8d11cf.IconsController{}
}

func (inst* p50e8d11cfb_controllers_IconsController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p50e8d11cf.IconsController)
	nop(ie, com)

	
    com.Conf = inst.getConf(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p50e8d11cfb_controllers_IconsController) getConf(ie application.InjectionExt)p477725ed1.ConfigProvider{
    return ie.GetComponent("#alias-477725ed1f3ba2156c23126bc5507aca-ConfigProvider").(p477725ed1.ConfigProvider)
}


func (inst*p50e8d11cfb_controllers_IconsController) getService(ie application.InjectionExt)p477725ed1.Service{
    return ie.GetComponent("#alias-477725ed1f3ba2156c23126bc5507aca-Service").(p477725ed1.Service)
}



// type p50e8d11cf.TypesController in package:github.com/starter-go/mimetypes-icons/app/web/controllers
//
// id:com-50e8d11cfb5c6e27-controllers-TypesController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p50e8d11cfb_controllers_TypesController struct {
}

func (inst* p50e8d11cfb_controllers_TypesController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-50e8d11cfb5c6e27-controllers-TypesController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p50e8d11cfb_controllers_TypesController) new() any {
    return &p50e8d11cf.TypesController{}
}

func (inst* p50e8d11cfb_controllers_TypesController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p50e8d11cf.TypesController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p50e8d11cfb_controllers_TypesController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p50e8d11cfb_controllers_TypesController) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}


