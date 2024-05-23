package objects

type Identifier[T string | int] struct {
	id T
}

func NewIdentifier[T string | int](value T) Identifier[T] {
	return Identifier[T]{id: value}
}

func (i Identifier[T]) Equal(other Identifier[T]) bool {
	return i.id == other.id
}

func GetStringIdentifier(i Identifier[string]) string {
	return i.id
}

func GetNumberIdentifier(i Identifier[int]) int {
	return i.id
}
