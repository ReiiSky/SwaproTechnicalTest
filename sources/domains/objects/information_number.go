package objects

type InformationNumber[T string | int] struct {
	value T
}

func NewInformationNumber[T string | int](value T) InformationNumber[T] {
	return InformationNumber[T]{value}
}

func GetStringInformationNumber(in InformationNumber[string]) string {
	return in.value
}
