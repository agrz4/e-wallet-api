package utils

import "fmt"

func GenerateWalletNumber(userID uint) string {
	if userID < 10 {
		return fmt.Sprint("10000%d", userID)
	} else if userID < 100 {
		return fmt.Sprint("1000%d", userID)
	} else {
		return fmt.Sprint("100%d", userID)
	}
}
