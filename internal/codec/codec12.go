package codec

import "encoding/hex"

// Codec12 hex helpers for job payloads.
func EncodeHex12(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex12(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad12(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
