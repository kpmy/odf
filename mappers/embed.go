package mappers

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
	"odf/xmlns"
	"odf/xmlns/draw"
	"odf/xmlns/svg"
	"odf/xmlns/text"
	"odf/xmlns/xlink"
	"strconv"
	"time"
)

type Draw struct {
	rd   io.Reader
	mime xmlns.Mime
}

func NewDraw(source io.Reader, _mime xmlns.Mime) *Draw {
	return &Draw{rd: source, mime: _mime}
}

func nextUrl() (ret string) {
	t := time.Now().UnixNano()
	seed := rand.Int63()
	h := md5.New()
	binary.Write(h, binary.LittleEndian, t)
	binary.Write(h, binary.LittleEndian, seed)
	data := h.Sum(nil)
	for _, x := range data {
		ret = ret + strconv.FormatInt(int64(x), 16)
	}
	return
}

//generators.Embeddable
func (d *Draw) Reader() io.Reader {
	return d.rd
}

func (d *Draw) MimeType() xmlns.Mime {
	return d.mime
}

//запись изображения
func (d *Draw) WriteTo(fm *Formatter, name string, w, h interface{}) string {
	fm.makePara()
	wr := fm.m.NewWriter(fm.rider)
	wr.WritePos(New(draw.Frame))
	wr.Attr(draw.Name, name).Attr(text.AnchorType, text.Paragraph).Attr(svg.Width, w).Attr(svg.Height, h)
	wr.WritePos(New(draw.Image))
	url := "Pictures/" + nextUrl()
	wr.Attr(xlink.Href, url).Attr(xlink.Type, xlink.Simple).Attr(xlink.Show, xlink.Embed).Attr(xlink.Actuate, xlink.OnLoad)
	return url
}
