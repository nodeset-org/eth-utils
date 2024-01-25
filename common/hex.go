package common

import (
	"encoding/hex"
	"strings"
)

const (
	hexPrefix string = "0x"
)

// Add a prefix to a hex string if not present
func EncodeHexWithPrefix(value []byte) string {
	hexString := hex.EncodeToString(value)
	return hexPrefix + hexString
}

// Convert a hex-encoded string to a byte array, removing the prefix if present
func DecodeHex(value string) ([]byte, error) {
	value = strings.TrimPrefix(value, hexPrefix)
	return hex.DecodeString(value)
}
