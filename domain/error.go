package domain

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrTodoNotFound = Error("todo not found")
)
