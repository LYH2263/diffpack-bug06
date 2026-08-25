package diffpack

import (
	"github.com/LYH2263/go-diffpack/internal/bundle"
)

func wireFromBundle(b *Bundle) *bundle.Wire {
	if b == nil {
		return nil
	}
	ops := make([]bundle.Op, len(b.Ops))
	for i, op := range b.Ops {
		ops[i] = bundle.Op{
			Kind:   bundle.OpKind(op.Kind),
			Offset: op.Offset,
			Length: op.Length,
			Data:   op.Data,
		}
	}
	return &bundle.Wire{
		BaseFingerprint: append([]byte(nil), b.BaseFingerprint...),
		TargetSize:      b.TargetSize,
		Ops:             ops,
	}
}

func (p *Packer) ExportBundle(jobID string) ([]byte, error) {
	rec, ok := p.store.Get(jobID)
	if !ok {
		return nil, ErrNotFound
	}
	if rec.Bundle == nil {
		return nil, ErrInvalidBundle
	}
	raw, err := bundle.Marshal(wireFromBundle(bundleFromStore(rec.Bundle)))
	if err != nil {
		return nil, err
	}
	// ExportBundle is a read of the stored bundle; it must not mutate the
	// store-internal record. Clearing rec.Bundle.Ops here left subsequent
	// exports (and ValidateStoredBundle) with an empty ops list, making the
	// second export's wire shorter than the first.
	return raw, nil
}
