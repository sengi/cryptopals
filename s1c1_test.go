package cryptopals

import (
	"encoding/base64"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual, err := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}
	s, err := base64.StdEncoding.DecodeString(actual)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
	if expected != actual {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
