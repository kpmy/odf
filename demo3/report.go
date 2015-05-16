package main

import (
	"bytes"
	"encoding/base64"
	"github.com/kpmy/golorem"
	"image/color"
	"io"
	"math/rand"
	"odf/generators"
	"odf/mappers"
	"odf/mappers/attr"
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/fo"
	"strconv"
	"strings"
	"time"
)

func r(suffix string, fm *mappers.Formatter) {
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
		fm.SetAttr(new(attr.TableCellAttributes).Border(attr.Border{Width: 0.01, Color: color.Black, Style: fo.Solid}))
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
		fm.WritePara("File not found because no IO allowed in browser.")
	}
}

func report() (io.Reader, error) {
	src := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	suffix := strconv.Itoa(src.Int())
	output := bytes.NewBuffer(nil)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	embed := make(map[string]generators.Embeddable)
	{
		const ImagePng xmlns.Mime = "image/png"
		if data, err := base64.StdEncoding.DecodeString(imgData); err == nil {
			img := bytes.NewBuffer(data)
			d := mappers.NewDraw(img, ImagePng)
			fm.SetAttr(new(attr.ParagraphAttributes).AlignRight())
			url := d.WriteTo(fm, "Two Gophers", 4.0, 4.0) //magic? real size of `.png` in cm
			embed[url] = d
		}
	}
	fm.SetAttr(new(attr.TextAttributes).Bold())
	fm.WriteString("\nSo Strange inc.")
	fm.WriteString("\n" + lorem.Email())
	fm.WriteString("\n" + lorem.Url())
	fm.SetAttr(nil)
	r(suffix, fm)
	generators.GeneratePackage(m, embed, output, fm.MimeType)
	return output, nil
}

var imgData = `iVBORw0KGgoAAAANSUhEUgAAAaQAAAGkCAYAAAB+TFE1AAAABmJLR0QA/wD/AP+gvaeTAAAACXBI
WXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH3wULEQENShOIkAAABSxJREFUeNrt18EJAjEURVEj01E6
STuzTjvpJDU914IbEfUr51QwvPlwSUtyAYBvu5oAAEECAEECQJAAQJAAECQAECQABAkABAkAQQIA
QQJAkABAkAAQJAAQJAAECQAECQBBAgBBAkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQA
ECQABAkABAkABAmAIg4T1Df2jBXgdaufzQpeSAAgSAAIEgAIEgCCBACCBIAgAYAgASBIACBIAAgS
AAgSAIIEAIIEgCABgCABIEgAIEgACBIACBIAggQAggSAIAGAIAEgSAAgSAAgSAAIEgAIEgCCBACC
BIAgAYAgASBIACBIAAgSAAgSAIIEAIIEgCABgCABIEgAIEgACBIACBIAggQAggSAIAGAIAEgSAAg
SAAIEgAIEgCCBACCBIAgAYAgASBIACBIACBIAAgSAAgSAIIEAIIEgCABgCABIEgAIEgACBIAPKEl
KfMxY8/4JQCftfrZvJAAQJAAECQAECQABAkABAkAQQIAQQJAkABAkAAQJAAQJAAECQAECQBBAgBB
AkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkAQQIAQQIAQQJAkABAkAAQ
JAAQJAAECQAECQBBAgBBAkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkA
QQIAQQJAkABAkAAQJAAQJAAECQAECQBBAgBBAgBBAkCQAECQABAkABAkAOo6TPAbVj+bFe6NPWMF
9+JevJAAQJAAECQAECQABAkABAkAQQIAQQJAkABAkAAQJAAQJAAECQAECQAECQBBAgBBAkCQAECQ
ABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkAQQIAQQJAkABAkAAQJAAQJAAECQAE
CQBBAgBBAkCQAECQABAkABAkAAQJAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkAQQIAQQJA
kABAkAAQJAAQJAAECQAECQBBAgBBAkCQAECQABAkAHijlsQKAHghAYAgASBIACBIAAgSAAgSAIIE
AIIEgCABgCABIEgAIEgACBIACBIAggQAggSAIAGAIAEgSAAgSAAIEgAIEgCCBACCBIAgAYAgASBI
ACBIAAgSAAgSAIIEAIIEgCCZAABBAgBBAkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQA
ECQABAkABAkAQQIAQQJAkABAkAAQJAAQJAAECQAECQBBAgBBAkCQAECQABAkABAkAAQJAAQJAAQJ
AEECAEECoJij0seMPeOXPLb62azgXtyLe/nne/FCAqAEQQJAkABAkAAQJAAQJAAECQAECQBBAgBB
AkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkAQQIAQQJAkABAkAAQJAAQ
JAAECQAECQBBAgBBAgBBAkCQAECQABAkABAkAAQJAAQJAEECAEECQJAAQJAAECQAECQABAkABAkA
QQIAQQJAkABAkAAQJAAQJAAECQAECQBBAgBBAkCQAECQABAkABAkAAQJAAQJAEECAEECAEECQJAA
QJAAKKglsQIAXkgAIEgACBIACBIAggQAggSAIAGAIAEgSAAgSAAIEgAIEgCCBACCBIAgAYAgASBI
ACBIAAgSAAgSAIIEAIIEgCABgCABIEgAIEgACBIACBIAggQAggSAIAGAIAEgSCYAoIIb4ogrP3v0
vosAAAAASUVORK5CYII=`
