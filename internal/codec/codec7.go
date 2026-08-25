package codec

import "encoding/hex"

// Codec7 hex helpers for job payloads.
func EncodeHex7(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex7(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad7(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
