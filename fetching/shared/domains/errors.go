package domains

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")
	ErrUnauthenticate      = errors.New("Unauthenticate")
)

type ResponseError struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
