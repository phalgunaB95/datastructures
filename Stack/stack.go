package stack

type Stack[T any] interface {
	Push(element T)
	Pop() (T, bool)
	Len() int
}
type stack[T any] struct {
	elements []T
}

func (s *stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *stack[T]) Pop() (T, bool) {
	var top T
	l := len(s.elements)
	if l == 0 {
		return top, false
	}

	top = s.elements[l-1]
	s.elements = s.elements[:l-1]
	return top, true
}

func (s *stack[T]) Len() int {
	return len(s.elements)
}
func NewStack[T any]() *stack[T] {
	return &stack[T]{make([]T, 0)}
}
