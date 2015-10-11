package main

import (
	"github.com/kpmy/odf/generators"
	"github.com/kpmy/odf/mappers"
	"github.com/kpmy/odf/model"
	_ "github.com/kpmy/odf/model/stub" //don't forget pimpl
	"github.com/kpmy/odf/xmlns"
	"os"
)

func main() {
	if output, err := os.Create("demo2.odf"); err == nil {
		//cleanup
		defer output.Close()
		//we need an empty model
		m := model.ModelFactory()
		//standard formatter
		fm := &mappers.Formatter{}
		//couple them
		fm.ConnectTo(m)
		//we want text
		fm.MimeType = xmlns.MimeText
		//yes we can
		fm.Init()
		//pretty simple
		fm.WriteString("Hello, World!")
		//store file
		generators.GeneratePackage(m, nil, output, fm.MimeType)
	}
}
