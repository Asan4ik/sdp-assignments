package gui

// WindowsCheckbox is a concrete product that belongs to the Windows family.
type WindowsCheckbox struct{}

func (windowsCheckbox WindowsCheckbox) Render() string {
	return "Rendering a Windows-style checkbox"
}
