package bybit

import (
	"fmt"
	"strconv"
)

func GetPrecision(value float64) (float64, error) {
	qtyString := fmt.Sprintf("%.3f", value)
	orderQty, err := strconv.ParseFloat(qtyString, 64)
	if err != nil {
		return 0, err
	}
	return orderQty, nil
}
