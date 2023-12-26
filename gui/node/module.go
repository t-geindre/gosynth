package node

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"gosynth/gui/fonts"
	"image/color"
)

type Module struct {
	Node
	//Layout                 *ebiten.Image
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewModule(width, height int) *Module {
	m := &Module{}
	m.Node = *NewNode(width, height, m)

	// Fill background
	bgColor := color.RGBA{R: 230, G: 230, B: 230, A: 255}
	m.Image.Fill(bgColor)

	// Draw border
	borderColor := color.RGBA{R: 172, G: 172, B: 172, A: 255}
	vector.StrokeRect(m.Image, 0, 0, float32(width), float32(height), 1, borderColor, false)

	// Load font
	ft, err := sfnt.Parse(fonts.LemonMilkMedium)
	if err != nil {
		panic(err)
	}

	fe, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size:    15,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	// Add centered title
	textColor := color.RGBA{R: 3, G: 3, B: 3, A: 255}
	w := font.MeasureString(fe, "VCA")
	text.Draw(m.Image, "VCA", fe, (width-w.Round())/2, 25, textColor)

	// Add volume slider / rectangle with rounded corners
	sliderHeight, sliderWidth := float32(height-200), float32(width-20)
	vector.DrawFilledRect(m.Image, 10, 40, sliderWidth, sliderHeight, color.Black, true)

	// Add 25 tick marks
	markOffColor := color.RGBA{R: 80, G: 80, B: 80, A: 255}
	markOnColor := color.RGBA{R: 0, G: 61, B: 194, A: 255}
	markHeight := (sliderHeight - 10 - (25 * 5)) / 25
	fmt.Println(markHeight)
	for i := float32(0); i < 25; i++ {
		if i > 10 {
			vector.DrawFilledRect(m.Image, 15, float32(i*(markHeight+5))+47.5, sliderWidth-10, markHeight, markOnColor, true)
			continue
		}
		vector.DrawFilledRect(m.Image, 15, float32(i*(markHeight+5))+47.5, sliderWidth-10, markHeight, markOffColor, true)
	}

	return m
}

func (m *Module) Update() error {
	if m.MouseLDown {
		x, y := ebiten.CursorPosition()
		m.MoveBy(x-m.LastMouseX, y-m.LastMouseY)
		m.LastMouseX, m.LastMouseY = x, y
	}

	return m.Node.Update()
}
func (m *Module) MouseLeftDown(target INode) {
	if m.GetParent() != nil {
		m.GetParent().MoveFront(m)
	}
	if m == target {
		m.MouseLDown = true
		m.LastMouseX, m.LastMouseY = ebiten.CursorPosition()
		m.Options.ColorScale.ScaleAlpha(0.8)
	}
	m.Node.MouseLeftDown(target)
}

func (m *Module) MouseLeftUp(target INode) {
	if m == target {
		m.Options.ColorScale.Reset()
		m.MouseLDown = false
	}
	m.Node.MouseLeftUp(target)
}
