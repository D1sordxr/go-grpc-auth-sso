package exceptions

type Exception struct {
	Message string
	Data    interface{}
	Error   error
}

func (e Exception) Err() string {
	return e.Error.Error()
}
