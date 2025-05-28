package analytics

var _ Message = (ErrorMessage)(nil)

type ErrorMessage interface {
	Message
	ErrorText() string
}

func NewErrorMessage(err error) ErrorMessage {
	return &errorMessage{message: message{event: "error"}, err: err}
}

var _ Message = (*errorMessage)(nil)

type errorMessage struct {
	message
	err error
}

func (e *errorMessage) ErrorText() string {
	return e.err.Error()
}

func (e *errorMessage) Validate() error {
	if err := e.err; err != nil {
		return err
	}
	if e.err == nil {
		return nil
	}
	return nil
}
