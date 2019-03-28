package md5crypt_test

import (
	"testing"

	"github.com/jsimonetti/pwscheme/md5crypt"
)

func TestValidPassword(t *testing.T) {
	pass := "test123"
	hash := "{MD5-CRYPT}$1$UNq3KKXM$hsrKHTk9BaYGZwafpr4K80"

	if res, err := md5crypt.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Valid password fails validation: %s", err)
	}
}

func TestInValidPassword(t *testing.T) {
	pass := "test12"
	hash := "{MD5-CRYPT}$1$UNq3KKXM$hsrKHTk9BaYGZwafpr4K80"

	if res, err := md5crypt.Validate(pass, hash); res != false {
		t.Errorf("Invalid password passes validation: %s", err)
	}
}

func TestGenerate4(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = md5crypt.Generate(pass, 4); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = md5crypt.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}

func TestGenerate8(t *testing.T) {
	pass := "test123"
	var hash string
	var err error
	var res bool

	if hash, err = md5crypt.Generate(pass, 8); err != nil {
		t.Errorf("Generate password fails: %s", err)
		return
	}

	if res, err = md5crypt.Validate(pass, hash); err != nil || res != true {
		t.Errorf("Generated hash can not be validated: %s", err)
	}
}

func TestGenerate0(t *testing.T) {
	if _, err := md5crypt.Generate("", 0); err != md5crypt.ErrSaltLengthInCorrect {
		t.Errorf("Generated hash with too short salt did not fail")
	}

}

func TestGenerate9(t *testing.T) {
	if _, err := md5crypt.Generate("", 9); err != md5crypt.ErrSaltLengthInCorrect {
		t.Errorf("Generated hash with too long salt did not fail")
	}

}
