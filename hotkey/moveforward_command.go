package hotkey

type MoveForwardCommand struct {
	Command
	tank *Tank
}

func NewMoveForwardCommand(tank *Tank) Command {
	return &MoveForwardCommand{
		tank: tank,
	}
}

func (moveforwardCmd *MoveForwardCommand) Execute() {
	moveforwardCmd.tank.MoveForward()
}

func (moveforwardCmd *MoveForwardCommand) Undo() {
	moveforwardCmd.tank.MoveBackward()
}

func (moveforwardCmd *MoveForwardCommand) GetName() string {
	return "MoveTankForward"
}
