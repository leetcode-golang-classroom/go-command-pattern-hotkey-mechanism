package hotkey

type ConnectCommand struct {
	Command
	telecom *Telecom
}

func NewConnectCommand(telecom *Telecom) Command {
	return &ConnectCommand{
		telecom: telecom,
	}
}

func (connectCmd *ConnectCommand) Execute() {
	connectCmd.telecom.Connect()
}
func (connectCmd *ConnectCommand) Undo() {
	connectCmd.telecom.Disconnect()
}

func (connectCmd *ConnectCommand) GetName() string {
	return "ConnectTelecom"
}
