package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes-icons/app/web/dto"
	"github.com/starter-go/mimetypes-icons/app/web/vo"
)

// TypesController ...
type TypesController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder  //starter:inject("#")
	TM     mimetypes.Manager //starter:inject("#")

}

func (inst *TypesController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *TypesController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *TypesController) route(rp libgin.RouterProxy) error {

	rp = rp.For("mediatypes")

	rp.GET("", inst.handleWithQuery)
	rp.GET(":type/:subtype", inst.handleWithType)
	rp.GET("suffix/:suffix", inst.handleWithSuffix)

	return nil
}

func (inst *TypesController) handle(c *gin.Context) {
	req := &myTypesRequest{
		context:    c,
		controller: inst,
	}
	req.execute(req.doExample)
}

func (inst *TypesController) handleWithType(c *gin.Context) {
	req := &myTypesRequest{
		context:              c,
		controller:           inst,
		wantRequestParamType: true,
	}
	req.execute(req.doQuery)
}

func (inst *TypesController) handleWithSuffix(c *gin.Context) {
	req := &myTypesRequest{
		context:                c,
		controller:             inst,
		wantRequestParamSuffix: true,
	}
	req.execute(req.doQuery)
}

func (inst *TypesController) handleWithQuery(c *gin.Context) {
	req := &myTypesRequest{
		context:          c,
		controller:       inst,
		wantRequestQuery: true,
	}
	req.execute(req.doQuery)
}

////////////////////////////////////////////////////////////////////////////////

type myTypesRequest struct {
	context    *gin.Context
	controller *TypesController

	wantRequestBody        bool
	wantRequestParamSuffix bool
	wantRequestParamType   bool
	wantRequestQuery       bool

	suffix   string
	typename string

	body1 vo.MediaTypes
	body2 vo.MediaTypes
}

func (inst *myTypesRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestParamSuffix {
		inst.suffix = c.Param("suffix")
	}

	if inst.wantRequestParamType {
		t1 := c.Param("type")
		t2 := c.Param("subtype")
		inst.typename = t1 + "/" + t2
	}

	if inst.wantRequestQuery {
		inst.suffix = c.Query("suffix")
		inst.typename = c.Query("type")
	}

	return nil
}

func (inst *myTypesRequest) send(err error) {
	d := &inst.body2
	s := d.Status
	res := &libgin.Response{
		Context: inst.context,
		Data:    d,
		Error:   err,
		Status:  s,
	}
	inst.controller.Sender.Send(res)
}

func (inst *myTypesRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myTypesRequest) doExample() error {
	return nil
}

func (inst *myTypesRequest) doQuery() error {

	tm := inst.controller.TM
	suffix := inst.suffix
	tn := inst.typename

	if suffix != "" {
		info, err := tm.FindBySuffix(suffix, nil)
		return inst.makeResult(info, err)
	}

	if tn != "" {
		info, err := tm.Find(mimetypes.Type(tn), nil)
		return inst.makeResult(info, err)
	}

	return nil
}

func (inst *myTypesRequest) makeResult(info *mimetypes.Info, err error) error {

	if err != nil {
		return err
	}

	dst := &dto.MediaType{
		Description: info.Description,
		Icon:        info.Icon,
		Label:       info.Label,
		Type:        info.Type,
	}

	inst.body2.Items = []*dto.MediaType{dst}
	return nil
}
