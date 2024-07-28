package carbon

// Component is an interface that represents a view that can be updated and rendered.
type Component interface {
	ComponentUpdate(*Msg, *Cmd) Component
	View() string
}
