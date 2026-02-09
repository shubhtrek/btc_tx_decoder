package decoder

// ExtractAddressHash extracts the hash part from scriptPubKey
func ExtractAddressHash(script []byte) []byte {

	// P2PKH: 76 a9 14 <20 bytes> 88 ac
	if len(script) == 25 &&
		script[0] == 0x76 &&
		script[1] == 0xa9 &&
		script[2] == 0x14 {
		return script[3:23]
	}

	// P2SH: a9 14 <20 bytes> 87
	if len(script) == 23 &&
		script[0] == 0xa9 &&
		script[1] == 0x14 {
		return script[2:22]
	}

	// P2WPKH: 00 14 <20 bytes>
	if len(script) == 22 &&
		script[0] == 0x00 &&
		script[1] == 0x14 {
		return script[2:22]
	}

	// P2WSH: 00 20 <32 bytes>
	if len(script) == 34 &&
		script[0] == 0x00 &&
		script[1] == 0x20 {
		return script[2:34]
	}

	return nil
}
