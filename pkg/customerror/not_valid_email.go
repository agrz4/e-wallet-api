package customerror

type NotValidEmailError struct {
}

func (e *NotValidEmailError) Error() string {
	return "not a valid email"
}
