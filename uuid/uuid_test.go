package uuid

import "testing"

func TestUUIDNew(t *testing.T) {
	uid, _ := New()
	t.Logf("UUID[%d:%d]", uid.MSB, uid.LSB)
	t.Log(uid)
}

func TestFromUInt(t *testing.T) {
	var msb, lsb uint64
	msb = 9775335782889704802
	lsb = 10642320787124414530
	uid := FromUint(msb, lsb)
	t.Logf("UUID[%d:%d]", uid.MSB, uid.LSB)
	t.Log(uid)
}

func TestFromString(t *testing.T) {
	var uidstr string
	uidstr = "87a8f81a-434e-4d62-93b1-1e5ed64e6842"
	uid, _ := FromString(uidstr)
	t.Logf("UUID[%d:%d]", uid.MSB, uid.LSB)
	t.Log(uid)
}

func TestIsEmpty(t *testing.T) {
	var uid UUID
	if uid.IsEmpty() {
		t.Log("Success")
		return
	}
	t.Log("failed")
}
