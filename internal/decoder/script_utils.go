package decoder

func DetectScriptType(script []byte) string{

	if len(script) == 25 &&
		script[0] == 0x76&&
		script[1] == 0xa9 &&
		script[2] == 0x14 &&
		script[23] == 0x88 &&
		script[24] == 0xac {
		return "P2PKH"
	}

	// P2SH: a9 14 <20 bytes> 87
	if len(script) == 23 &&
		script[0] == 0xa9 &&
		script[1] == 0x14 &&
		script[22] == 0x87 {
		return "P2SH"
	}

	// P2WPKH: 00 14 <20 bytes>
	if len(script) == 22 &&
		script[0] == 0x00 &&
		script[1] == 0x14 {
		return "P2WPKH (SegWit)"
	}
	// P2WSH: 00 20 <32 bytes>
	if len(script) == 34 &&
		script[0] == 0x00 &&
		script[1] == 0x20 {
		return "P2WSH (SegWit)"
	}

	return "Unknown"
}