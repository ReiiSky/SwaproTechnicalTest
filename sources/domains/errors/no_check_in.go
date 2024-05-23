package errors

type NoAvailableCheckIn struct{}

func (err NoAvailableCheckIn) Error() string {
	return "No Avilable checkin attendance"
}
