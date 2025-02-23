package customerror

type WalletAlreadyExistsError struct {
}

func (e *WalletAlreadyExistsError) Error() string {
	return "wallet already exists"
}
