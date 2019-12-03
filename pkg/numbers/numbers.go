package numbers

import (
	"math/big"
	"strconv"
)

// BigIntToFloat64 converts big.Int to float64 type
func BigIntToFloat64(bigInt *big.Int) float64 {
	f := new(big.Float).SetInt(bigInt)
	val, _ := f.Float64()
	return val
}

// Float64ToBigInt converts float64 to big.Int type
func Float64ToBigInt(float float64) *big.Int {
	bigInt := new(big.Int)
	bigInt.SetString(strconv.FormatFloat(float, 'f', -1, 64), 10)
	return bigInt
}

// ListIntToListString converts a list of big.int to a list of string
func ListIntToListString(listInt []int) []string {
	listString := make([]string, len(listInt))
	for idx, i := range listInt {
		listString[idx] = strconv.Itoa(i)
	}
	return listString
}

// IntOrEmptyInt returns an empty int if a int pointer is nil.
// Otherwise returns the int value of the pointer.
func IntOrEmptyInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

// IntToPtr is a convenience func that converts a int to it's pointer
func IntToPtr(s int) *int {
	return &s
}
