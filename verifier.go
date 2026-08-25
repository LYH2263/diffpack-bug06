package diffpack

import (
	wirepkg "github.com/LYH2263/go-diffpack/internal/bundle"
)

func (p *Packer) VerifyBundle(b *Bundle) error {
	if b == nil {
		return ErrInvalidBundle
	}
	var size int64
	for _, op := range b.Ops {
		switch op.Kind {
		case OpCopy:
			size += int64(op.Length)
		case OpInsert:
			size += int64(len(op.Data))
		default:
			return ErrInvalidBundle
		}
	}
	if size != b.TargetSize {
		return ErrInvalidBundle
	}
	if len(b.BaseFingerprint) == 0 {
		return ErrInvalidBundle
	}
	return nil
}

func (p *Packer) ValidateStoredBundle(jobID string) error {
	rec, ok := p.store.Get(jobID)
	if !ok {
		return ErrNotFound
	}
	if rec.Bundle == nil {
		return ErrInvalidBundle
	}
	w := wireFromBundle(bundleFromStore(rec.Bundle))
	return wirepkg.ValidateWire(w)
}
