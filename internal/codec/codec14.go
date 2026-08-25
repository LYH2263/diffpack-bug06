package codec

import "encoding/hex"

// Codec14 hex helpers for job payloads.
func EncodeHex14(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex14(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad14(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
