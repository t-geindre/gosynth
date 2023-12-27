package node

type AppendOption uint8

type AppendOptions struct {
	CenterHorizontally bool
	FillVertically     int
	FillHorizontally   int
	PaddingTop         int
	PaddingBottom      int
	PaddingLeft        int
	PaddingRight       int
	MarginTop          int
	MarginBottom       int
	MarginLeft         int
	MarginRight        int
}

func NewAppendOptions() *AppendOptions {
	return &AppendOptions{}
}

func (a *AppendOptions) Paddings(top, bottom, left, right int) *AppendOptions {
	a.PaddingTop = top
	a.PaddingBottom = bottom
	a.PaddingLeft = left
	a.PaddingRight = right
	return a
}

func (a *AppendOptions) Padding(padding int) *AppendOptions {
	a.PaddingTop = padding
	a.PaddingBottom = padding
	a.PaddingLeft = padding
	a.PaddingRight = padding
	return a
}

func (a *AppendOptions) Margins(top, bottom, left, right int) *AppendOptions {
	a.MarginTop = top
	a.MarginBottom = bottom
	a.MarginLeft = left
	a.MarginRight = right
	return a
}

func (a *AppendOptions) Margin(margin int) *AppendOptions {
	a.MarginTop = margin
	a.MarginBottom = margin
	a.MarginLeft = margin
	a.MarginRight = margin
	return a
}

func (a *AppendOptions) HorizontallyCentered() *AppendOptions {
	a.CenterHorizontally = true
	return a
}

func (a *AppendOptions) VerticallyFill(fill int) *AppendOptions {
	a.FillVertically = fill
	return a
}

func (a *AppendOptions) HorizontallyFill(fill int) *AppendOptions {
	a.FillHorizontally = fill
	return a
}

func ComputeLayout(parent INode) {
	if len(parent.GetChildren()) == 0 {
		return
	}

	baseYOffset := 0
	baseXOffset := 0

	aWidth := parent.GetWidth()
	aHeight := parent.GetHeight()

	pOptions := parent.GetAppendOptions()
	if pOptions != nil {
		baseYOffset += pOptions.PaddingTop
		baseXOffset += pOptions.PaddingLeft
		aWidth -= pOptions.PaddingLeft + pOptions.PaddingRight
		aHeight -= pOptions.PaddingTop + pOptions.PaddingBottom
	}

	for _, HandleVFill := range [2]bool{false, true} {
		yOffset := baseYOffset
		xOffset := baseXOffset
		for _, child := range parent.GetChildren() {
			child.SetPosition(xOffset, yOffset)

			cOpt := child.GetAppendOptions()
			if cOpt != nil {
				child.MoveBy(cOpt.MarginLeft, cOpt.MarginTop)

				if cOpt.CenterHorizontally {
					child.SetPositionX(xOffset + (aWidth-child.GetWidth())/2)
				}

				if cOpt.FillHorizontally > 0 {
					w := aWidth*cOpt.FillHorizontally/100 - cOpt.MarginLeft - cOpt.MarginRight
					child.SetWidth(w)
				}

				if cOpt.FillVertically > 0 {
					if HandleVFill {
						h := aHeight*cOpt.FillVertically/100 - cOpt.MarginTop - cOpt.MarginBottom
						child.SetHeight(h)
					}
				} else {
					if !HandleVFill {
						aHeight -= child.GetOuterHeight()
					}
				}
			} else {
				if !HandleVFill {
					aHeight -= child.GetOuterHeight()
				}
			}
			yOffset += child.GetOuterHeight()
		}
	}
}
