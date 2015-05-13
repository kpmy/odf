package main

import (
	"odf/generators"
	"odf/mappers"
	"odf/model"
	_ "odf/model/stub" //don't forget pimpl
	"odf/xmlns"
	"os"
)

func main() {
	if output, err := os.Create("demo2.odf"); err == nil {
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
		//cleanup
		defer output.Close()
	}
}
