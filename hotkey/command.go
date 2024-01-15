package hotkey

type Command interface {
	Execute()
	Undo()
	GetName() string
}
