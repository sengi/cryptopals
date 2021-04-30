package cryptopals

import "bytes"

func fitness(v []byte) float32 {
	f := float32(0)
	// http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
	fr := map[byte]float32{'e': 12.02, 't': 9.1, 'a': 8.12, 'o': 7.68, 'i': 7.31, 'n': 6.95}
	for _, b := range bytes.ToLower(v) {
		if b < ' ' || b > '~' {
			f -= 12
		} else {
			f += fr[b]
		}
	}
	return f
}

func SingleByteXOR(v []byte) (key byte) {
	var max_score float32
	for i := 0; i <= 0xff; i++ {
		decrypt := FixedXOR(v, []byte{byte(i)})
		score := fitness(decrypt)
		if score > max_score {
			max_score = score
			key = byte(i)
		}
	}
	return
}
