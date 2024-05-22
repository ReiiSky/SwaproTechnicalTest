package objects

type InformationNumber[T string | int] struct {
	value T
}

func NewInformationNumber[T string | int](value T) InformationNumber[T] {
	return InformationNumber[T]{value}
}

func (in *InformationNumber[T]) GetNumber() T {
	return in.value
}
