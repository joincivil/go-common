package eth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
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
// Handles signatures with and without 0x prefixes
func VerifyEthSignature(address string, message string, signature string) (bool, error) {
	// sig must be a 65-byte compact ECDSA signature containing the
	// recovery id as the last element.
	addressBytes, err := hexutil.Decode(address)
	if err != nil {
		return false, errors.New("address appears to be invalid")
	}

	signatureBytes, err := DecodeSignatureBytes(signature)
	if err != nil {
		return false, errors.New("signature appears to be invalid")
	}

	// TODO this is a hack to set the ECDSA signature recovery ID to 0
	// web3 is returning 27 or 28, but should be 0 or 1
	// https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethsign
	// https://github.com/ethereum/go-ethereum/blob/6cd6b921ac57480d95af8b9bec2424e1f89fa196/crypto/secp256k1/secp256.go
	if signatureBytes[64] == 27 {
		signatureBytes[64] = 0
	}
	if signatureBytes[64] == 28 {
		signatureBytes[64] = 1
	}

	message = AsEthereumSignature(message)
	hash := crypto.Keccak256Hash([]byte(message))

	pubKey, err := crypto.SigToPub(hash.Bytes(), signatureBytes)
	if err != nil {
		return false, err
	}

	sigAddr := crypto.PubkeyToAddress(*pubKey)

	// If addresses don't match, then return false
	if !bytes.Equal(addressBytes, sigAddr.Bytes()) {
		return false, nil
	}

	return true, nil
}

// VerifyEthSignatureWithPubkey accepts an ECDSA public key, a message string, and a signature and
// confirms that the signature was indeed signed by the key specified
// Handles signatures with and without 0x prefixes
func VerifyEthSignatureWithPubkey(pubKey ecdsa.PublicKey, message string, signature string) (bool, error) {
	addr := crypto.PubkeyToAddress(pubKey)
	return VerifyEthSignature(addr.Hex(), message, signature)
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

// SignEthMessage signs a given message with the given ECDSA private key.
// Returns the signature as a hex string with 0x prefix.
func SignEthMessage(pk *ecdsa.PrivateKey, message string) (string, error) {
	message = AsEthereumSignature(message)
	hash := crypto.Keccak256([]byte(message))
	signature, err := crypto.Sign(hash, pk)
	if err != nil {
		return "", err
	}
	hexSig := hexutil.Encode(signature)
	return hexSig, nil
}

// AsEthereumSignedMessage adds a prefix and len to the message to identify it
// as an Ethereum specific signature.
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sign
func AsEthereumSignature(msg string) string {
	return fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
}

// DecodeSignatureBytes decodes a signature to bytes.  Handles signature
// hex strings with or without 0x prefix.
func DecodeSignatureBytes(signature string) ([]byte, error) {
	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		signatureBytes, err = hex.DecodeString(signature)
		if err != nil {
			return nil, err
		}
	}
	return signatureBytes, nil
}
