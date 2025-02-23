package customerror

type InsufficientBalanceError struct {
}

func (e *InsufficientBalanceError) Error() string {
	return "insufficient ballance"
}
