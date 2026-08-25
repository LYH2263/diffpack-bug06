package blocksplit

import (
    "github.com/LYH2263/go-diffpack/internal/fingerprint"
)

type Block struct {
    Offset int
    Length int
    Adler  uint32
    SHA    []byte
}

type Index struct {
    Blocks []Block
    byAdler map[uint32][]int
}

func BuildIndex(base []byte, blockSize int) *Index {
    if blockSize <= 0 {
        blockSize = 64
    }
    idx := &Index{byAdler: make(map[uint32][]int)}
    if len(base) == 0 {
        return idx
    }
    for off := 0; off < len(base); off += blockSize {
        end := off + blockSize
        if end > len(base) {
            end = len(base)
        }
        chunk := base[off:end]
        ad := fingerprint.Adler32(chunk)
        sha := fingerprint.SHA256(chunk)
        b := Block{Offset: off, Length: end - off, Adler: ad, SHA: sha}
        idx.Blocks = append(idx.Blocks, b)
        idx.byAdler[ad] = append(idx.byAdler[ad], len(idx.Blocks)-1)
    }
    return idx
}

func (idx *Index) Match(base, target []byte, at int, blockSize int) (copyOff int, copyLen int, ok bool) {
    if at >= len(target) {
        return 0, 0, false
    }
    end := at + blockSize
    if end > len(target) {
        end = len(target)
    }
    chunk := target[at:end]
    ad := fingerprint.Adler32(chunk)
    candidates := idx.byAdler[ad]
    want := fingerprint.SHA256(chunk)
    for _, bi := range candidates {
        b := idx.Blocks[bi]
        if b.Length != len(chunk) {
            continue
        }
        if string(b.SHA) != string(want) {
            continue
        }
        if b.Offset+b.Length > len(base) {
            continue
        }
        if string(base[b.Offset:b.Offset+b.Length]) != string(chunk) {
            continue
        }
        return b.Offset, b.Length, true
    }
    return 0, 0, false
}
