package codec

import "encoding/hex"

// Codec11 hex helpers for job payloads.
func EncodeHex11(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex11(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad11(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
