package api

import (
	"fmt"
	"strings"
)

type RequestValidationError struct {
	messages []string
	model    string
	method   string
}

func (rve *RequestValidationError) Error() string {
	msgs := strings.Join(rve.messages, ", ")
	return fmt.Sprintf("Error validation at %s, method: %s: %s", rve.model, rve.method, msgs)
}

type ResponseError struct {
	message string
}

func (re *ResponseError) Error() string {
	return re.message
}

type CatalogParentError string

func (cpe CatalogParentError) Error() string {
	return "Not found catalog parent"
}

