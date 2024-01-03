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
		if orientVertical {
			contentWidth += c.GetWantedSize().GetHeight() + c.GetMargin().GetVertical()
		} else {
			contentWidth += c.GetWantedSize().GetWidth() + c.GetMargin().GetHorizontal()
		}
	}

	// Calculate scaling factor and free space
	innerWidth := l.GetSize().GetWidth() - l.GetPadding().GetHorizontal()
	innerHeight := l.GetSize().GetHeight() - l.GetPadding().GetVertical()

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
	xOffset := l.GetPadding().GetLeft()
	yOffset := l.GetPadding().GetTop()
	for _, c := range children {
		fill := float64(0)
		if c.GetFill() > 0 && freeSpace > 0 {
			fill = c.GetFill() / 100 * freeSpace
		}

		if orientVertical {
			c.GetPosition().Set(
				xOffset+c.GetMargin().GetLeft(),
				yOffset+c.GetMargin().GetTop()*scaleFact,
			)
			c.GetSize().Set(
				innerHeight-c.GetMargin().GetHorizontal(),
				c.GetWantedSize().GetHeight()*scaleFact+fill*scaleFact,
			)
			yOffset += c.GetSize().GetHeight() + c.GetMargin().GetVertical()*scaleFact
		} else {
			c.GetPosition().Set(
				xOffset+c.GetMargin().GetLeft()*scaleFact,
				yOffset+c.GetMargin().GetTop(),
			)
			c.GetSize().Set(
				c.GetWantedSize().GetWidth()*scaleFact+fill*scaleFact,
				innerHeight-c.GetMargin().GetVertical(),
			)
			xOffset += c.GetSize().GetWidth() + c.GetMargin().GetHorizontal()*scaleFact
		}
	}
}
