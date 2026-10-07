package style

const (
	FrameWidth    = 120
	FrameHeight   = 36
	ContentHeight = FrameHeight - HeaderHeight - 1 // FrameHeight - HeaderHeight - FooterHeight
)

const (
	// The horizontal padding on each side of the content area.
	ContentPadX = 2
	// The horizontal gap between the two panels.
	ColumnGap = 2
)

// Header and Footer spacing
const (
	HeaderHeight = 2
	FooterHeight = 1 // Not in use at the moment
)

// Projects spacing
// 37 + ColumnGap(2) + 77 = 116 = FrameWidth - 2*ContentPadX.
const (
	ProjectListWidth   = 37
	ProjectDetailWidth = 77
)
