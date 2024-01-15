package hotkey

import (
	"bufio"
	"fmt"
	"strings"
)

type MainController struct {
	ioWriter                *bufio.Writer
	commandHistoryStack     *Stack
	undoCommandHistoryStack *Stack
	commands                []Command
	originCmdLength         int
	keyboard                *KeyBoard
}

func NewMainController(commandHistoryStack *Stack, undoCommandHistoryStack *Stack, ioWriter *bufio.Writer, commands []Command, keyboard *KeyBoard) *MainController {
	orginCmdLength := len(commands)
	return &MainController{
		commandHistoryStack:     commandHistoryStack,
		undoCommandHistoryStack: undoCommandHistoryStack,
		ioWriter:                ioWriter,
		commands:                commands,
		originCmdLength:         orginCmdLength,
		keyboard:                keyboard,
	}
}

func (mainController *MainController) BindKeyHandler(key string, command int) {
	mainController.keyboard.AddHotKey(key, command)
}

func (mainController *MainController) Undo() {
	prevCommand := mainController.commandHistoryStack.Pop()
	if prevCommand != nil {
		prevCommand.Undo()
		mainController.undoCommandHistoryStack.Push(prevCommand)
	}
}
func (mainController *MainController) Redo() {
	prevCommand := mainController.undoCommandHistoryStack.Pop()
	if prevCommand != nil {
		prevCommand.Execute()
		mainController.commandHistoryStack.Push(prevCommand)
	}
}
func (mainController *MainController) SetMacro(key string, commands []int) {
	macroCommand := NewMacroCommand(mainController.FindCommands(commands))
	commandIdx := len(mainController.keyboard.hotkeysMap) + len(mainController.commands)
	mainController.keyboard.AddMacroCommand(commandIdx, macroCommand)
	mainController.keyboard.AddHotKey(key, commandIdx)
}
func (mainController *MainController) Handle(key string) {
	command, ok := mainController.keyboard.CheckHotKey(key)
	if !ok {
		mainController.ioWriter.WriteString(fmt.Sprintf("%s command not support\n", key))
		return
	}
	var currentCmd Command
	if command < mainController.originCmdLength {
		currentCmd = mainController.commands[command]
	} else if _, ok := mainController.keyboard.macroCommands[command]; ok {
		currentCmd = mainController.keyboard.macroCommands[command]
	} else {
		mainController.ioWriter.WriteString(fmt.Sprintf("%d command not found\n", command))
		return
	}
	currentCmd.Execute()
	mainController.commandHistoryStack.Push(currentCmd)
	mainController.undoCommandHistoryStack.Clear()
}
func (mainController *MainController) FindCommands(commands []int) []Command {
	resultCommands := []Command{}
	hitMap := make(map[int]struct{})
	for _, command := range commands {
		hitMap[command] = struct{}{}
	}
	for commandIdx, command := range mainController.commands {
		if _, ok := hitMap[commandIdx]; ok {
			resultCommands = append(resultCommands, command)
		}
	}
	return resultCommands
}

func (mainController *MainController) ShowCommandList() {
	commandList := ""
	for idx, command := range mainController.commands {
		commandList += fmt.Sprintf("(%d): %s\n", idx, command.GetName())
	}
	mainController.ioWriter.WriteString(commandList)
	mainController.ioWriter.Flush()
}

func (mainController *MainController) ShowHotKeysList() {
	commandList := ""
	for key, command := range mainController.keyboard.GetHotkeyMap() {
		if command < len(mainController.commands) {
			commandList += fmt.Sprintf("%s: %s\n", key, mainController.commands[command].GetName())
		} else if _, ok := mainController.keyboard.GetMacroCommands()[command]; ok {
			commandList += fmt.Sprintf("%s: %s\n", key, mainController.keyboard.macroCommands[command].GetName())
		}
	}
	mainController.ioWriter.WriteString(commandList)
	mainController.ioWriter.Flush()
}
func (mainController *MainController) PromptOption() {
	mainController.ioWriter.WriteString("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) (q) 停止執行 按下按鍵:")
	mainController.ioWriter.Flush()
}
func (mainController *MainController) PromptStop() {
	mainController.ioWriter.WriteString("停止執行\n")
	mainController.ioWriter.Flush()
}
func (mainController *MainController) HandleSetup(isMacro bool, key string) bool {
	if isMacro {
		cmds, ok := mainController.keyboard.ProccessCommands()
		if !ok {
			return ok
		}
		mainController.SetMacro(key, cmds)
	} else {
		cmd, ok := mainController.keyboard.ProccessCommand()
		if !ok {
			return ok
		}
		mainController.BindKeyHandler(key, cmd)
	}
	return true
}
func (mainController *MainController) HandleAskIsMacro() (string, bool) {
	mainController.ioWriter.WriteString("設置巨集指令 (y/n)：")
	mainController.ioWriter.Flush()
	return mainController.keyboard.ProccessKey()
}
func (mainController *MainController) HandleAskKey() (string, bool) {
	mainController.ioWriter.WriteString("Key:")
	mainController.ioWriter.Flush()
	return mainController.keyboard.ProccessKey()
}
func (mainController *MainController) HandleAskOption() (string, bool) {
	mainController.ShowHotKeysList()
	mainController.PromptOption()
	return mainController.keyboard.ProccessKey()
}
func (mainController *MainController) PromptForSetup(isMacro bool, key string) {
	if isMacro {
		mainController.ioWriter.WriteString(fmt.Sprintf("要將哪些指令設置成快捷鍵 %s 的巨集（輸入多個數字，以空白隔開）: \n", key))
	} else {
		mainController.ioWriter.WriteString(fmt.Sprintf("要將哪一道指令設置到快捷鍵 %s 上:\n", key))
	}
	mainController.ioWriter.Flush()
	mainController.ShowCommandList()
}
func (mainController *MainController) HandleFlow() {
	isRun := true
	for isRun {
		option, ok := mainController.HandleAskOption()
		if !ok {
			continue
		}
		switch option {
		case "q":
			mainController.PromptStop()
			isRun = false
			continue
		case "1":
			isSetMacro, ok := mainController.HandleAskIsMacro()
			if !ok {
				continue
			}
			key, ok := mainController.HandleAskKey()
			if !ok {
				continue
			}
			setMarcoCmd := strings.Compare(isSetMacro, "y") == 0
			mainController.PromptForSetup(setMarcoCmd, key)
			if !mainController.HandleSetup(setMarcoCmd, key) {
				continue
			}
		case "2":
			mainController.Undo()
		case "3":
			mainController.Redo()
		default:
			mainController.Handle(option)
		}
	}
}
