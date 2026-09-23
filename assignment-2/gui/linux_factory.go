package gui

/* LinuxFactory is the Linux counterpart of WindowsFactory. Same idea,
   different family of products. */
type LinuxFactory struct{}

func (linuxFactory LinuxFactory) CreateButton() Button {
	return LinuxButton{}
}

func (linuxFactory LinuxFactory) CreateCheckbox() Checkbox {
	return LinuxCheckbox{}
}
