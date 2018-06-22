package twofactor

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letters = "1234567890!@#$%^&*()abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString() string {
	b := make([]byte, rand.Intn(len(letters)))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return base32.StdEncoding.EncodeToString(b)
}

func TestPadding(t *testing.T) {
	for i := 0; i < 300; i++ {
		b := randString()
		origEncoding := string(b)
		modEncoding := strings.Replace(string(b), "=", "", -1)
		str, err := base32.StdEncoding.DecodeString(origEncoding)
		if err != nil {
			fmt.Println("Can't decode: ", string(b))
			t.FailNow()
		}

		paddedEncoding := Pad(modEncoding)
		if origEncoding != paddedEncoding {
			fmt.Println("Padding failed:")
			fmt.Printf("Expected: '%s'", origEncoding)
			fmt.Printf("Got: '%s'", paddedEncoding)
			t.FailNow()
		} else {
			mstr, err := base32.StdEncoding.DecodeString(paddedEncoding)
			if err != nil {
				fmt.Println("Can't decode: ", paddedEncoding)
				t.FailNow()
			}

			if string(mstr) != string(str) {
				fmt.Println("Re-padding failed:")
				fmt.Printf("Expected: '%s'", str)
				fmt.Printf("Got: '%s'", mstr)
				t.FailNow()
			}
		}
	}
}
