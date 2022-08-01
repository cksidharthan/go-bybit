package bybit

import (
	"fmt"
	"strconv"
	"time"
)

func GetPrecision(value float64) (float64, error) {
	qtyString := fmt.Sprintf("%.3f", value)
	orderQty, err := strconv.ParseFloat(qtyString, 64)
	if err != nil {
		return 0, err
	}
	return orderQty, nil
}

func GetCurrentTime() string {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	return strconv.Itoa(intNow)
}
