# ODF
[Open Document Format](http://docs.oasis-open.org/office/v1.2/OpenDocument-v1.2.html) (ODF) producer library for Go (golang).

## Описание
Формирование документа в формате Open Document (ODF) для языка Go (golang).

Формирование документа заключается в последовательном вызове инструкций Форматтера, который выполняет модификацию одной или нескольких частей модели документа.
Затем вызывается процедура генерации файла-архива .odf

Клиентский код изолируется от особенностей структуры документа ODF. 

Необходимость форматтера обсуловлена тем, что стандарт ODF предполагает изменение видимого содержимого документа посредством изменений в нескольких местах модели документа (стили, встроенные файлы, и т.д.)

## Пример
    go get github.com/kpmy/odf
В пакете demo есть пример использования ODF для формирования отчета.

## Description
This library is for generation of ODF document with Go.

You can produce a document with content by calling the Formatter methods.
Then you can save this document to zip-file .odf

No need for your code to handle with ODF XML content. 
More examples in demo/report.go

## Example

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

## Moar

It works in browser now. Got Demo3, GopherJS + Dart.

[http://kpmy.github.io/odf/](http://kpmy.github.io/odf/)