package uuid

import (
	uuid "github.com/satori/go.uuid"
	mh "github.com/tawakhal/util/math"
)

var (
	Empty = UUID{}
)

type UUID struct {
	MSB uint64
	LSB uint64
}

func New() (UUID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return Empty, err
	}

	msb := mh.BytesToUint64(id[:8])
	lsb := mh.BytesToUint64(id[8:])

	return UUID{MSB: msb, LSB: lsb}, nil
}

func FromInt(msb, lsb uint64) UUID {
	return UUID{MSB: msb, LSB: lsb}
}

func FromString(hex string) (UUID, error) {
	id, err := uuid.FromString(hex)
	if err != nil {
		return Empty, err
	}

	msb := mh.BytesToUint64(id[:8])
	lsb := mh.BytesToUint64(id[8:])

	return UUID{MSB: msb, LSB: lsb}, nil
}

func (id UUID) String() string {
	msb := mh.Uint64ToBytes(id.MSB)
	lsb := mh.Uint64ToBytes(id.LSB)
	uid, err := uuid.FromBytes(append(msb, lsb...))
	if err != nil {
		return ""
	}

	return uid.String()
}
