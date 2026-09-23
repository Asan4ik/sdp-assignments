package gui

// LinuxCheckbox is a concrete product that belongs to the Linux family.
type LinuxCheckbox struct{}

func (LinuxCheckbox LinuxCheckbox) Render() string {
	return "Rendering a Linux-style checkbox"
}
