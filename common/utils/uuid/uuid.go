package uuid

import (
	"encoding/hex"
	"math/big"

	"github.com/google/uuid"
)

type UUID struct {
	id uuid.UUID
}

func GenUUID() *UUID {
	id, _ := uuid.NewUUID()
	return &UUID{
		id: id,
	}
}

func (u *UUID) ConvertToBigInt() *big.Int {
	return new(big.Int).SetBytes(u.id[:])
}

func (u *UUID) String() string {
	return u.ConvertToBigInt().String()
}

func (u *UUID) Encode() string {
	return hex.EncodeToString(u.id[:])
}
