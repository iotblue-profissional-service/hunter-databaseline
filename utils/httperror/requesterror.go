package httperror

import "net/http"

type RestErr struct {
	Code           int    `json:"code"`
	MessageKeys    string `json:"messageKeys"`
	Message        string `json:"message"`
	IsInternalOnly bool   `json:"isInternalOnly,omitempty"`
	ToBeLogged     bool   `json:"toBeLogged,omitempty"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Code:        http.StatusBadRequest,
		MessageKeys: "invalid body",
		Message:     message,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        http.StatusInternalServerError,
		MessageKeys: "Internal server error",
	}
}

func NewUnAuthenticatedServerError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        http.StatusUnauthorized,
		MessageKeys: "user_is_unauthorized",
	}
}

func NewCreateFailedError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        4001,
		MessageKeys: "Creation Failure",
	}
}

func NewUpdateFailedError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        4002,
		MessageKeys: " Updating Failure",
	}
}

func NewDeleteFailedError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        4004,
		MessageKeys: " Deleting Failure",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Code:           4004,
		MessageKeys:    "resource_not_found",
		Message:        message,
		IsInternalOnly: false,
		ToBeLogged:     false,
	}
}
