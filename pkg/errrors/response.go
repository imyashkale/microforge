package errors

import (
	"net/http"
	"sort"

	"github.com/go-playground/validator/v10"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func NewResponse(err interface{}) ErrorResponse {
	switch e := err.(type) {
	case ErrorResponse:
		return err.(ErrorResponse)
	case validator.ValidationErrors:
		return InvalidInput(e)
	}
	return InternalServerError("")
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: msg,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

type InvalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InvalidInput creates a new error response for data validation errors (HTTP 400).
// It uses validator.ValidationErrors to construct a detailed list of field errors.
func InvalidInput(errs validator.ValidationErrors) ErrorResponse {
	var details []InvalidField

	// Extract field names and sort them for consistent error ordering.
	var fields []string
	for _, err := range errs {
		fields = append(fields, err.Field())
	}
	sort.Strings(fields)

	// Construct a detailed message for each field error.
	for _, field := range fields {
		for _, err := range errs {
			if err.Field() == field {
				details = append(details, InvalidField{
					Field: field,
					Error: err.Error(),
				})
				break // Break the inner loop once the matching field error is added.
			}
		}
	}

	// Return the ErrorResponse struct with sorted field errors.
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}