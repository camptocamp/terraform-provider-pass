/*
 * goprogressbar
 *     Copyright (c) 2016-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package goprogressbar

import (
	"bytes"
	"testing"
)

func TestCursorMovement(t *testing.T) {
	buf := &bytes.Buffer{}
	Stdout = buf

	moveCursorUp(5)
	if buf.String() != "\033[5A" {
		t.Errorf("Unexpected cursor up movement behaviour")
	}
	buf.Reset()

	moveCursorDown(5)
	if buf.String() != "\033[5B" {
		t.Errorf("Unexpected cursor down movement behaviour")
	}
	buf.Reset()
}
