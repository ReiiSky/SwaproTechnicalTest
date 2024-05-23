package errors

type LocationNotExist struct{}

func (err LocationNotExist) Error() string {
	return "Location not exist"
}
