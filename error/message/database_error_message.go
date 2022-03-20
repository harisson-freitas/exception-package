package message

import (
	"fmt"

	"github.com/harisson-freitas/exception-package/error"
)

const TYPE_NAME = "DATABASE"

func GetMessageErrorOpenConnection() (err *error.Error) {
	return error.CreateError(TYPE_NAME, "Error Open Connection!")
}

func GetMessageErrorConfigDatabase() (err *error.Error) {
	return error.CreateError(TYPE_NAME, "Database configuration error!")
}

func GetMessageErrorFind(e string) (err *error.Error) {
	msg := fmt.Sprintf("Error fetching user from database! %s", e)
	return error.CreateError(TYPE_NAME, msg)
}

func GetMessageErrorCreate(e string) (err *error.Error) {
	msg := fmt.Sprintf("Error inserting user into database! %s", e)
	return error.CreateError(TYPE_NAME, msg)
}

func GetMessageErrorUpdate(e string) (err *error.Error) {
	msg := fmt.Sprintf("Error updating user in database! %s", e)
	return error.CreateError(TYPE_NAME, msg)
}

func GetMessageErrorDelete(e string) (err *error.Error) {
	msg := fmt.Sprintf("Error deleting user in database! %s", e)
	return error.CreateError(TYPE_NAME, msg)
}
