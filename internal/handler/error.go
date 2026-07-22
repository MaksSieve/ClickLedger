package handler

import (
	"clickledger/internal/service"
	"errors"
	"fmt"
	"net/http"
)

func HandleError(writer http.ResponseWriter, err error) {
	if errors.Is(err, &ErrRequestParameterInvalid{}) {
		http.Error(writer, err.Error(), 400)
	} else if errors.Is(err, service.ErrLinkNotFound) {
		http.Error(writer, err.Error(), 404)
	} else {
		http.Error(writer, err.Error(), 500)
	}
}

type ErrRequestParameterInvalid struct {
	parameter           string
	expectedValueOrType string
	givenValueOrType    string
}

func NewErrRequestParameterInvalid(param string, expectedValueOrType string, givenValueOrType string) *ErrRequestParameterInvalid {
	return &ErrRequestParameterInvalid{
		parameter:           param,
		expectedValueOrType: expectedValueOrType,
		givenValueOrType:    givenValueOrType,
	}
}
func (err *ErrRequestParameterInvalid) Error() string {
	return fmt.Sprintf("%s should be %s, but got %s", err.parameter, err.expectedValueOrType, err.givenValueOrType)
}
