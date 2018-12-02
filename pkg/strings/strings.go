// Package strings contains various common utils for strings
package strings

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

var (
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|ethereum|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
)

// IsValidEthAPIURL returns true if the given string matches a valid
// eth endpoint URL
func IsValidEthAPIURL(url string) bool {
	return validEthAddressExp.MatchString(url)
}

// IsValidContractAddress returns true is the given string matches a valid
// smart contract address
func IsValidContractAddress(address string) bool {
	return common.IsHexAddress(address)
}

// RandomHexStr generates a hex string from a byte slice of n random numbers.
func RandomHexStr(n int) (string, error) {
	bys := make([]byte, n)
	if _, err := rand.Read(bys); err != nil {
		return "", err
	}
	return hex.EncodeToString(bys), nil
}