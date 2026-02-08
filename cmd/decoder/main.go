package main

import (
	"encoding/hex"
	"fmt"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	fmt.Println("Bitcoin Transaction Decoder")

	// raw transaction hex (sample bitcoin transaction)

	hexTx := "0100000001e1c7a1d2f3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f001122334455667788990000000000ffffffff0100f2052a010000000000000000"

	raw, err := hex.DecodeString(hexTx)
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}

	tx, err := decoder.Decode(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	decoder.PrettyPrint(tx)
	decoder.PrintSummary(tx)


	fmt.Println("Transaction Decoder Done")
}
