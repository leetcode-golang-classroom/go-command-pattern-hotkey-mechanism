package hotkey

type MoveBackwardCommand struct {
	Command
	tank *Tank
}

func NewMoveBackwardCommand(tank *Tank) Command {
	return &MoveBackwardCommand{
		tank: tank,
	}
}

func (movebackwardCmd *MoveBackwardCommand) Execute() {
	movebackwardCmd.tank.MoveBackward()
}
func (movebackwardCmd *MoveBackwardCommand) Undo() {
	movebackwardCmd.tank.MoveForward()
}
func (movebackwardCmd *MoveBackwardCommand) GetName() string {
	return "MoveTankBackward"
}
