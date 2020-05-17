package domains

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error.")
	ErrNotFound            = errors.New("Your requested item is not found.")
	ErrConflict            = errors.New("Your item already exist.")
	ErrBadParamInput       = errors.New("Given param is not valid.")
	ErrUnprocessableEntity = errors.New("Unprocessable entity.")
	ErrUnauthenticate      = errors.New("Unauthenticated.")
)

type ResponseError struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
