package hotkey

import "bufio"

type Telecom struct {
	isConnected bool
	ioWriter    *bufio.Writer
}

func NewTelecom(ioWriter *bufio.Writer) *Telecom {
	return &Telecom{
		isConnected: false,
		ioWriter:    ioWriter,
	}
}

func (telecom *Telecom) Connect() {
	telecom.isConnected = true
	telecom.ioWriter.WriteString("The telecom has been turned on.\n")
	telecom.ioWriter.Flush()
}

func (telecom *Telecom) Disconnect() {
	telecom.isConnected = false
	telecom.ioWriter.WriteString("The telecom has been turned off.\n")
	telecom.ioWriter.Flush()
}
