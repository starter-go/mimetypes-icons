package webunits

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/mimetypes"
	mimetypescommon "github.com/starter-go/mimetypes-common"
	"github.com/starter-go/vlog"
)

// TestImageListController ...
type TestImageListController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	TM mimetypes.Manager //starter:inject("#")
}

func (inst *TestImageListController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *TestImageListController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *TestImageListController) route(rp libgin.RouterProxy) error {
	rp = rp.For("test")
	// rp.GET("", inst.handleGetJSON)
	rp.GET("/test/mediatype-image-list", inst.handleGetPage)
	return nil
}

func (inst *TestImageListController) handleGetPage(c *gin.Context) {

	builder := &bytes.Buffer{}
	inst.makePageData(builder)

	data := builder.Bytes()
	code := http.StatusOK
	mediatype := "text/html; charset=UTF-8"
	c.Data(code, mediatype, data)
}

func (inst *TestImageListController) makePageData(b *bytes.Buffer) {

	title := "MediaType Image List"
	imageStyle := "width:24px; height:24px;"

	all := mimetypescommon.GetRequiredTypeNameList()

	b.WriteString("<!DOCTYPE html>")
	b.WriteString("<html>")
	b.WriteString("<head>")
	b.WriteString("  <meta content='text/html; charset=utf-8' />")
	b.WriteString("</head>")
	b.WriteString("<body>")
	b.WriteString("  <h1>" + title + "</h1>")
	b.WriteString("  <ol>")

	for _, t := range all {

		info := inst.findTypeInfo(t)

		b.WriteString("<li>")

		b.WriteString("<div>" + t.String() + "</div>")
		b.WriteString("<div>" + info.Label + "</div>")
		b.WriteString("<div>" + info.Description + "</div>")
		b.WriteString(fmt.Sprintf("<img style='%s' src='%s' />", imageStyle, info.Icon))

		b.WriteString("</li>")
	}

	b.WriteString("</ol>")
	b.WriteString("</body>")
	b.WriteString("</html>")
}

func (inst *TestImageListController) findTypeInfo(t mimetypes.Type) *mimetypes.Info {

	info, err := inst.TM.Find(t, nil)
	if err == nil {
		return info
	}

	info = &mimetypes.Info{
		Icon: "/image/bad/icon",
	}
	vlog.Warn(err.Error())
	return info
}
