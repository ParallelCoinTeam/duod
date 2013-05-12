package btc

import (
	"errors"
	"encoding/binary"
)

type Block struct {
	Raw []byte
	Hash *Uint256
	Txs []*Tx

	Version uint32
	Parent []byte
	MerkleRoot []byte
	BlockTime uint32
	Bits uint32

	// if the block is trusted, we do not check signatures and some other things...
	Trusted bool
}


func NewBlock(data []byte) (*Block, error) {
	if len(data)<81 {
		return nil, errors.New("Block too short")
	}

	var bl Block
	bl.Hash = NewSha2Hash(data[:80])
	bl.Raw = data
	bl.Version = binary.LittleEndian.Uint32(data[0:4])
	bl.Parent = data[4:36]
	bl.MerkleRoot = data[36:68]
	bl.BlockTime = binary.LittleEndian.Uint32(data[68:72])
	bl.Bits = binary.LittleEndian.Uint32(data[72:76])
	return &bl, nil
}


// Parses block's transactions and adds them to the structure, calculating hashes BTW.
// It would be more elegant to use bytes.Reader here, but this solution is ~20% faster.
func (bl *Block) BuildTxList() (e error) {
	offs := int(80)
	txcnt, n := VLen(bl.Raw[offs:])
	offs += n
	bl.Txs = make([]*Tx, txcnt)

	for i:=0; i<useThreads; i++ {
		taskDone <- false
	}

	for i:=0; i<int(txcnt); i++ {
		bl.Txs[i], n = NewTx(bl.Raw[offs:])
		if bl.Txs[i] == nil {
			e = errors.New("NewTx failed")
			break
		}
		bl.Txs[i].Size = uint32(n)
		_ = <- taskDone // wait if we have too many threads already
		go func(h **Uint256, b []byte) {
			*h = NewSha2Hash(b)
			taskDone <- true
		}(&bl.Txs[i].Hash, bl.Raw[offs:offs+n])
		offs += n
	}

	// Wait for pending hashing to finish...
	for i:=0; i<useThreads; i++ {
		_ = <- taskDone
	}
	return
}


func GetBlockReward(height uint32) (uint64) {
	return 50e8 >> (height/210000)
}
