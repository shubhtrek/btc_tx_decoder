package decoder

type Transaction struct {
	Version  uint32      `json:"version"`
	Inputs   []TxInput  `json:"input"`
	Outputs  []TxOutput `json:"output"`
	LockTime uint32     `json:"locktime"`
}

type TxInput struct {
	PrevTxId   []byte `json:"prevtxid"`
	PrevIndex  uint32 `json:"previndex"`
	ScriptSign []byte `json:"scriptsign"`
	Sequence   uint32 `json:"sequence"`
}

type TxOutput struct {
	Value  uint64 `json:"value"`
	PubKey []byte `json:"pubkey"`
}
