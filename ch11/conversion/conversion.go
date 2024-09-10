package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strs []string) ([]float64, error) {
	l := len(strs)
	ans := make([]float64, l)
	for idx, num := range strs {
		res, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return nil, errors.New("error converting from string to float64")
		}
		ans[idx] = res
	}
	return ans, nil
}
