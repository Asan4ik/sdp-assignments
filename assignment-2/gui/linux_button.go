package gui

// LinuxButton is a concrete product that belongs to the Linux family.
type LinuxButton struct{}

func (LinuxButton LinuxButton) Render() string {
	return "Rendering a Linux-style button"
}
