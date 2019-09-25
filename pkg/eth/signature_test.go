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
	sigAddr := common.BytesToAddress(crypto.Keccak256(pubKeyBys[1:])[12:])
	message := LoginChallenge()
	signature, _ := eth.SignEthMessage(privKey, message)

	req := eth.ChallengeRequest{
		ExpectedPrefix: "Civil Test",
		GracePeriod:    100,
		InputChallenge: message,
		InputAddress:   sigAddr.String(),
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
	address := "0x7c342E040D73639FA20b8e4f539BA6A29319DcCc"
	message := "Civil Test @ 2018-01-09T20:08:57Z"
	signature := "0x120fbe013f535b5ace9590b7e902adc3aaf72cb52349a7b3975bb0ffb3992b8b469afe5b590cc6342afc78d6119a06f12ca5fd2c431468f8dfe619c739c1f96400"

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
	signature := "0xfea80f27862777116d2d5557aaa5a8887bf411a70e3c3cd377368128f88546ce2f59ffbd8263d64d328a324f21ab5dfe6516467bc2129b35012d9ca8bb39f0eb01"
	pubKeyHex := "044f4a35cb48550342ac99b5be887a1e90b3d0ea30010e74ab77faa077a935f65e5dcc9ab02755f982e85e06a52679ccae6221d631c89dfaba6359f305444b93ef"

	// Ensure proper conversion between hex store format and bytes to pub key
	// Test with standard package, does not append 0x to hex strings as in go-ethereum
	// hexutils
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
