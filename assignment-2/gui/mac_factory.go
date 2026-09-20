package gui

/* MacFactory is the Mac counterpart of WindowsFactory. Same idea,
   different family of products. */
type MacFactory struct{}

func (macFactory MacFactory) CreateButton() Button {
	return MacButton{}
}

func (macFactory MacFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}
