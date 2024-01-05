package connection

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

type Cable struct {
	src            *Plug
	dst            *Plug
	sX, sY, eX, eY float32
	color          color.RGBA
	path           vector.Path
	vertices       []ebiten.Vertex
	indices        []uint16
	changed        bool
}

func NewCable(src *Plug) *Cable {
	return &Cable{
		src:   src,
		color: CycleColor(),
	}
}

func (c *Cable) SetSrc(src *Plug) {
	c.src = src
}

func (c *Cable) SetDst(dst *Plug) {
	if dst != nil && dst.GetDirection() != c.src.GetDirection() {
		c.dst = dst
		return
	}
	c.dst = nil
}

func (c *Cable) GetSrc() *Plug {
	return c.src
}

func (c *Cable) GetDst() *Plug {
	return c.dst
}

func (c *Cable) Draw(img *ebiten.Image, shiftX, shiftY float64) {
	c.computeEndStartPoints()
	c.computeVertices(shiftX, shiftY)

	opBis := &ebiten.DrawTrianglesOptions{}
	opBis.AntiAlias = true
	opBis.FillRule = ebiten.FillAll
	opBis.ColorScaleMode = ebiten.ColorScaleModePremultipliedAlpha
	img.DrawTriangles(c.vertices, c.indices, whiteImage, opBis)
}

func (c *Cable) getCenteredPosition(p *Plug) (float64, float64) {
	x, y := p.GetLayout().GetAbsolutePosition()
	w, h := p.GetLayout().GetSize()

	return x + w/2, y + h/2
}

func (c *Cable) computeEndStartPoints() {
	x0, y0 := c.getCenteredPosition(c.src)

	x1, y1 := float64(0), float64(0)
	if c.dst != nil {
		x1, y1 = c.getCenteredPosition(c.dst)
	} else {
		mx, my := ebiten.CursorPosition()
		x1, y1 = float64(mx), float64(my)
	}

	sX, sY, eX, eY := float32(x0), float32(y0), float32(x1), float32(y1)
	if sX != c.sX || sY != c.sY || eX != c.eX || eY != c.eY {
		c.sX, c.sY, c.eX, c.eY = sX, sY, eX, eY
		c.changed = true
	}
}

func (c *Cable) computeVertices(shiftX, shiftY float64) {
	if !c.changed {
		return
	}

	c.changed = false

	var sX, sY, eX, eY float32
	if c.sY < c.eY {
		sX, sY, eX, eY = c.eX, c.eY, c.sX, c.sY
	} else {
		sX, sY, eX, eY = c.sX, c.sY, c.eX, c.eY
	}

	s2X, s2Y := sX+(eX-sX)/2, sY+float32(math.Abs(float64(eX-sX)))

	c.path = vector.Path{}
	c.path.MoveTo(sX, sY)
	c.path.QuadTo(s2X, s2Y, eX, eY)

	op := &vector.StrokeOptions{}
	op.Width = 5
	c.vertices, c.indices = c.path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	for i := range c.vertices {
		c.vertices[i].DstX = c.vertices[i].DstX - float32(shiftX)
		c.vertices[i].DstY = c.vertices[i].DstY - float32(shiftY)
		c.vertices[i].ColorR = float32(c.color.R) / float32(0xff)
		c.vertices[i].ColorG = float32(c.color.G) / float32(0xff)
		c.vertices[i].ColorB = float32(c.color.B) / float32(0xff)
		c.vertices[i].ColorA = .5
	}
}

var whiteImage *ebiten.Image

func init() {
	whiteImage = ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)
}
