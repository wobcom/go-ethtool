package util

import (
	"encoding/hex"
	"unicode/utf8"
)

// GetValidUtf8String returns in any case valid utf8 string for any input
func GetValidUtf8String(raw []byte) string {
	if utf8.Valid(raw) {
		return string(raw)
	}
	return hex.EncodeToString(raw)
}
