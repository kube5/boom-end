package invite_code

import (
	"strings"
)

const (
	BASE    = "1E8S2DZX9WYLTN6BQF7CP5IK3MJUAR4HV"
	DECIMAL = 33
	PAD     = "A"
	LEN     = 4
)

func Encode(uid uint64) string {
	id := uid
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func Decode(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)
	baseArr := []byte(BASE)
	baseRev := make(map[byte]int)
	for k, v := range baseArr {
		baseRev[v] = k
	}
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += uint64(index) * b
		r++
	}
	return res
}
