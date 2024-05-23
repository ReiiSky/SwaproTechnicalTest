package services

import "github.com/ReiiSky/SwaproTechnical/sources/applications/services/auth"

type Auth interface {
	Decode(string) (auth.AuthPayload, error)
	Encode(auth.AuthPayload) string
}
