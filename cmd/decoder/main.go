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

	// REAL Bitcoin transaction (testnet example)

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/decoder <tx_hex_file>")
		return
	}

	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	hexTx := strings.TrimSpace(string(data))

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
