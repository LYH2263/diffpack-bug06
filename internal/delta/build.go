package delta

import (
	"context"
	"fmt"

	"github.com/LYH2263/go-diffpack/internal/blocksplit"
	"github.com/LYH2263/go-diffpack/internal/fingerprint"
)

type OpKind uint8

const (
	OpCopy   OpKind = 1
	OpInsert OpKind = 2
)

type Op struct {
	Kind   OpKind
	Offset int64
	Length int32
	Data   []byte
}

type Hunk struct {
	Index     int
	Kind      string
	BaseStart int64
	Length    int32
	InsertLen int
}

type Built struct {
	Fingerprint []byte
	TargetSize  int64
	Ops         []Op
	Hunks       []Hunk
}

func Build(ctx context.Context, base, target []byte, blockSize int) (*Built, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	fp := fingerprint.SHA256(base)
	idx := blocksplit.BuildIndex(base, blockSize)
	var ops []Op
	var hunks []Hunk
	pos := 0
	hunkIdx := 0
	for pos < len(target) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		off, ln, ok := idx.Match(base, target, pos, blockSize)
		if ok && ln > 0 {
			ops = append(ops, Op{Kind: OpCopy, Offset: int64(off), Length: int32(ln)})
			hunks = append(hunks, Hunk{
				Index:     hunkIdx,
				Kind:      "copy",
				BaseStart: int64(off),
				Length:    int32(ln),
			})
			hunkIdx++
			pos += ln
			continue
		}
		data := target[pos : pos+1]
		ops = append(ops, Op{Kind: OpInsert, Data: data})
		hunks = append(hunks, Hunk{
			Index:     hunkIdx,
			Kind:      "insert",
			InsertLen: 1,
		})
		hunkIdx++
		pos++
	}
	return &Built{
		Fingerprint: fp,
		TargetSize:  int64(len(target)),
		Ops:         ops,
		Hunks:       hunks,
	}, nil
}

func Apply(built *Built, base []byte) ([]byte, error) {
	if built == nil {
		return nil, fmt.Errorf("delta: nil")
	}
	out := make([]byte, 0, int(built.TargetSize))
	for _, op := range built.Ops {
		switch op.Kind {
		case OpCopy:
			start := int(op.Offset)
			end := start + int(op.Length)
			if start < 0 || end > len(base) {
				return nil, fmt.Errorf("delta: copy out of range")
			}
			out = append(out, base[start:end]...)
		case OpInsert:
			out = append(out, op.Data...)
		default:
			return nil, fmt.Errorf("delta: bad op")
		}
	}
	if int64(len(out)) != built.TargetSize {
		return nil, fmt.Errorf("delta: size mismatch")
	}
	return out, nil
}

func ApplyBundle(bundleOps []Op, targetSize int64, base []byte) ([]byte, error) {
	return Apply(&Built{TargetSize: targetSize, Ops: bundleOps}, base)
}
