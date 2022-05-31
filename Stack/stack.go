package stack

const DefaultSize = 128

type Stack[T any] interface {
	Push(element T)
	Top() T
	Pop()
	Len() int
	IsEmpty() bool
}
type stack[T any] struct {
	elements []T
	sp       int
}

func (s *stack[T]) expand() {
	l := len(s.elements) << 1
	x := make([]T, l)
	copy(x, s.elements)
	s.elements = x
}

func (s *stack[T]) Push(element T) {
	s.sp++
	if s.sp == len(s.elements) {
		s.expand()
	}
	s.elements[s.sp] = element
}
func (s *stack[T]) Top() T {
	return s.elements[s.sp]
}
func (s *stack[T]) Pop() {
	s.sp--
}
func (s *stack[T]) Len() int {
	return s.sp
}
func (s *stack[T]) IsEmpty() bool {
	return s.sp == -1
}
func NewStack[T any]() *stack[T] {
	return &stack[T]{make([]T, DefaultSize), -1}
}
func NewStackN[T any](n int) *stack[T] {
	return &stack[T]{make([]T, n), -1}
}
