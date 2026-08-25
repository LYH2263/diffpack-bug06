package codec

import "encoding/hex"

// Codec2 hex helpers for job payloads.
func EncodeHex2(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex2(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad2(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
