package decoder

import (
	"crypto/sha256"
	"encoding/hex"
)

// DoubleSHA256 = sha256(sha256(data))
func DoubleSHA256(data []byte) []byte {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// TXID (legacy hash - no witness)
func CalculateTXID(raw []byte) string {
	hash := DoubleSHA256(raw)
	reversed := ReverseBytes(hash) // Bitcoin display format
	return hex.EncodeToString(reversed)
}

// WTXID (segwit hash - with witness)
func CalculateWTXID(raw []byte) string {
	hash := DoubleSHA256(raw)
	reversed := ReverseBytes(hash)
	return hex.EncodeToString(reversed)
}

