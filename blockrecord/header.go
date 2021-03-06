// Copyright (c) 2014-2016 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockrecord

import (
	"encoding/binary"
	"github.com/bitmark-inc/bitmarkd/blockdigest"
	"github.com/bitmark-inc/bitmarkd/difficulty"
	"github.com/bitmark-inc/bitmarkd/fault"
	"github.com/bitmark-inc/bitmarkd/merkle"
)

// packed records are just a byte slice
type PackedHeader []byte
type PackedBlock []byte

// block version
const (
	Version = 1
)

// maximum transactions in a block
// limited by uint16 field
const (
	MaximumTransactions = 10000
)

// byte sizes for various fields
const (
	VersionSize          = 2                   // Block version number
	TransactionCountSize = 2                   // Count of transactions
	NumberSize           = 8                   // This block's number
	PreviousBlockSize    = blockdigest.Length  // 256-bit Argon2d hash of the previous block header
	MerkleRootSize       = merkle.DigestLength // 256-bit SHA3 hash based on all of the transactions in the block
	TimestampSize        = 8                   // Current timestamp as seconds since 1970-01-01T00:00 UTC
	DifficultySize       = 8                   // Current target difficulty in compact format
	NonceSize            = 8                   // 64-bit number (starts at 0)
)

// offsets of the fields
const (
	versionOffset          = 0
	transactionCountOffset = versionOffset + VersionSize
	numberOffset           = transactionCountOffset + TransactionCountSize
	previousBlockOffset    = numberOffset + NumberSize
	merkleRootOffset       = previousBlockOffset + PreviousBlockSize
	timestampOffset        = merkleRootOffset + MerkleRootSize
	difficultyOffset       = timestampOffset + TimestampSize
	nonceOffset            = difficultyOffset + DifficultySize

	// the total size is exported
	TotalBlockSize = nonceOffset + NonceSize // total bytes in the header
)

// the unpacked header structure
// the types here must match Bitcoin header types
type Header struct {
	Version          uint16                 `json:"version"`
	TransactionCount uint16                 `json:"transactionCount"`
	Number           uint64                 `json:"number,string"`
	PreviousBlock    blockdigest.Digest     `json:"previousBlock"`
	MerkleRoot       merkle.Digest          `json:"merkleRoot"`
	Timestamp        uint64                 `json:"timestamp,string"`
	Difficulty       *difficulty.Difficulty `json:"difficulty"`
	Nonce            NonceType              `json:"nonce"`
}

// create a new header with attached difficulty item
func New() *Header {
	return &Header{
		Difficulty: difficulty.New(),
	}
}

// turn a byte slice into a record
func (record PackedHeader) Unpack() (*Header, error) {
	if len(record) != TotalBlockSize {
		return nil, fault.ErrInvalidBlockHeader
	}

	header := New()

	header.Version = binary.LittleEndian.Uint16(record[versionOffset:])
	header.TransactionCount = binary.LittleEndian.Uint16(record[transactionCountOffset:])
	header.Number = binary.LittleEndian.Uint64(record[numberOffset:])

	err := blockdigest.DigestFromBytes(&header.PreviousBlock, record[previousBlockOffset:merkleRootOffset])
	if nil != err {
		return nil, err
	}

	err = merkle.DigestFromBytes(&header.MerkleRoot, record[merkleRootOffset:timestampOffset])
	if nil != err {
		return nil, err
	}

	header.Timestamp = binary.LittleEndian.Uint64(record[timestampOffset:difficultyOffset])
	header.Difficulty.SetBytes(record[difficultyOffset:nonceOffset])
	header.Nonce = NonceType(binary.LittleEndian.Uint64(record[nonceOffset:]))

	return header, nil
}

// digest for a packed header
// make sure to truncate bytes to correct length
func (record PackedHeader) Digest() blockdigest.Digest {
	return blockdigest.NewDigest(record[:TotalBlockSize])
}

// turn a record into an array of bytes
func (header *Header) Pack() PackedHeader {
	buffer := make([]byte, TotalBlockSize)

	binary.LittleEndian.PutUint16(buffer[versionOffset:], header.Version)
	binary.LittleEndian.PutUint16(buffer[transactionCountOffset:], header.TransactionCount)
	binary.LittleEndian.PutUint64(buffer[numberOffset:], header.Number)

	// these are in little endian order so can just copy them
	copy(buffer[previousBlockOffset:], header.PreviousBlock[:])
	copy(buffer[merkleRootOffset:], header.MerkleRoot[:])

	binary.LittleEndian.PutUint64(buffer[timestampOffset:], header.Timestamp)
	binary.LittleEndian.PutUint64(buffer[difficultyOffset:], header.Difficulty.Bits())
	binary.LittleEndian.PutUint64(buffer[nonceOffset:], uint64(header.Nonce))

	return buffer
}
