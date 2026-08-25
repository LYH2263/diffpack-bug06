package codec

import "encoding/hex"

// Codec1 hex helpers for job payloads.
func EncodeHex1(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex1(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad1(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
