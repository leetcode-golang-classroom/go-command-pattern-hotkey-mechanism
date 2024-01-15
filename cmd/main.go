package main

import (
	"bufio"
	"os"

	"github.com/leetcode-golang-classroom/go-command-pattern-hotkey-mechanism/hotkey"
)

func main() {
	ioReader := bufio.NewReader(os.Stdin)
	ioWriter := bufio.NewWriter(os.Stdout)
	keyBoard := hotkey.NewKeyBoard(ioReader, ioWriter)
	tank := hotkey.NewTank(ioWriter)
	telecom := hotkey.NewTelecom(ioWriter)
	commandHistoryStack := hotkey.NewStack()
	undoCommandHistoryStack := hotkey.NewStack()
	mainController := hotkey.NewMainController(commandHistoryStack, undoCommandHistoryStack, ioWriter,
		[]hotkey.Command{
			hotkey.NewMoveForwardCommand(tank),
			hotkey.NewMoveBackwardCommand(tank),
			hotkey.NewConnectCommand(telecom),
			hotkey.NewDisconnectCommand(telecom),
			hotkey.NewResetCommand(keyBoard),
		}, keyBoard)

	mainController.HandleFlow()
}
