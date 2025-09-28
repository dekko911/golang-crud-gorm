package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

const (
	ISE    = http.StatusInternalServerError
	OK     = http.StatusOK
	BR     = http.StatusBadRequest
	NF     = http.StatusNotFound
	CRD    = http.StatusCreated
	FORBID = http.StatusForbidden
)

var Validate = validator.New()
