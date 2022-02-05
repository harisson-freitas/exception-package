package error

type Error struct {
	TypeName string `json:"type"`
	Message  string `json:"message"`
}

func CreateError(typeName string, message string) (e *Error) {
	return &Error{TypeName: typeName, Message: message}
}
