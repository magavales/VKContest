package dialog

type State struct {
	Type  Type
	Name  Name
	Value string

	PrevState *State
}
