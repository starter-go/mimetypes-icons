package main

import (
	"os"

	"github.com/starter-go/mimetypes-icons/modules/mimetypesicons"
	"github.com/starter-go/starter"
)

func main() {
	m := mimetypesicons.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
