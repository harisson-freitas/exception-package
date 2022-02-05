package error

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	http_code "github.com/harisson-freitas/exception-package/error/enum/http"
	system "github.com/harisson-freitas/exception-package/error/enum/system"
)

func GetErrorWithStatusCode(w http.ResponseWriter, error *Error, statusCode int) {
	message := createMessage(error, statusCode)
	response, _ := json.Marshal(message)

	w.WriteHeader(statusCode)
	w.Write(response)
}

func GetStatusCode(error *Error, statusCode int) int {
	isConflict := strings.Contains(error.Message, system.ConflictSave)
	if isConflict {
		return http.StatusConflict
	} else {
		return statusCode
	}
}

func createMessage(error *Error, statusCode int) *Error {
	switch statusCode {
	case http.StatusNotFound:
		return getError(error, http_code.StatusNotFound)
	case http.StatusConflict:
		return getError(error, http_code.StatusConflict)
	case http.StatusInternalServerError:
		return getError(error, http_code.StatusInternalServerError)
	default:
		return nil
	}
}

func getError(error *Error, statusCode int) *Error {
	error.Message = fmt.Sprintf("%s: %s", getStatusText(statusCode), error.Message)
	return error
}

func getStatusText(statusCode int) string {
	return http.StatusText(statusCode)
}
