package decoder

type Transaction struct{
	Version int64
	Inputs []TxInput
	Outputs []TxOutput
	LockTime uint32
}

type TxInput struct{
	PrevTxId []byte
	PrevIndex uint32
	ScriptSign []byte
	Sequence uint32

}

type TxOutput struct{
	Value uint64
	PubKey []byte

}