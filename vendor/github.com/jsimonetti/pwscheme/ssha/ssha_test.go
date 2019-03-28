package ssha_test

import (
	"testing"

	"github.com/jsimonetti/pwscheme/ssha"
)

func TestValidPassword(t *testing.T) {
	pass := "test123"
	hash := "{SSHA}JFZFs0oHzxbMwkSJmYVeI8MnTDy/276a"

	if res, err := ssha.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Valid password fails validation: %s", err)
	}
}

func TestInValidPassword(t *testing.T) {
	pass := "test12"
	hash := "{SSHA}JFZFs0oHzxbMwkSJmYVeI8MnTDy/276a"

	if res, err := ssha.Validate(pass, hash); res != false {
		t.Errorf("Invalid password passes validation: %s", err)
	}
}

func TestGenerate4(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha.Generate(pass, 4); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}

func TestGenerate8(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha.Generate(pass, 8); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}
