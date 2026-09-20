package gui

// MacButton is a concrete product that belongs to the Mac family.
type MacButton struct{}

func (macButton MacButton) Render() string {
	return "Rendering a Mac-style button"
}
