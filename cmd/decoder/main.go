package main

import (
	"fmt"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	fmt.Println("Bitcoin Transaction Decoder")

	raw := []byte{}
	_, err := decoder.Decoder(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction Decoder Done")
}
