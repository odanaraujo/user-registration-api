package httperr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"err,omitempty"`
	Code    int      `json:"code"`
	Fields  []Fields `json:"fields,omitempty"`
}

type Fields struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value,omitempty"`
	Message string      `json:"message"`
}

func (err *RestErr) Error() string {
	return err.Message
}

func NewRestErr(message, err string, code int, fields []Fields) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Fields:  fields,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationsError(message string, fields []Fields) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Fields:  fields,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNoTFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
