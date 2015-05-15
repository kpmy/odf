package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kpmy/ypk/assert"
	"github.com/mitchellh/mapstructure"
	"log"
	_ "odf/model/stub" //don't forget pimpl
	"sync"
)

type Msg struct {
	Typ string
}

type Handler func(m *Msg)

var wg *sync.WaitGroup = &sync.WaitGroup{}
var busChan chan *Msg

//этот хэндлер только пишет сообщения в канал главной горутины
func busHandler(m *Msg) {
	busChan <- m
}

//этот хэндлер обрабатывает сообщения в рамках главной горутины
func handle(m *Msg) {
	switch m.Typ {
	case "init":
		panic("not implemented")
	}
}

func Process(m *Msg) {
	assert.For(m != nil, 20)
	js.Global.Call("postMessage", m)
}

func Init(handler Handler) {
	js.Global.Set("onmessage", func(oEvent *js.Object) {
		data := oEvent.Get("data").Interface()
		m := &Msg{}
		err := mapstructure.Decode(data, m)
		assert.For(err == nil, 40)
		handler(m)
	})
}

func main() {
	log.Println("odf loading... ")
	Init(busHandler)
	busChan = make(chan *Msg)
	wg.Add(1)
	go func(wg *sync.WaitGroup, c chan *Msg) {
		log.Println("done")
		Process(&Msg{Typ: "init"})
		for {
			select {
			case m := <-c:
				handle(m)
			}
		}
	}(wg, busChan)
	wg.Wait()
	log.Println("odf closed")
}
