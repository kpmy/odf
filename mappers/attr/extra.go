package attr

import (
	"fmt"
	"github.com/kpmy/ypk/assert"
	"image/color"
	"strconv"
)

type Border struct {
	Width float64
	Color color.Color
	Style string
}

func (bb Border) String() string {
	assert.For(bb.Style != "", 20)
	assert.For(bb.Color != nil, 21)
	r, g, b, _ := bb.Color.RGBA()
	return fmt.Sprint(strconv.FormatFloat(bb.Width, 'f', 8, 64)+"cm", " ", bb.Style, " ", fmt.Sprintf("#%02X%02X%02X", uint8(r), uint8(g), uint8(b)))
}
