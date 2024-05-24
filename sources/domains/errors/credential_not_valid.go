package errors

type CredentialNotValid struct{}

func (err CredentialNotValid) Error() string {
	return "Credential not valid"
}
