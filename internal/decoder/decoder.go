package decoder

func Decode(raw []byte) (*Transaction, error) {
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
		Inputs:  make([]TxInput,0),
	}

	for i := uint64(0); i < inputCount; i++ {
		prevTxID,err  := r.read(32)
		if err != nil {
			return nil, err
		}

		prevIndex, err := r.ReadUint32()
		if err != nil {
			return nil, err
		}

		scriptLen,err := r.ReadVarInt()
		if err != nil {
			return nil, err
		}

		scriptsig,err := r.read(int(scriptLen))
		if err != nil {
			return nil, err
		}

		sequence,err := r.ReadUint32()
		if err != nil {
			return nil, err
		}

		input := TxInput{
			PrevTxID: prevTxID,
			PrevIndex: prevIndex,
			ScriptSig: scriptsig,
			Sequence: sequence,
		}

		tx.Inputs = append(tx.Inputs, input)
	}

	return tx, nil
}
