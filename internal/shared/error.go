package shared

type Error interface {
	error
	GetMessage() string
	GetCode() int
	SetCode(int) Error
}

type Err struct {
	message string
	code    int
}

func NewErr(message string) *Err {
	defaultCode := 500
	return &Err{message: message, code: defaultCode}
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) GetMessage() string { return e.message }

func (e *Err) GetCode() int { return e.code }

func (e *Err) SetCode(code int) Error {
	e.code = code
	return e
}
