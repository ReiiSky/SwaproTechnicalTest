package auth

import (
	"errors"

	jwt "github.com/golang-jwt/jwt"
)

var (
	ErrEncodedTextNotValid = errors.New("encoded string not valid")
	ErrParsedClaimNotValid = errors.New("claim not valid")
)

type JWTAuthentication struct {
	token []byte
}

func NewJWTAuthentication(token string) JWTAuthentication {
	return JWTAuthentication{[]byte(token)}
}

func (auth *JWTAuthentication) Encode(payload AuthPayload) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"employee_id": payload.EmployeeID,
	})

	encodedToken, _ := claims.SignedString(auth.token)
	return encodedToken
}

func (auth *JWTAuthentication) Decode(encoded string) (AuthPayload, error) {
	payload := AuthPayload{}
	parsedClaim, err := jwt.ParseWithClaims(encoded, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return auth.token, nil
	})

	if err != nil {
		return payload, ErrEncodedTextNotValid
	}

	kvClaim := *parsedClaim.Claims.(*jwt.MapClaims)

	if id, ok := kvClaim["employee_id"].(float64); ok {
		payload.EmployeeID = int(id)
	} else {
		return payload, ErrParsedClaimNotValid
	}

	return payload, nil
}
