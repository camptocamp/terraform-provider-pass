package xkcdpwgen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_GeneratePassword(t *testing.T) {
	g := NewGenerator()
	password := g.GeneratePasswordString()
	assert.Equal(t, 4, len(strings.Split(password, " ")), "not the right number of words for default settings")
}

func TestGenerator_GeneratePasswordWithCustomDelimiter(t *testing.T) {
	g := NewGenerator()
	g.SetDelimiter("%")
	password := g.GeneratePasswordString()
	assert.Equal(t, 4, len(strings.Split(password, "%")), "not the right number of words with custom delimiter")
}

func TestGenerator_GeneratePasswordWithCustomWordcount(t *testing.T) {
	g := NewGenerator()
	g.SetNumWords(5)
	password := g.GeneratePasswordString()
	assert.Equal(t, 5, len(strings.Split(password, " ")), "not the right number of words with custom wordcount")
}

func TestGenerator_GeneratePasswordWithCustomWordlist(t *testing.T) {
	g := NewGenerator()
	g.UseCustomWordlist([]string{"muh", "muh"})
	password := g.GeneratePasswordString()
	assert.Equal(t, "muh muh muh muh", password, "not the expected password for custom wordlist")
}

func TestGenerator_GeneratePasswordWithCapitalization(t *testing.T) {
	g := NewGenerator()
	g.UseCustomWordlist([]string{"muh", "muh"})
	g.SetCapitalize(true)
	password := g.GeneratePasswordString()
	assert.Equal(t, "Muh Muh Muh Muh", password, "not the correct capitalization")
}

func TestGenerator_SetLanguage(t *testing.T) {
	g := NewGenerator()
	err := g.UseLangWordlist("de")
	assert.Equal(t, nil, err, "Setting language to \"de\" should not raise an error")
	password := g.GeneratePasswordString()
	assert.Equal(t, 4, len(strings.Split(password, " ")), "not the right number of words for default settings")
}

func TestGenerator_SetInvalidLanguageRaisesError(t *testing.T) {
	g := NewGenerator()
	err := g.UseLangWordlist("muh")
	assert.EqualError(t, err, "language \"muh\" has no matching wordlist", "not the correct error message")
}
