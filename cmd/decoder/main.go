package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	fmt.Println("Bitcoin Transaction Decoder")

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/decoder <tx_hex_file>")
		return
	}

	filePath := os.Args[1]

	// read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	// clean hex
	hexTx := strings.TrimSpace(string(data))

	// hex -> bytes
	raw, err := hex.DecodeString(hexTx)
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}

	// decode tx
	tx, err := decoder.Decode(raw)
	if err != nil {
		fmt.Println("Decode failed ❌")
		fmt.Println("Reason:", err)
		return
	}

	// output
	decoder.PrettyPrint(tx)
	decoder.PrintSummary(tx)

	fmt.Println("Transaction Decoder Done")
}
