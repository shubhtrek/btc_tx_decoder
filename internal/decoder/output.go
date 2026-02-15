package decoder

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(tx *Transaction) {
	b, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println(string(b))
}

func PrintInputs(tx *Transaction) {
	fmt.Println("---- INPUTS ----")
	for i, in := range tx.Inputs {
		fmt.Println("Input", i)
		fmt.Println(" PrevTxID:", FormatTXID(in.PrevTxID))
		fmt.Println(" Index:", in.PrevIndex)
		fmt.Println(" ScriptSig:", fmt.Sprintf("%x", in.ScriptSig))
		fmt.Println(" Sequence:", in.Sequence)
		fmt.Println()
	}
}

func PrintOutputs(tx *Transaction) {
	fmt.Println("---- OUTPUTS ----")
	for i, out := range tx.Outputs {
		fmt.Println("Output", i)
		fmt.Println(" Value:", out.Value)
		fmt.Println(" ScriptPubKey:", fmt.Sprintf("%x", out.ScriptPubkey))
		fmt.Println()
	}
}
