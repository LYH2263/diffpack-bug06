package codec

import "encoding/hex"

// Codec8 hex helpers for job payloads.
func EncodeHex8(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex8(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad8(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
