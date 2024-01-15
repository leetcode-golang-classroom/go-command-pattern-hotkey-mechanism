package hotkey

import (
	"bufio"
)

type Tank struct {
	position int
	ioWriter *bufio.Writer
}

func NewTank(ioWriter *bufio.Writer) *Tank {
	return &Tank{
		position: 0,
		ioWriter: ioWriter,
	}
}

func (tank *Tank) MoveForward() {
	tank.position += 1
	tank.ioWriter.WriteString("The tank has moved forward.\n")
	tank.ioWriter.Flush()
}

func (tank *Tank) MoveBackward() {
	tank.position -= 1
	tank.ioWriter.WriteString("The tank has moved backward.\n")
	tank.ioWriter.Flush()
}
