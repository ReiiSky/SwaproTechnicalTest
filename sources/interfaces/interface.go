package interfaces

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications"
)

type StopableProcess interface {
	applications.Process
	Close()
}

type IKernel interface {
	NewProcess() StopableProcess
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Status string         `json:"status"`
	Data   interface{}    `json:"data"`
	Error  *ErrorResponse `json:"error"`
}

func NewOKResponse(data interface{}) Response {
	return Response{Status: "ok", Data: data}
}

func NewFailedResponse(data interface{}, err error) Response {
	var errResponse *ErrorResponse

	if err != nil {
		errResponse = &ErrorResponse{
			Message: err.Error(),
		}
	}

	return Response{
		Status: "failed",
		Data:   data,
		Error:  errResponse,
	}
}
