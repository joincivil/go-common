package eth_test

import (
	"encoding/hex"

	"testing"
	"time"

	"github.com/joincivil/go-common/pkg/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifyEthChallengeAndSignature(t *testing.T) {
	privKey, _ := crypto.GenerateKey()
	pubKeyBys := crypto.FromECDSAPub(&privKey.PublicKey)
	pubKey, _ := crypto.UnmarshalPubkey(pubKeyBys)
	sigAddr := crypto.PubkeyToAddress(*pubKey)
	message := LoginChallenge()
	signature, _ := eth.SignEthMessage(privKey, message)

	req := eth.ChallengeRequest{
		ExpectedPrefix: "Civil Test",
		GracePeriod:    100,
		InputChallenge: message,
		InputAddress:   sigAddr.Hex(),
		Signature:      signature,
	}

	err := eth.VerifyEthChallengeAndSignature(req)
	if err != nil {
		t.Errorf("verify eth challenge and signature threw error: err: %v", err)
	}
}

func TestVerifyEthChallengeAndSignatureGracePeriod(t *testing.T) {
	privKey, _ := crypto.GenerateKey()
	pubKeyBys := crypto.FromECDSAPub(&privKey.PublicKey)
	sigAddr := common.BytesToAddress(crypto.Keccak256(pubKeyBys[1:])[12:])
	message := LoginChallenge()
	signature, _ := eth.SignEthMessage(privKey, message)

	req := eth.ChallengeRequest{
		ExpectedPrefix: "Civil Test",
		GracePeriod:    3,
		InputChallenge: message,
		InputAddress:   sigAddr.String(),
		Signature:      signature,
	}

	time.Sleep(4 * time.Second)

	err := eth.VerifyEthChallengeAndSignature(req)
	if err == nil {
		t.Errorf("verify eth challenge should have expired grace period")
	}
}

func TestVerifySignature(t *testing.T) {
	address := "0x5385A3a9a1468b7D900A93E6f21E903E30928764"
	message := "Log in to Civil @ 2019-09-25T17:01:19.545Z"
	signature := "0x0808658f881758de3ada6d9e072f2c9b89d8aad580885af19e21268d012d528733950be423d4a726c03c57cf5523d88aaac57c6a96526c3cbaa938592a15db7c1b"

	var result, err = eth.VerifyEthSignature(address, message, signature)

	if err != nil {
		t.Fatalf("error thrown: %s", err)
	}
	if !result {
		t.Errorf("signature was not verified")
	}
}

func TestVerifyInvalidSignature(t *testing.T) {
	address := "0x7c342E040D73639FA20b8e4f539BA6A29319DcCc"
	message := "Civil Test @ 2018-01-09T20:08:57Z"

	// Sign this message with a different key; the verification should fail
	privKey, _ := crypto.GenerateKey()
	signature, _ := eth.SignEthMessage(privKey, message)

	result, err := eth.VerifyEthSignature(address, message, signature)
	if err != nil {
		t.Errorf("error should not have been thrown: %s", err)
	}
	if result {
		t.Errorf("signature should not have been verified")
	}
}

func TestVerifyChallengeMalformed(t *testing.T) {

	const prefix = "Civil Test"
	const challenge = "Invalid prefix @ 2018-01-09T20:08:57Z"

	var err = eth.VerifyEthChallenge(prefix, 100, challenge)

	if err == nil {
		t.Fatalf("challenge was verified when it should not have been")
	} else if err.Error() != "challenge does not start with `Civil Test`" {
		t.Fatalf("did not expect this error message: '%v'", err.Error())
	}
}

func TestVerifyChallengeExpired(t *testing.T) {

	const prefix = "Civil Test"
	const challenge = "Civil Test @ 2018-01-09T20:08:57Z"

	var err = eth.VerifyEthChallenge(prefix, 100, challenge)

	if err == nil {
		t.Fatalf("challenge was verified when it should not have been")
	} else if err.Error() != "expired" {
		t.Fatalf("did not expect this error message: " + err.Error())
	}
}

func LoginChallenge() string {
	return "Civil Test @ " + time.Now().Format(time.RFC3339)
}

func TestVerifyChallengeValid(t *testing.T) {

	const prefix = "Civil Test"
	challenge := LoginChallenge()

	var err = eth.VerifyEthChallenge(prefix, 100, challenge)

	if err != nil {
		t.Fatalf("error thrown: %s", err)
	}
}

func TestSignEthMessage(t *testing.T) {
	message := "Civil Test @ 2018-01-09T20:08:57Z"

	acct, _ := eth.MakeAccount()
	pk := acct.Key

	signature, err := eth.SignEthMessage(pk, message)
	if err != nil {
		t.Errorf("did not expect error signing: err: %v", err)
	}

	ethAddress := eth.GetEthAddressFromPrivateKey(pk)

	// Verify that the signature is valid with same message
	result, err := eth.VerifyEthSignature(ethAddress.Hex(), message, signature)

	if err != nil {
		t.Fatalf("error thrown: %s", err)
	}
	if !result {
		t.Errorf("signature was not verified for generated signature")
	}

	// Test
}

func TestVerifyEthSignatureWithPubkey(t *testing.T) {
	message := "Civil Test @ 2018-01-09T20:08:57Z"

	acct, _ := eth.MakeAccount()
	pk := acct.Key

	signature, err := eth.SignEthMessage(pk, message)
	if err != nil {
		t.Errorf("did not expect error signing: err: %v", err)
	}

	// Ensure proper conversion between hex store format and bytes to pub key
	// Test with standard package, does not append 0x to hex strings as in go-ethereum
	// hexutils
	pubKeyBys := crypto.FromECDSAPub(&pk.PublicKey)
	pubKeyHex := hex.EncodeToString(pubKeyBys)
	pubKeyBys, _ = hex.DecodeString(pubKeyHex)
	pubKey, _ := crypto.UnmarshalPubkey(pubKeyBys)

	// Verify the signature using the public key
	result, err := eth.VerifyEthSignatureWithPubkey(*pubKey, message, signature)
	if err != nil {
		t.Fatalf("Should not have failed to verify")
	}
	if !result {
		t.Errorf("should have verified the generated signature")
	}

}

func TestVerifyEthSignatureWithPubkeyHex(t *testing.T) {
	message := "Civil Test @ 2018-01-09T20:08:57Z"
	signature := "0x9ff69ebcb961ab7617733a0f75efc5d419cea57067e4de0455d53da9779e5ec665cc46e67b32d698e9eab3653dd9e9ed735f9b81dce3c7d599c4391aa064f97001"
	pubKeyHex := "04f621285c7e14fc5eabecb65a1796afdc1a3f8f23926b7faa853b13eacb6d212c527d2eefd0d7e82bc1a59dc99220797c6034751eb9dadf23fc9fc6ed1274fca0"

	pubKeyBys, _ := hex.DecodeString(pubKeyHex)
	pubKey, _ := crypto.UnmarshalPubkey(pubKeyBys)

	// Verify the signature using the public key
	result, err := eth.VerifyEthSignatureWithPubkey(*pubKey, message, signature)
	if err != nil {
		t.Errorf("Should not have failed to verify")
	}
	if !result {
		t.Errorf("should have verified the generated signature")
	}

}
