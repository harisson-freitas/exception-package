package error_handler

import (
	// "encoding/json"
	// "fmt"
	"fmt"
	"net/http"

	"github.com/harisson-freitas/exception-package/error"
	type_name "github.com/harisson-freitas/exception-package/error/handler/enum"
)

func ErrorValidationHandler(w http.ResponseWriter, err *error.Error) {
	code := getStatusCode(err)
	getErrorValidation(w, err, code)
}

func ErrorValidationRouterHandler(w http.ResponseWriter, err *error.Error, statusCode int) {
	if statusCode == http.StatusNotFound {
		statusNotFoundError(w, err)
	}
}

func getStatusCode(err *error.Error) int {
	switch err.TypeName {
	case type_name.DATABASE:
		return error.GetStatusCode(err, http.StatusInternalServerError)
	case type_name.JSON, type_name.PARSE:
		return http.StatusInternalServerError
	default:
		return 0
	}
}

func getErrorValidation(w http.ResponseWriter, err *error.Error, statusCode int) {
	if statusCode == http.StatusInternalServerError {
		internalServerError(w, err)
	} else if statusCode == http.StatusConflict {
		statusConflictError(w, err)
	} else {
		fmt.Println("FAIL CODE!!")
	}
}

func internalServerError(w http.ResponseWriter, err *error.Error) {
	error.GetErrorWithStatusCode(w, err, http.StatusInternalServerError)
}

func statusConflictError(w http.ResponseWriter, err *error.Error) {
	error.GetErrorWithStatusCode(w, err, http.StatusConflict)
}

func statusNotFoundError(w http.ResponseWriter, err *error.Error) {
	error.GetErrorWithStatusCode(w, err, http.StatusNotFound)
}
