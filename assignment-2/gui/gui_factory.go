package gui

/* GUIFactory is the abstract factory. It declares one creation method
   per product in the family, so any concrete factory is guaranteed to
   produce a matching, consistent set of components. */
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}
