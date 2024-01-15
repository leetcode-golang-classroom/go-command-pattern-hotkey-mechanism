package hotkey

type DisconnectCommand struct {
	Command
	telecom *Telecom
}

func NewDisconnectCommand(telecom *Telecom) Command {
	return &DisconnectCommand{
		telecom: telecom,
	}
}

func (disconnectCmd *DisconnectCommand) Execute() {
	disconnectCmd.telecom.Disconnect()
}

func (disconnectCmd *DisconnectCommand) Undo() {
	disconnectCmd.telecom.Connect()
}
func (disconnectCmd *DisconnectCommand) GetName() string {
	return "DisconnectTelecom"
}
