package ssha256_test

import (
	"testing"

	"github.com/jsimonetti/pwscheme/ssha256"
)

func TestValidPassword(t *testing.T) {
	pass := "test123"
	hash := "{SSHA256}czO44OTV17PcF1cRxWrLZLy9xHd7CWyVYplr1rOhuMlx/7IK"

	if res, err := ssha256.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Valid password fails validation: %s", err)
	}
}

func TestInValidPassword(t *testing.T) {
	pass := "test12"
	hash := "{SSHA256}czO44OTV17PcF1cRxWrLZLy9xHd7CWyVYplr1rOhuMlx/7IK"

	if res, err := ssha256.Validate(pass, hash); res != false {
		t.Errorf("Invalid password passes validation: %s", err)
	}
}

func TestGenerate4(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha256.Generate(pass, 4); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha256.Validate(pass, hash); err != nil && res != false {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}

func TestGenerate8(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha256.Generate(pass, 8); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha256.Validate(pass, hash); err != nil && res != false {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}
