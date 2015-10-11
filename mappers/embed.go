package mappers

import (
	"crypto/md5"
	"encoding/binary"
	"github.com/kpmy/odf/xmlns"
	"github.com/kpmy/odf/xmlns/draw"
	"github.com/kpmy/odf/xmlns/svg"
	"github.com/kpmy/odf/xmlns/text"
	"github.com/kpmy/odf/xmlns/xlink"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//Draw holds image data and it's mimetype
type Draw struct {
	rd   io.Reader
	mime xmlns.Mime
}

//NewDraw constructor
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

//Reader implements generators.Embeddable
func (d *Draw) Reader() io.Reader {
	return d.rd
}

func (d *Draw) MimeType() xmlns.Mime {
	return d.mime
}

//WriteTo writes image (link to image) to document model, don't forget to pass Draw as Embeddable to package generator
func (d *Draw) WriteTo(fm *Formatter, name string, w, h interface{}) string {
	fm.defaultParaMapper.makePara()
	wr := fm.m.NewWriter()
	wr.Pos(fm.defaultParaMapper.rider.Pos())
	wr.WritePos(New(draw.Frame))
	wr.Attr(draw.Name, name).Attr(text.AnchorType, text.Paragraph).Attr(svg.Width, w).Attr(svg.Height, h)
	wr.WritePos(New(draw.Image))
	url := "Pictures/" + nextUrl()
	wr.Attr(xlink.Href, url).Attr(xlink.Type, xlink.Simple).Attr(xlink.Show, xlink.Embed).Attr(xlink.Actuate, xlink.OnLoad)
	return url
}
