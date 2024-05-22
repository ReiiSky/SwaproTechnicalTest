package objects

type CryptedInformation struct {
	password string
	method   string
}

func NewMD5CryptedInformation(information string) CryptedInformation {
	return CryptedInformation{
		password: information,
		method:   "md5",
	}
}
