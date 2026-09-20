package gui

/* Button is an abstract product. Every platform provides its own
   version of a button that knows how to render itself. */
type Button interface {
	Render() string
}
