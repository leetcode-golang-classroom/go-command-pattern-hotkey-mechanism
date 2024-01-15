package hotkey

type ResetCommand struct {
	Command
	keyBoard      *KeyBoard
	prevHotKey    map[string]int
	prevMacroCmds map[int]Command
}

func NewResetCommand(keyBoard *KeyBoard) Command {
	return &ResetCommand{
		keyBoard:      keyBoard,
		prevHotKey:    make(map[string]int),
		prevMacroCmds: make(map[int]Command),
	}
}
func (resetCmd *ResetCommand) Execute() {
	resetCmd.prevHotKey = resetCmd.keyBoard.hotkeysMap
	resetCmd.prevMacroCmds = resetCmd.keyBoard.macroCommands
	resetCmd.keyBoard.ResetHotKey()
	resetCmd.keyBoard.ResetMacroCommand()
}
func (resetCmd *ResetCommand) Undo() {
	resetCmd.keyBoard.hotkeysMap, resetCmd.prevHotKey = resetCmd.prevHotKey, resetCmd.keyBoard.hotkeysMap
	resetCmd.keyBoard.macroCommands, resetCmd.prevMacroCmds = resetCmd.prevMacroCmds, resetCmd.keyBoard.macroCommands
}
func (resetCmd *ResetCommand) GetName() string {
	return "ResetMainControlKeyboard"
}
