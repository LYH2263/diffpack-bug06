package codec

import "encoding/hex"

// Codec13 hex helpers for job payloads.
func EncodeHex13(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex13(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad13(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
