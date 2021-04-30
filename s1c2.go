package cryptopals

func FixedXOR(a, b []byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i%len(b)]
	}
	return res
}
