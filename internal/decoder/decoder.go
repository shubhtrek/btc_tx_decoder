package decoder

func Decoder(raw []byte) (*Transaction, error) {
	r := NewReader(raw)

	version, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}

	inputCount, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	tx := &Transaction{
		Version: version,
		Inputs:  make([]TxInput, inputCount),
	}

	return tx, nil
}
