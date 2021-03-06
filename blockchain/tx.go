package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/TualatinX/blockchain-go/wallet"
)

type TxOutput struct {
	// Value would be representative of the amount of coins in a transaction
	Value int

	PubKeyHash []byte
}

type TxOutputs struct {
	Outputs []TxOutput
}

//TxInput is representative of a reference to a previous TxOutput
type TxInput struct {
	// ID will find the Transaction that a specific output is inside of
	ID []byte

	// Out will be the index of the specific output we found within a transaction.
	// For example if a transaction has 4 outputs, we can use this "Out" field to specify which output we are looking for
	Out int

	Signature []byte
	PubKey    []byte
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))
	return txo
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)
	return bytes.Equal(lockingHash, pubKeyHash)
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-wallet.ChecksumLength]
	out.PubKeyHash = pubKeyHash
}
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)

}

func (outs *TxOutputs) Serialize() []byte {
	var content bytes.Buffer

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(outs)
	Handle(err)

	return content.Bytes()
}

func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&outputs)
	Handle(err)
	return outputs
}
