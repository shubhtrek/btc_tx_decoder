package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	// ---- flags ----
	summaryOnly := flag.Bool("summary", false, "Show transaction summary only")
	jsonOut := flag.Bool("json", false, "Show transaction as JSON")
	showInputs := flag.Bool("inputs", false, "Show inputs only")
	showOutputs := flag.Bool("outputs", false, "Show outputs only")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage:")
		fmt.Println("  btc_tx_decoder <tx_hex_file> [--summary] [--json] [--inputs] [--outputs]")
		return
	}

	filePath := flag.Arg(0)

	fmt.Println("Bitcoin Transaction Decoder")

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

	// ---- output modes ----
	if *jsonOut {
		decoder.PrintJSON(tx)
		return
	}

	if *summaryOnly {
		decoder.PrintSummary(tx)
		return
	}

	if *showInputs {
		decoder.PrintInputs(tx)
		return
	}

	if *showOutputs {
		decoder.PrintOutputs(tx)
		return
	}

	// output
	decoder.PrettyPrint(tx)
	decoder.PrintSummary(tx)

	fmt.Println("Transaction Decoder Done")
}
