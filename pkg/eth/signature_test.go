package eth_test

import (
	"testing"
	"time"

	"github.com/joincivil/go-common/pkg/eth"
)

func TestVerifyEthChallengeAndSignature(t *testing.T) {
	req := eth.ChallengeRequest{
		ExpectedPrefix: "Civil Test",
		GracePeriod:    100,
		InputChallenge: LoginChallenge(),
		InputAddress:   "0xddB9e9957452d0E39A5E43Fd1AB4aE818aecC6aD",
		Signature:      "0x520c1f6a0f1f968db5aaa39c08055bf2bd33dc9162d0237423549d31e91b6c661aa171e475cca20e1f0347685eaca6a0e443ecf5de3f53fb88dbb006ade5fc001b",
	}

	err := eth.VerifyEthChallengeAndSignature(req)
	if err != nil {
		t.Errorf("verify eth challenge and signature threw error: err: %v", err)
	}
}

func TestVerifySignature(t *testing.T) {

	const address = "0xddB9e9957452d0E39A5E43Fd1AB4aE818aecC6aD"
	const message = "Civil Test @ 2018-01-09T20:08:57Z"
	const signature = "0x520c1f6a0f1f968db5aaa39c08055bf2bd33dc9162d0237423549d31e91b6c661aa171e475cca20e1f0347685eaca6a0e443ecf5de3f53fb88dbb006ade5fc001b"

	var result, err = eth.VerifyEthSignature(address, message, signature)

	if err != nil {
		t.Fatalf("error thrown: %s", err)
	}
	if !result {
		t.Errorf("signature was not verified")
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
