package main

import (
	"fmt"

	"github.com/shubhtrek/btc_tx_decoder/internal/decoder"
)

func main() {

	fmt.Println("Bitcoin Transaction Decoder")

	raw := []byte{
			0x01, 0x00, 0x00, 0x00, // version = 1
	0x01, // input count = 1
	}
	_, err := decoder.Decode(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction Decoder Done")
}
