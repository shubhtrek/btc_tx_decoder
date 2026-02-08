package decoder

import "encoding/hex"

func ReverseBytes(b []byte) []byte {
	out := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		out[i] = b[len(b)-1-i]
	}
	return out
}

func FormatTXID(txid []byte) string {
	reversed := ReverseBytes(txid)
	return hex.EncodeToString(reversed)
}