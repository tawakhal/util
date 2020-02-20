package uuid

import (
	uuid "github.com/satori/go.uuid"
	mh "github.com/tawakhal/util/math"
)

// UUID is struct of uuid
type UUID struct {
	MSB uint64
	LSB uint64
}

// IsEmpty is checking uuid is empty or not
func (et UUID) IsEmpty() bool {
	var uid UUID
	if et == uid {
		return true
	}
	return false
}

// String is generate uuid to string
func (et UUID) String() string {
	msb := mh.Uint64ToBytes(et.MSB)
	lsb := mh.Uint64ToBytes(et.LSB)
	uid, err := uuid.FromBytes(append(msb, lsb...))
	if err != nil {
		return ""
	}
	return uid.String()
}

// New is generate uuid with version 4
func New() (uid UUID, err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return uid, err
	}
	uid = UUID{MSB: mh.BytesToUint64(id[:8]), LSB: mh.BytesToUint64(id[8:])}
	return uid, nil
}

// FromUint is generate uuid from uint64
func FromUint(msb, lsb uint64) UUID {
	return UUID{MSB: msb, LSB: lsb}
}

// FromString is generate uuid from string
func FromString(hex string) (uid UUID, err error) {
	id, err := uuid.FromString(hex)
	if err != nil {
		return uid, err
	}
	uid = UUID{MSB: mh.BytesToUint64(id[:8]), LSB: mh.BytesToUint64(id[8:])}
	return uid, nil
}
