package main

import (
	"encoding/hex"
	"fmt"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	fmt.Println("Bitcoin Transaction Decoder")

	// REAL Bitcoin transaction (testnet example)
	hexTx := "02000000000101b1c2d3e4f5a6978899aabbccddeeff00112233445566778899aabbccddeeff0000000000ffffffff02e8030000000000001600144f3c1a2b3d4e5f60718293a4b5c6d7e8f901234567880000000000000000160014abcdefabcdefabcdefabcdefabcdefabcdefabcd02473044022011223344556677889900aabbccddeeff00112233445566778899aabbccddeeff022011223344556677889900aabbccddeeff00112233445566778899aabbccddeeff012102abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef"

	raw, err := hex.DecodeString(hexTx)
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}

	tx, err := decoder.Decode(raw)
	if err != nil {
		fmt.Println("Decode failed ❌")
		fmt.Println("Reason:", err)
		return
	}

	decoder.PrettyPrint(tx)
	decoder.PrintSummary(tx)

	fmt.Println("Transaction Decoder Done")
}
