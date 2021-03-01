package wen

import "testing"

func TestVersionParse(t *testing.T) {

	v := VersionParse("0.1.1", 4)
	t.Log(v)
}
