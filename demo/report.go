package main

import (
	"github.com/kpmy/ypk/assert"
	"odf/generators"
	"odf/mappers"
	"odf/model"
	_ "odf/model/stub" // необходимо
	"odf/xmlns"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func do(wg *sync.WaitGroup, suffix string) {
	assert.For(suffix != "", 20)
	output, _ := os.OpenFile("test-"+suffix+".odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	generators.Generate(m, nil, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(4)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go do(wg, strconv.Itoa(i))
		wg.Add(1)
	}
	wg.Wait()
}
