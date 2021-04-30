package cryptopals

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func hexToBytes(t *testing.T, s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestFixedXOR(t *testing.T) {
	a := hexToBytes(t, "1c0111001f010100061a024b53535009181c")
	b := hexToBytes(t, "686974207468652062756c6c277320657965")
	expected := hexToBytes(t, "746865206b696420646f6e277420706c6179")
	t.Logf("a=%v, b=%v, expected=%v", string(a), string(b), string(expected))
	actual := FixedXOR(a, b)
	if !bytes.Equal(expected, actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
