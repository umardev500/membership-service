package helper

import (
	"strconv"
	"strings"
)

func RemovePriceDot(price string) (out int64) {
	price = strings.Split(price, ".")[0]
	priceInt, _ := strconv.Atoi(price)

	return int64(priceInt)
}
