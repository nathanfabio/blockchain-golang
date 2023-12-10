package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//Static difficulty, but if it were to be used in the business, it would have to be an algorithm that is gradually incremented
const Difficulty = 10

type Proof struct {
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - Difficulty))
	pow := &Proof{b, target}
	return pow
}

func (pow *Proof) InitData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.Block.PrevHash, pow.Block.Data, ToHexadecimal(int64(nonce)), ToHexadecimal(int64(Difficulty))}, []byte{},)
	return data 
}

func ToHexadecimal(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}

func (pow *Proof) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
	}

	return nonce, hash[:]
}