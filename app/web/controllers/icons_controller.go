package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/mimetypes-icons/app/classes/icons"
)

// IconsController ...
type IconsController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Conf    icons.ConfigProvider //starter:inject("#")
	Service icons.Service        //starter:inject("#")

}

func (inst *IconsController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *IconsController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *IconsController) route(rp libgin.RouterProxy) error {

	cfg := inst.Conf.Configuration()
	path1 := cfg.WebPathPrefix

	rp.GET(path1+"/*name", func(ctx *gin.Context) {
		inst.handle(ctx, cfg)
	})
	return nil
}

func (inst *IconsController) handle(c *gin.Context, cfg *icons.Configuration) {
	name := c.Param("name")
	code := http.StatusOK
	prefix := cfg.WebPathPrefix
	img, err := inst.Service.FindImage(c, prefix+"/"+name)
	if err != nil {
		msg := "HTTP 404"
		img = &icons.Image{
			Type: "text/plain",
			Data: []byte(msg),
		}
		code = http.StatusNotFound
	}
	c.Data(code, img.Type.String(), img.Data)
}

////////////////////////////////////////////////////////////////////////////////
