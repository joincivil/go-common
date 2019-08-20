package eth

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// ChallengeRequest contains a signature and the fields needed to verify it
type ChallengeRequest struct {
	ExpectedPrefix string
	GracePeriod    int64
	InputAddress   string
	InputChallenge string
	Signature      string
}

// VerifyEthChallengeAndSignature accepts a ChallengeRequest and verifies that it is valid
func VerifyEthChallengeAndSignature(request ChallengeRequest) error {
	if err := VerifyEthChallenge(request.ExpectedPrefix, request.GracePeriod, request.InputChallenge); err != nil {
		return err
	}
	verified, err := VerifyEthSignature(request.InputAddress, request.InputChallenge, request.Signature)

	if err != nil {
		return err
	}
	if !verified {
		return errors.New("could not verify signature")
	}

	return nil
}

// VerifyEthSignature accepts an Ethereum address, a message string, and a signature and
// confirms that the signature was indeed signed by the address specified
func VerifyEthSignature(address string, message string, signature string) (bool, error) {
	// sig must be a 65-byte compact ECDSA signature containing the
	// recovery id as the last element.
	addressBytes, err := hexutil.Decode(address)
	if err != nil {
		return false, errors.New("Address appears to be invalid")
	}
	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return false, errors.New("Signature appears to be invalid")
	}
	// TODO this is a hack to set the ECDSA signature recovery ID to 0
	// web3 is returning 27 or 28, but should be 0 or 1
	// https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethsign
	// https://github.com/ethereum/go-ethereum/blob/6cd6b921ac57480d95af8b9bec2424e1f89fa196/crypto/secp256k1/secp256.go
	if signatureBytes[64] == 27 || signatureBytes[64] == 28 {
		signatureBytes[64] = 0
	}

	var hash = crypto.Keccak256([]byte(message))
	sigAddress, err := crypto.Ecrecover(hash, signatureBytes)

	if err != nil {
		return false, err
	}

	if bytes.Equal(addressBytes, sigAddress) {
		return false, errors.New("signature does not match")
	}

	return true, nil
}

// VerifyEthChallenge confirms that a "challenge" string has a timestamp that is
// within {gracePeriod} number of seconds from the current time
// this is used to ensure that an attacker cannot reuse previously signed messages
// Challenge should be in the form of "{prefix} @ 2018-01-04T17:48:32-05:00"
func VerifyEthChallenge(prefix string, gracePeriod int64, challenge string) error {
	var parts = strings.Split(challenge, " @ ")
	if parts[0] != prefix {
		return errors.New("challenge does not start with `" + prefix + "`")
	}
	if len(parts) != 2 {
		return errors.New("challenge does not have time portion")
	}

	var now = time.Now()
	var input, err = time.Parse(time.RFC3339, parts[1])
	if err != nil {
		return err
	}

	if (now.Unix() - input.Unix()) > gracePeriod {
		return errors.New("expired")
	}

	return nil
}
