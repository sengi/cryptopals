package cryptopals

import (
	"testing"
)

func TestSingleByteXOR(t *testing.T) {
	v := hexToBytes(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expected := byte(88)
	actual := SingleByteXOR(v)
	if expected != actual {
		t.Fatalf("expected %v, got %v", expected, actual)
	}

	d := FixedXOR(v, []byte{expected})
	t.Logf("decrypt=%v", string(d))
}
