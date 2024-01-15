package hotkey

import (
	"bufio"
	"strconv"
	"strings"
)

type KeyBoard struct {
	ioReader      *bufio.Reader
	ioWriter      *bufio.Writer
	hotkeysMap    map[string]int
	macroCommands map[int]Command
}

func NewKeyBoard(ioReader *bufio.Reader, ioWriter *bufio.Writer) *KeyBoard {
	return &KeyBoard{
		ioReader:      ioReader,
		ioWriter:      ioWriter,
		hotkeysMap:    make(map[string]int),
		macroCommands: make(map[int]Command),
	}
}
func (keyboard *KeyBoard) AddMacroCommand(idx int, command Command) {
	keyboard.macroCommands[idx] = command
}
func (keyboard *KeyBoard) AddHotKey(key string, command int) {
	keyboard.hotkeysMap[key] = command
}
func (keyboard *KeyBoard) ResetHotKey() {
	keyboard.hotkeysMap = make(map[string]int)
}
func (keyboard *KeyBoard) ResetMacroCommand() {
	keyboard.macroCommands = make(map[int]Command)
}
func (keyboard *KeyBoard) CheckHotKey(key string) (int, bool) {
	command, ok := keyboard.hotkeysMap[key]
	return command, ok
}
func (keyboard *KeyBoard) ProccessKey() (string, bool) {
	input, _, _ := keyboard.ioReader.ReadLine()
	key := string(input)
	var rKey rune
	for idx, r := range key {
		if idx == 0 {
			rKey = r
			break
		}
	}
	ok := keyboard.IsValidKey(rKey) && len(key) == 1
	return key, ok
}
func (keyboard *KeyBoard) ProccessCommand() (int, bool) {
	input, _, _ := keyboard.ioReader.ReadLine()
	command := string(input)
	command = strings.TrimSpace(command)
	resultCmd := 0
	parseResult, err := strconv.Atoi(command)
	if err != nil {
		return 0, false
	}
	resultCmd = parseResult
	return resultCmd, true
}
func (keyboard *KeyBoard) ProccessCommands() ([]int, bool) {
	input, _, _ := keyboard.ioReader.ReadLine()
	commands := string(input)
	commands = strings.TrimSpace(commands)
	splitCmds := strings.Split(commands, " ")
	resultCommands := []int{}
	for _, cmd := range splitCmds {
		cmdIdx, err := strconv.Atoi(cmd)
		if err != nil {
			return []int{}, false
		}
		resultCommands = append(resultCommands, cmdIdx)
	}
	return resultCommands, true
}
func (keyboard *KeyBoard) IsValidKey(key rune) bool {
	return (key >= 'a' && key <= 'z') || (key >= 49 && key <= 51)
}
func (keyboard *KeyBoard) GetHotkeyMap() map[string]int {
	return keyboard.hotkeysMap
}
func (keyboard *KeyBoard) GetMacroCommands() map[int]Command {
	return keyboard.macroCommands
}
