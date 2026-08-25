package fingerprint

import (
    "crypto/sha256"
    "hash"
    "hash/adler32"
)

func Adler32(b []byte) uint32 {
    return adler32.Checksum(b)
}

func SHA256(b []byte) []byte {
    h := sha256.Sum256(b)
    out := make([]byte, len(h))
    copy(out, h[:])
    return out
}

func RollingAdler(h uint32, remove, add byte, window int) uint32 {
    const mod = 65521
    var a = byte(h & 0xff)
    var b = byte((h >> 8) & 0xff)
    a = a - remove + add
    b = b - byte(window)*remove + a
    return uint32(a) | (uint32(b) << 8)
}

type Hasher struct {
    adler hash.Hash32
    sha   hash.Hash
}

func NewHasher() *Hasher {
    return &Hasher{adler: adler32.New(), sha: sha256.New()}
}

func (h *Hasher) Write(p []byte) (int, error) {
    if _, err := h.adler.Write(p); err != nil {
        return 0, err
    }
    return h.sha.Write(p)
}

func (h *Hasher) Sum() (adler uint32, sha []byte) {
    s := h.sha.Sum(nil)
    out := make([]byte, len(s))
    copy(out, s)
    return h.adler.Sum32(), out
}
