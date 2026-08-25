package codec

import "encoding/hex"

// Codec5 hex helpers for job payloads.
func EncodeHex5(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex5(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad5(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
