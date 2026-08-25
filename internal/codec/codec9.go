package codec

import "encoding/hex"

// Codec9 hex helpers for job payloads.
func EncodeHex9(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex9(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad9(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
