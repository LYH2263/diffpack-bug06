package codec

import "encoding/hex"

// Codec3 hex helpers for job payloads.
func EncodeHex3(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex3(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad3(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
