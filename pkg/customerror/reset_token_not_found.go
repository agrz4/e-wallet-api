package customerror

type ResetTokenNotFound struct {
}

func (e *ResetTokenNotFound) Error() string {
	return "invalid reset token"
}
