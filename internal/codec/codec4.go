package codec

import "encoding/hex"

// Codec4 hex helpers for job payloads.
func EncodeHex4(b []byte) string {
    return hex.EncodeToString(b)
}

func DecodeHex4(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func Pad4(b []byte, n int) []byte {
    if n <= 0 {
        return b
    }
    out := make([]byte, len(b)+n)
    copy(out, b)
    return out
}
