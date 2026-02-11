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

	// ---- SegWit detection ----
	isSegWit := false

	marker, err := r.read(1)
	if err != nil {
		return nil, err
	}

	flag, err := r.read(1)
	if err != nil {
		return nil, err
	}

	if marker[0] == 0x00 && flag[0] == 0x01 {
		isSegWit = true
	} else {
		// not segwit → rewind reader position
		r.pos -= 2
	}

	inputCount, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	tx := &Transaction{
		Version:  version,
		IsSegWit: isSegWit,
		Inputs:   make([]TxInput, 0),
		Outputs:  make([]TxOutput, 0),
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

	// ---- Witness data (SegWit only) ----
	if tx.IsSegWit {
		for i := 0; i < len(tx.Inputs); i++ {

			witnessCount, err := r.ReadVarInt()
			if err != nil {
				return nil, err
			}

			var witnessStack [][]byte

			for j := uint64(0); j < witnessCount; j++ {
				itemLen, err := r.ReadVarInt()
				if err != nil {
					return nil, err
				}

				item, err := r.read(int(itemLen))
				if err != nil {
					return nil, err
				}

				witnessStack = append(witnessStack, item)
			}

			tx.Inputs[i].Witness = witnessStack
		}
	}

	lockTime, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	tx.LockTime = lockTime

	// ---- Hashes ----
	tx.TXID = CalculateTXID(raw)
	if tx.IsSegWit {
		tx.WTXID = CalculateWTXID(raw)
	} else {
		tx.WTXID = tx.TXID
	}

	return tx, nil

}

func PrettyPrint(tx *Transaction) {
	fmt.Println("========== TX DECODE ==========")
	fmt.Println("----- Bitcoin Transaction -----")
	fmt.Println("Version:", tx.Version)
	fmt.Println("SegWit:", tx.IsSegWit)
	fmt.Println("TXID:", tx.TXID)
	fmt.Println("WTXID:", tx.WTXID)
	fmt.Println("Inputs:", len(tx.Inputs))
	fmt.Println("Outputs:", len(tx.Outputs))
	fmt.Println("LockTime:", tx.LockTime)
	fmt.Println()
	fmt.Println("========== END TX =============")


	for i, in := range tx.Inputs {
		fmt.Println("Input", i)
		fmt.Println("  PrevTxID:", FormatTXID(in.PrevTxID))
		fmt.Println("  PrevIndex:", in.PrevIndex)
		fmt.Println("  ScriptSig (hex):", fmt.Sprintf("%x", in.ScriptSig))
		fmt.Println("  ScriptSig (len):", len(in.ScriptSig), "bytes")
		fmt.Println("  ScriptSig Type:", DetectScriptType(in.ScriptSig))

		fmt.Println("  Sequence:", in.Sequence)
		fmt.Println()

		if tx.IsSegWit && len(in.Witness) > 0 {
			fmt.Println("  Witness:")
			for wi, w := range in.Witness {
				fmt.Println("   -", wi, ":", fmt.Sprintf("%x", w))
			}
		}

	}

	for i, out := range tx.Outputs {
		fmt.Println("Output", i)
		fmt.Println("  Value (sats):", out.Value)
		fmt.Println("  ScriptPubKey (hex):", fmt.Sprintf("%x", out.ScriptPubkey))
		fmt.Println("  ScriptPubKey (len):", len(out.ScriptPubkey), "bytes")
		fmt.Println("  ScriptPubKey Type:", DetectScriptType(out.ScriptPubkey))
		hash := ExtractAddressHash(out.ScriptPubkey)
		if hash != nil {
			fmt.Println("  Address Hash:", fmt.Sprintf("%x", hash))

			if DetectScriptType(out.ScriptPubkey) == "P2PKH" {
				addr := P2PKHAddress(hash)
				fmt.Println("  Bitcoin Address:", addr)
			}
		}

		fmt.Println()
	}
}

func PrintSummary(tx *Transaction) {
	fmt.Println("----Tx Summary----")
	fmt.Println("Version:", tx.Version)
	fmt.Println("Inputs:", len(tx.Inputs))
	fmt.Println("OutPuts:", len(tx.Outputs))
	fmt.Println("LockTime:", tx.LockTime)

	var total uint64 = 0
	for _, out := range tx.Outputs {
		total += out.Value
	}

	fmt.Println("Total Output (sats):", total)
	fmt.Println("----------------")
	fmt.Println()
}
