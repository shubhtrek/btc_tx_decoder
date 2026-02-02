package decoder

import "fmt"

func Decoder(raw []byte) (*Transaction, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("Empty Tx")
	}

	tx := &Transaction{}

	return tx ,nil
}