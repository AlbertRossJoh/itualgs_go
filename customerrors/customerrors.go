package customerrors

type ErrQueueFull struct{}
type ErrQueueEmpty struct{}
type ErrEmptyStack struct{}

func (e *ErrQueueFull) Error() string {
	return "Queue is full"
}

func (e *ErrQueueEmpty) Error() string {
	return "Queue is empty"
}

func (e *ErrEmptyStack) Error() string {
	return "Stack is empty"
}
