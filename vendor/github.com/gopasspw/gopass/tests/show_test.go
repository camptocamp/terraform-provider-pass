package tests

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	ts := newTester(t)
	defer ts.teardown()

	_, err := ts.run("show")
	assert.Error(t, err)

	ts.initStore()

	out, err := ts.run("show")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" show [name]\n", out)

	out, err = ts.run("show foo")
	assert.Error(t, err)
	assert.Equal(t, "\nError: failed to retrieve secret 'foo': Entry is not in the password store\n", out)

	ts.initSecrets("")

	_, err = ts.run("show foo")
	assert.NoError(t, err)
	_, err = ts.run("show -f foo")
	assert.NoError(t, err)
	_, err = ts.run("show foo -force")
	assert.NoError(t, err)

	out, err = ts.run("show fixed/secret")
	assert.NoError(t, err)
	assert.Contains(t, out, "safe content to display, you can force display with show -f.\nCopying password instead.")

	_, err = ts.run("config autoclip false")
	assert.NoError(t, err)
	_, err = ts.run("show fixed/secret")
	assert.Error(t, err)

	out, err = ts.run("show -f fixed/secret")
	assert.NoError(t, err)
	assert.Equal(t, "moar", out)

	out, err = ts.run("show fixed/twoliner")
	assert.NoError(t, err)
	assert.Equal(t, "more stuff", out)

	out, err = ts.run("show fixed/twoliner -f")
	assert.NoError(t, err)
	assert.Equal(t, "and\nmore stuff", out)

	out, err = ts.run("show fixed/twoliner -c")
	assert.NoError(t, err)
	assert.NotContains(t, out, "safe content to display")

	_, err = ts.run("config safecontent false")
	assert.NoError(t, err)

	out, err = ts.run("show fixed/twoliner")
	assert.NoError(t, err)
	assert.Equal(t, "and\nmore stuff", out)

	out, err = ts.run("show fixed/secret")
	assert.NoError(t, err)
	assert.Equal(t, "moar", out)

	out, err = ts.run("show --qr fixed/secret")
	assert.NoError(t, err)
	assert.Equal(t, "\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[40m  \x1b[0m\x1b[40m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\n\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m\x1b[47m  \x1b[0m", out)
}
