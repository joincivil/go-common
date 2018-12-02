// Package time_test contains tests for the string strings

package strings_test

import (
	"testing"

	"github.com/joincivil/go-common/pkg/strings"
)

func TestIsValidEthAPIURL(t *testing.T) {
	if strings.IsValidEthAPIURL("thisisnotavalidurl") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if strings.IsValidEthAPIURL("http//thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if strings.IsValidEthAPIURL("http/thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if !strings.IsValidEthAPIURL("http://thisisvalid.co") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("http://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("https://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("https://thisisvalid.longtld") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("ws://thisisvalid.ether/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("wss://thisisvalid.com/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("wss://localhost/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("wss://localhost:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("wss://127.0.0.1/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !strings.IsValidEthAPIURL("wss://127.0.0.1:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
}

func TestIsValidContractAddress(t *testing.T) {
	if strings.IsValidContractAddress("") {
		t.Error("Should have failed on an empty contract address")
	}
	if strings.IsValidContractAddress("thisisnotavalidaddress") {
		t.Error("Should have failed on an invalid contract address")
	}
	if strings.IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55f") {
		t.Error("Should have failed on an invalid contract address")
	}
	if !strings.IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55") {
		t.Error("Should have not have failed on an valid contract address")
	}
}

func TestRandomHex(t *testing.T) {
	s, err := strings.RandomHexStr(32)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 64 {
		t.Errorf("Should have been a 64 char hex string: %v", len(s))
	}

	s, err = strings.RandomHexStr(10)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 20 {
		t.Errorf("Should have been a 20 char hex string: %v", len(s))
	}

	s, err = strings.RandomHexStr(0)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 0 {
		t.Errorf("Should have been a 0 char hex string: %v", len(s))
	}
}