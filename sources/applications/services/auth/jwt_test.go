package auth_test

import (
	"testing"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/services/auth"
)

var (
	token         = "secret-token"
	authenticator = auth.NewJWTAuthentication(token)
	payload       = auth.AuthPayload{EmployeeID: 1}
)

func TestEncode(t *testing.T) {
	encoded := authenticator.Encode(payload)

	if len(encoded) <= 0 {
		t.Error("Empty encoded token length")
	}
}

func TestDecode(t *testing.T) {
	encoded := authenticator.Encode(payload)
	decoded, err := authenticator.Decode(encoded)

	if err != nil {
		t.Errorf("Error decode not nill expected nill with message: %s", err.Error())
	}

	if decoded.EmployeeID != payload.EmployeeID {
		t.Errorf("decoded employee id is not %d", decoded.EmployeeID)
	}
}
