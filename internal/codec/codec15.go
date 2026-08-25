package codec

import "encoding/hex"

// Codec15 hex helpers for job payloads.
func EncodeHex15(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex15(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad15(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
