package codec

import "encoding/hex"

// Codec6 hex helpers for job payloads.
func EncodeHex6(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex6(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad6(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
