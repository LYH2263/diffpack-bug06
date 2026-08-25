package diffpack

import (
	"context"

	"github.com/LYH2263/go-diffpack/internal/delta"
)

func (p *Packer) ApplyDelta(ctx context.Context, bundle *Bundle, base []byte) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	p.mu.RLock()
	closed := p.closed
	p.mu.RUnlock()
	if closed {
		return nil, ErrClosed
	}
	if bundle == nil {
		return nil, ErrInvalidBundle
	}
	ops := make([]delta.Op, len(bundle.Ops))
	for i, op := range bundle.Ops {
		ops[i] = delta.Op{
			Kind:   delta.OpKind(op.Kind),
			Offset: op.Offset,
			Length: op.Length,
			Data:   op.Data,
		}
	}
	return delta.ApplyBundle(ops, bundle.TargetSize, base)
}
