package main

import (
	"testing"

	"github.com/beevik/ntp"
)

func TestNtp1(t *testing.T) {
	_, err := ntp.Time("bla.bla")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestNtp2(t *testing.T) {
	_, err := ntp.Time("ntp1.ntp-servers.net")
	if err != nil {
		t.Error("Expected error")
	}
}
