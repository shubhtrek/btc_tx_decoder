package decoder

import (
	"encoding/binary"
	"fmt"
)

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
		Inputs:  make([]TxInput, 0),
		Outputs: make([]TxOutput, 0),
	}

	for i := uint64(0); i < inputCount; i++ {
		prevTxID, err := r.read(32)
		if err != nil {
			return nil, err
		}

		prevIndex, err := r.ReadUint32()
		if err != nil {
			return nil, err
		}

		scriptLen, err := r.ReadVarInt()
		if err != nil {
			return nil, err
		}

		scriptsig, err := r.read(int(scriptLen))
		if err != nil {
			return nil, err
		}

		sequence, err := r.ReadUint32()
		if err != nil {
			return nil, err
		}

		input := TxInput{
			PrevTxID:  prevTxID,
			PrevIndex: prevIndex,
			ScriptSig: scriptsig,
			Sequence:  sequence,
		}

		tx.Inputs = append(tx.Inputs, input)
	}

	outputCount, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	for i := uint64(0); i < outputCount; i++ {
		valueByte, err := r.read(8)

		if err != nil {
			return nil, err
		}
		value := binary.LittleEndian.Uint64(valueByte)

		scriptLen, err := r.ReadVarInt()
		if err != nil {
			return nil, err
		}

		scriptPubkey, err := r.read(int(scriptLen))
		if err != nil {
			return nil, err
		}

		output := TxOutput{
			Value:        value,
			ScriptPubkey: scriptPubkey,
		}

		tx.Outputs = append(tx.Outputs, output)
	}

	lockTime, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	tx.LockTime = lockTime
	return tx, nil
}

func PrettyPrint(tx *Transaction) {
	fmt.Println("----- Bitcoin Transaction -----")
	fmt.Println("Version:", tx.Version)
	fmt.Println("Inputs:", len(tx.Inputs))
	fmt.Println("Outputs:", len(tx.Outputs))
	fmt.Println("LockTime:", tx.LockTime)
	fmt.Println()

	for i, in := range tx.Inputs {
		fmt.Println("Input", i)
		fmt.Println("  PrevTxID:", FormatTXID(in.PrevTxID))
		fmt.Println("  PrevIndex:", in.PrevIndex)
		fmt.Println("  ScriptSig:", fmt.Sprintf("%x", in.ScriptSig))
		fmt.Println("  Sequence:", in.Sequence)
		fmt.Println()
	}

	for i, out := range tx.Outputs {
		fmt.Println("Output", i)
		fmt.Println("  Value (sats):", out.Value)
		fmt.Println("  ScriptPubKey:", fmt.Sprintf("%x", out.ScriptPubkey))
		fmt.Println()
	}
}
