package customerror

type PasswordNotSame struct {
}

func (e *PasswordNotSame) Error() string {
	return "password is not the same as confirm password"
}
