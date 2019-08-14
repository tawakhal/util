package math

func BytesToUint64(b []byte) uint64 {
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

func BytesToUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func BytesToUint16(b []byte) uint16 {
	return uint16(b[1]) | uint16(b[0])<<8
}

func Bytes16ToUint32(b []byte) uint32 {
	return uint32(b[1]) | uint32(b[0])<<8
}

func Uint64ToBytes(n uint64) []byte {
	bytes := make([]byte, 8)
	bytes[0] = byte((n >> 56) & 0xFF)
	bytes[1] = byte((n >> 48) & 0xFF)
	bytes[2] = byte((n >> 40) & 0xFF)
	bytes[3] = byte((n >> 32) & 0xFF)
	bytes[4] = byte((n >> 24) & 0xFF)
	bytes[5] = byte((n >> 16) & 0xFF)
	bytes[6] = byte((n >> 8) & 0xFF)
	bytes[7] = byte(n & 0xFF)

	return bytes
}

func Uint32ToBytes(n uint32) []byte {
	bytes := make([]byte, 4)
	bytes[0] = byte((n >> 24) & 0xFF)
	bytes[1] = byte((n >> 16) & 0xFF)
	bytes[2] = byte((n >> 8) & 0xFF)
	bytes[3] = byte(n & 0xFF)

	return bytes
}

func Uint16ToBytes(n uint16) []byte {
	bytes := make([]byte, 2)
	bytes[0] = byte((n >> 8) & 0xFF)
	bytes[1] = byte(n & 0xFF)

	return bytes
}

func UintToBytes(n uint64, length int) []byte {
	bytes := make([]byte, length)
	for index := 0; index < length; index++ {
		shift := uint64((length - index - 1) * 8)
		if shift > 0 {
			bytes[index] = byte((n >> shift) & 0xFF)
		} else {
			bytes[index] = byte(n & 0xFF)
		}
	}
	return bytes
}

func BytesToUint(bytes []byte, length int) uint64 {
	val := uint64(bytes[0])
	for index := 1; index < length; index++ {
		val = (val << 8) | uint64(bytes[index])
	}
	return val
}

func BytesToInt64(b []byte) int64 {
	return int64(b[7]) | int64(b[6])<<8 | int64(b[5])<<16 | int64(b[4])<<24 |
		int64(b[3])<<32 | int64(b[2])<<40 | int64(b[1])<<48 | int64(b[0])<<56
}

func Int64ToBytes(n int64) []byte {
	bytes := make([]byte, 8)
	bytes[0] = byte((n >> 56) & 0xFF)
	bytes[1] = byte((n >> 48) & 0xFF)
	bytes[2] = byte((n >> 40) & 0xFF)
	bytes[3] = byte((n >> 32) & 0xFF)
	bytes[4] = byte((n >> 24) & 0xFF)
	bytes[5] = byte((n >> 16) & 0xFF)
	bytes[6] = byte((n >> 8) & 0xFF)
	bytes[7] = byte(n & 0xFF)

	return bytes
}
