package gui

// WindowsButton is a concrete product that belongs to the Windows family.
type WindowsButton struct{}

func (windowsButton WindowsButton) Render() string {
	return "Rendering a Windows-style button"
}
