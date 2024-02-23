package hexutil

import "encoding/hex"

const ZeroAddress = "0x0000000000000000000000000000000000000000"

func Encode(b []byte) string {
	hx := hex.EncodeToString(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hx) == 0 {
		hx = "0"
	}

	return "0x" + hx
}

func EncodeWith0x(s string) string {
	if s[0:2] == "0x" {
		return s
	}

	return "0x" + s
}

func CopyBytes(data []byte) []byte {
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret
}

func Decode(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		if len(s) >= 2 && s[0:2] == "0x" {
			s = s[2:]
		}
		h, _ := hex.DecodeString(s)

		return h
	}
	return nil
}
