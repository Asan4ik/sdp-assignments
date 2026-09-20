package gui

// WindowsFactory is a concrete factory. It only ever produces Windows
// components, so a client using it can never end up with a mismatched
// set of widgets.
type WindowsFactory struct{}

func (windowsFactory WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (windowsFactory WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}
