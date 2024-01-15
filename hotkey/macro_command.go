package hotkey

type MacroCommand struct {
	Command
	commands []Command
}

func NewMacroCommand(commands []Command) Command {
	return &MacroCommand{
		commands: commands,
	}
}
func (macroCmd *MacroCommand) Execute() {
	for _, command := range macroCmd.commands {
		command.Execute()
	}
}
func (macroCmd *MacroCommand) Undo() {
	lastIdx := len(macroCmd.commands) - 1
	for commandIdx := lastIdx; commandIdx >= 0; commandIdx-- {
		macroCmd.commands[commandIdx].Undo()
	}
}
func (macroCmd *MacroCommand) GetName() string {
	commandName := ""
	for idx, command := range macroCmd.commands {
		commandName += command.GetName()
		if idx != len(macroCmd.commands)-1 {
			commandName += " & "
		}
	}
	return commandName
}
