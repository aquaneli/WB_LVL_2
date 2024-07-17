package main

import (
	"testing"

	"github.com/beevik/ntp"
)

func Test1(t *testing.T) {
	_, err := ntp.Time("bla.bla")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
