package socks5

import (
	"testing"
)

func TestRadiusCfg(t *testing.T) {
	creds := RadiusCfg{
		Server: "foo",
		Secret: "bar",
	}

	if creds.getServer() != "foo" {
		t.Fatalf("expect valid")
	}
}
