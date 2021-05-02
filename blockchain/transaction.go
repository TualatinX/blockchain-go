package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

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

// CoinbaseTx is the function that will run when someone on a node succesfully "mines" a block. The reward inside as it were.
func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}
	// Since this is the "first" transaction of the block, it has no previous output to reference.
	// This means that we initialize it with no ID, and it's OutputIndex is -1
	txIn := TxInput{[]byte{}, -1, data}
	// txOut will represent the amount of tokens(reward) given to the person(toAddress) that executed CoinbaseTx
	txOut := TxOutput{reward, toAddress} // You can see it follows {value, PubKey}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx

}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (tx *Transaction) IsCoinbase() bool {
	// This checks a transaction and will only return true if it is a newly minted "coin"
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// Find Spendable Outputs
// Check if we have enough money to send the amount that we are asking
// If we do, make inputs that point to the outputs we are spending
// If there is any leftover money, make new outputs from the difference.
// Initialize a new transaction with all the new inputs and outputs we made
// Set a new ID, and return it.
func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: Not enough funds!")
	}
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
