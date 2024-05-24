package persistences

import "errors"

var (
	ErrQuery         = errors.New("error call query to db")
	ErrNoInsertAny   = errors.New("no insert any data to db")
	ErrPartialInsert = errors.New("partial insert data to db")
)
