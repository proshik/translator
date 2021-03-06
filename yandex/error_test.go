package yandex

import (
	"github.com/pkg/errors"
	"testing"
)

var err = ExternalError{errors.New("error message"), UNEXPECTED_ERROR}

func TestExternalError_String(t *testing.T) {
	result := err.String()
	if result != "error message, with code=6" {
		t.Error("incorrect string message")
	}

}
