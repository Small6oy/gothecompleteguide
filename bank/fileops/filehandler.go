package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func StoreFloatInFile(accountBalance string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalance, []byte(balanceText), 0644)
}
func GetFloatFromFile(accountBalance string) (float64, error) {
	data, err := os.ReadFile(accountBalance)
	if err != nil {
		return 0, errors.New("Failed to Read File")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("Failed to Parse File")
	}

	return balance, nil
}
