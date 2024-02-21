package tax

import (
	"errors"
	"fmt"
)

type taxBand struct {
	taxRate   float64
	bandStart float64
	bandEnd   float64
}

func (t taxBand) calculateTax(amount float64) float64 {
	fmt.Println(t)
	return amount
}

func initTaxBand() [3]taxBand {

	taxes := [3]taxBand{{
		taxRate:   0,
		bandStart: 0,
		bandEnd:   10,
	}, {
		taxRate:   10,
		bandStart: 11,
		bandEnd:   20,
	}, {
		taxRate:   20,
		bandStart: 21,
		bandEnd:   30,
	}}

	return taxes
}

func GetTaxAmount(amount float64) (float64, error) {
	var taxAmount float64
	taxes := initTaxBand()
	for _, val := range taxes {
		if val.bandStart >= amount && val.bandEnd <= amount {
			taxAmount = val.calculateTax(amount)
			return taxAmount, nil
		}
	}
	return 0, errors.New("Tax Amount No Found in Tax Band")
}
