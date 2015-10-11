package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
	_ "github.com/kpmy/odf/model/stub" //don't forget pimpl
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"io"
	"log"
	"sync"
)

type Msg struct {
	Type  string
	Param string
	Data  string
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
	switch m.Type {
	case "init":
		log.Println("message bus connected")
	case "get":
		var rd io.Reader
		if m.Param == "demo" {
			rd, _ = demo()
		} else if m.Param == "report" {
			rd, _ = report()
		}
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, rd)
		m := &Msg{Type: "data"}
		m.Data = base64.StdEncoding.EncodeToString(buf.Bytes())
		Process(m)
	default:
		halt.As(100, "not implemented", m.Type)
	}
}

func Process(m *Msg) {
	assert.For(m != nil, 20)
	s, _ := json.Marshal(m)
	js.Global.Call("postMessage", string(s))
}

func Init(handler Handler) {
	js.Global.Set("onmessage", func(oEvent *js.Object) {
		_data := oEvent.Get("data").Interface().(string)
		log.Println(_data)
		m := &Msg{}
		err := json.Unmarshal([]byte(_data), m)
		log.Println(m)
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
		Process(&Msg{Type: "init"})
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
