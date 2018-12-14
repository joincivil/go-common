package numbers_test

import (
	"math/big"
	"testing"

	"github.com/joincivil/go-common/pkg/numbers"
)

func TestBigIntToFloat64(t *testing.T) {
	floatVal := float64(3)
	bigInt := big.NewInt(3)
	floatNum := numbers.BigIntToFloat64(bigInt)
	if floatVal != floatNum {
		t.Errorf("Bigint to Float64 conversion failed")
	}
}

func TestFloat64ToBigInt(t *testing.T) {
	bigIntVal := big.NewInt(34)
	float := float64(34)
	bigInt := numbers.Float64ToBigInt(float)
	if bigInt == bigIntVal {
		t.Errorf("Float64 to Bigint conversion failed")
	}
}
