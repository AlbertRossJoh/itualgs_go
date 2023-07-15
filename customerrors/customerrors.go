package customerrors

type ErrQueueFull struct{}
type ErrQueueEmpty struct{}
type ErrEmptyStack struct{}
type ErrZeroVector struct{}
type ErrVectorCross struct{}
type ErrMatrixNotSquare struct{}

func (e *ErrQueueFull) Error() string {
	return "Queue is full"
}

func (e *ErrQueueEmpty) Error() string {
	return "Queue is empty"
}

func (e *ErrEmptyStack) Error() string {
	return "Stack is empty"
}

func (e *ErrZeroVector) Error() string {
	return "Vector is empty"
}

func (e *ErrVectorCross) Error() string {
	return "Crossing is currently only supported for 2D vectors"
}

func (e *ErrMatrixNotSquare) Error() string {
	return "Invalid operation"
}
