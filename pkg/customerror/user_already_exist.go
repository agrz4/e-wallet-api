package customerror

type UserAlreadyExistError struct {
}

func (e *UserAlreadyExistError) Error() string {
	return "user already exist"
}
