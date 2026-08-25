package codec

import "encoding/hex"

// Codec10 hex helpers for job payloads.
func EncodeHex10(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex10(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad10(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
