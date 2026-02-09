package decoder

import (
	"crypto/sha256"
	"math/big"
)

// Base58 alphabet (Bitcoin)
var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// EncodeBase58 encodes bytes into base58 string
func EncodeBase58(input []byte) string {
	var result []byte

	x := new(big.Int).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	// handle leading zeros
	for _, b := range input {
		if b == 0x00 {
			result = append(result, base58Alphabet[0])
		} else {
			break
		}
	}

	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Base58CheckEncode = version + payload + checksum → base58
func Base58CheckEncode(version byte, payload []byte) string {
	data := append([]byte{version}, payload...)

	// checksum = first 4 bytes of double sha256
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:4]

	full := append(data, checksum...)
	return EncodeBase58(full)
}

// P2PKH address generator
func P2PKHAddress(hash []byte) string {
	// version 0x00 = mainnet P2PKH
	return Base58CheckEncode(0x00, hash)
}
