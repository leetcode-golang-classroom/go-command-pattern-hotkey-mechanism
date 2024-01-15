package hotkey

type Stack struct {
	commands []Command
}

func NewStack() *Stack {
	return &Stack{
		commands: []Command{},
	}
}

func (stack *Stack) Pop() Command {
	if len(stack.commands) == 0 {
		return nil
	}
	topIdx := len(stack.commands) - 1
	prevCmd := stack.commands[topIdx]
	stack.commands = stack.commands[:topIdx]
	return prevCmd
}

func (stack *Stack) Push(command Command) {
	stack.commands = append(stack.commands, command)
}

func (stack *Stack) Clear() {
	stack.commands = []Command{}
}
