package ssha512_test

import (
	"testing"

	"github.com/jsimonetti/pwscheme/ssha512"
)

func TestValidPassword(t *testing.T) {
	pass := "test123"
	hash := "{SSHA512}xPUl/px+1cG55rUH4rzcwxdOIPSB2TingLpiJJumN2xyDWN4Ix1WQG3ihnvHaWUE8MYNkvMi5rf0C9NYixHsE6Yh59M="

	if res, err := ssha512.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Valid password fails validation: %s", err)
	}
}

func TestInValidPassword(t *testing.T) {
	pass := "test12"
	hash := "{SSHA512}xPUl/px+1cG55rUH4rzcwxdOIPSB2TingLpiJJumN2xyDWN4Ix1WQG3ihnvHaWUE8MYNkvMi5rf0C9NYixHsE6Yh59M="

	if res, err := ssha512.Validate(pass, hash); res != false {
		t.Errorf("Invalid password passes validation: %s", err)
	}
}

func TestGenerate4(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha512.Generate(pass, 4); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha512.Validate(pass, hash); err != nil && res != false {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}

func TestGenerate8(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = ssha512.Generate(pass, 8); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = ssha512.Validate(pass, hash); err != nil && res != false {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}
