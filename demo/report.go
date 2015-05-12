package main

import (
	"bytes"
	"github.com/kpmy/golorem"
	"github.com/kpmy/ypk/assert"
	"io"
	"math/rand"
	"odf/generators"
	"odf/mappers"
	"odf/mappers/attr"
	"odf/model"
	_ "odf/model/stub" // необходимо
	"odf/xmlns"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func report(suffix string, fm *mappers.Formatter) {
	{ //first page
		fm.WriteString("\n\n\n\n\n\n\n\n\n\n")
		fm.SetAttr(new(attr.TextAttributes).Size(32).Bold()).SetAttr(new(attr.ParagraphAttributes).AlignCenter())
		fm.WritePara("Periodic report")
		fm.SetAttr(nil).SetAttr(new(attr.TextAttributes).Italic())
		fm.WritePara("Report:\t")
		fm.SetAttr(new(attr.TextAttributes).Bold())
		fm.WriteString(suffix + "\n")
		fm.SetAttr(new(attr.TextAttributes).Italic())
		fm.WriteString("Date:\t")
		fm.SetAttr(new(attr.TextAttributes).Bold())
		fm.WriteString(time.Now().String())
		fm.SetAttr(new(attr.ParagraphAttributes).PageBreak())
		fm.WritePara("")
		fm.SetAttr(nil)
	}
	para := func() {
		fm.SetAttr(new(attr.TextAttributes).Bold().Size(18))
		fm.WritePara(strings.ToUpper(lorem.Word(5, 15)))
		fm.WriteLn()
		fm.SetAttr(nil)
		para := lorem.Paragraph(5, 10)
		fm.WritePara("\t" + para + "\n\n")
	}
	for i := 0; i < 5; i++ {
		para()
	}
	{ //huge table
		fm.SetAttr(new(attr.TextAttributes).Bold().Size(18))
		fm.WritePara("TABLE 50x5")
		fm.SetAttr(nil)
		tm := &mappers.TableMapper{}
		tm.ConnectTo(fm)
		tm.Write("test", 50+1, 5) //50+header row
		tt := tm.List["test"]
		tm.Span(tt, 0, 0, 1, 5)
		fm.SetAttr(new(attr.ParagraphAttributes).AlignCenter()).SetAttr(new(attr.TextAttributes).Bold())
		tm.Pos(tt, 0, 0).WriteString("Header")
		fm.SetAttr(nil)
		for i := 1; i < 51; i++ {
			for j := 0; j < 5; j++ {
				if j == 0 {
					fm.SetAttr(new(attr.TextAttributes).Bold())
				} else {
					fm.SetAttr(nil)
				}
				tm.Pos(tt, i, j).WriteString(strconv.Itoa(i * j))
			}
		}
	}
	{ //appendix
		fm.RegisterFont("Courier New", "Courier New") // may not work in Linux/MacOS
		fm.SetAttr(nil).SetAttr(new(attr.ParagraphAttributes).PageBreak())
		fm.SetAttr(new(attr.TextAttributes).Size(18).Bold())
		fm.WritePara("Appendix A.\nListing of report.go")
		fm.WriteLn()
		fm.SetAttr(nil).SetAttr(new(attr.TextAttributes).FontFace("Courier New").Size(6))
		if f, err := os.Open("report.go"); err == nil {
			defer f.Close()
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, f)
			fm.WritePara(string(buf.Bytes()))
		} else {
			fm.WritePara("File not found.")
		}
	}
}

func do(wg *sync.WaitGroup) {
	src := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	suffix := strconv.Itoa(src.Int())
	output, _ := os.OpenFile("test-"+suffix+".odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	embed := make(map[string]generators.Embeddable)
	{
		const ImagePng xmlns.Mime = "image/png"
		img, _ := os.Open("10414104.png")
		d := mappers.NewDraw(img, ImagePng)
		fm.SetAttr(new(attr.ParagraphAttributes).AlignRight())
		url := d.WriteTo(fm, "Two Gophers", 4.0, 4.0) //magic? real size of `.png` in cm
		embed[url] = d
	}
	fm.SetAttr(new(attr.TextAttributes).Bold())
	fm.WriteString("\nSo Strange inc.")
	fm.WriteString("\n" + lorem.Email())
	fm.WriteString("\n" + lorem.Url())
	fm.SetAttr(nil)
	report(suffix, fm)
	generators.GeneratePackage(m, embed, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(4)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go do(wg)
		wg.Add(1)
	}
	wg.Wait()
}
