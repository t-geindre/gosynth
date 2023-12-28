package layout

type IBlock interface {
	// Tree

	GetChildren() []IBlock
	GetParent() IBlock

	// Spacing

	GetMargin() (int, int, int, int)
	SetMargin(margin int)
	SetMarginTop(margin int)
	SetMarginBottom(margin int)
	SetMarginLeft(margin int)
	SetMarginRight(margin int)

	GetPadding() (int, int, int, int)
	SetPadding(padding int)
	SetPaddingTop(padding int)
	SetPaddingBottom(padding int)
	SetPaddingLeft(padding int)
	SetPaddingRight(padding int)

	// Placement

	GetPosition() (int, int)
	SetPosition(x, y int)
	SetPositionX(x int)
	SetPositionY(y int)
	MoveBy(x, y int)
	MoveByX(x int)
	MoveByY(y int)

	HorizontallyCentered(centered bool)
	VerticallyCentered(centered bool)

	// Sizing
	GetSize() (int, int)
	SetSize(width, height int)
	SetWidth(width int)
	SetHeight(height int)

	HorizontallyFluid(fluid bool) // Will have the H size of contained components
	VerticallyFluid(fluid bool)   // Will have the V size of contained components
	HorizontallyFill(percent int)
	VerticallyFill(percent int)

	// Computing

	ComputeLayout()
}
