package layout

import (
	"math"
)

func computeHorizontal(l ILayout) {
	// No children to handle
	if len(l.GetChildren()) == 0 {
		return
	}

	// Filter out absolute children
	children := make([]ILayout, 0, len(l.GetChildren()))

	for _, c := range l.GetChildren() {
		if !c.GetAbsolutePositioning() {
			children = append(children, c)
		}
	}

	// No children to handle left
	if len(children) == 0 {
		return
	}

	// Store orientation for further usage
	orientVertical := l.GetContentOrientation() == Vertical

	// Compute required size
	contentWidth := float64(0)
	for _, c := range children {
		cww, cwh := c.GetWantedSize().Get()
		cmt, cmb, cml, cmr := c.GetMargin().Get()

		if orientVertical {
			contentWidth += cwh + cmt + cmb
		} else {
			contentWidth += cww + cml + cmr
		}
	}

	// Calculate scaling factor and free space
	lw, lh := l.GetSize().Get()
	lpt, lpb, lpl, lpr := l.GetPadding().Get()

	innerWidth := lw - lpl - lpr
	innerHeight := lh - lpt - lpb

	if orientVertical {
		innerWidth, innerHeight = innerHeight, innerWidth
	}

	scaleFact := innerWidth / contentWidth
	freeSpace := math.Max(0, innerWidth-contentWidth)

	// If there is free space, we need to distribute it across the filler children
	if freeSpace > 0 {
		totalFill := float64(0)
		for _, c := range children {
			if c.GetFill() > 0 {
				totalFill += c.GetFill()
			}
		}

		if totalFill > 100 {
			panic("Total fill cannot be greater than 100%")
		}

		contentWidth += freeSpace * totalFill / 100
		scaleFact = innerWidth / contentWidth
	}

	// Place children and apply scale
	xOffset := lpl
	yOffset := lpt
	for _, c := range children {
		cww, cwh := c.GetWantedSize().Get()
		cmt, cmb, cml, cmr := c.GetMargin().Get()

		_, _ = cww, cwh
		_, _ = cmr, cmb

		fill := float64(0)
		if c.GetFill() > 0 && freeSpace > 0 {
			fill = c.GetFill() / 100 * freeSpace
		}

		if orientVertical {
			c.GetPosition().Set(xOffset+cml, yOffset+cmt*scaleFact)
			ch := (cwh + fill) * scaleFact
			c.GetSize().Set(innerHeight-cml-cmr, ch)
			yOffset += ch + (cmt+cmb)*scaleFact
		} else {
			c.GetPosition().Set(xOffset+cml*scaleFact, yOffset+cmt)
			cw := (cww + fill) * scaleFact
			c.GetSize().Set(cw, innerHeight-cmt-cmb)
			xOffset += cw + c.GetMargin().GetHorizontal()*scaleFact
		}
	}
}
