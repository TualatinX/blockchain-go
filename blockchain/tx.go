package blockchain

type TxOutput struct {
	// Value would be representative of the amount of coins in a transaction
	Value int

	// The Pubkey is needed to "unlock" any coins within an Output. This indicated that YOU are the one that sent it.
	// You are indentifiable by your PubKey
	// PubKey in this iteration will be very straightforward, however in an actual application this is a more complex algorithm
	PubKey string
}

//TxInput is representative of a reference to a previous TxOutput
type TxInput struct {
	// ID will find the Transaction that a specific output is inside of
	ID []byte

	// Out will be the index of the specific output we found within a transaction.
	// For example if a transaction has 4 outputs, we can use this "Out" field to specify which output we are looking for
	Out int

	// This would be a script that adds data to an outputs' PubKey. However for this tutorial the Sig will be indentical to the PubKey.
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
