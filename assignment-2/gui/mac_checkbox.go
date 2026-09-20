package gui

// MacCheckbox is a concrete product that belongs to the Mac family.
type MacCheckbox struct{}

func (macCheckbox MacCheckbox) Render() string {
	return "Rendering a Mac-style checkbox"
}
