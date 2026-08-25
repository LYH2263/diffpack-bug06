package diffpack

import (
	"context"
	"sync"

	"github.com/LYH2263/go-diffpack/internal/audit"
	"github.com/LYH2263/go-diffpack/internal/delta"
	"github.com/LYH2263/go-diffpack/internal/jobstore"
)

type Packer struct {
	opts   Options
	store  *jobstore.Store
	audit  *audit.Logger
	mu     sync.RWMutex
	closed bool
}

func New(opts Options) (*Packer, error) {
	o := opts.withDefaults()
	p := &Packer{
		opts:  o,
		store: jobstore.New(o.MaxJobs),
	}
	if o.AuditPath != "" {
		lg, err := audit.Open(o.AuditPath)
		if err != nil {
			return nil, err
		}
		p.audit = lg
	}
	return p, nil
}

func bundleFromBuilt(b *delta.Built) *Bundle {
	if b == nil {
		return nil
	}
	ops := make([]Op, len(b.Ops))
	for i, op := range b.Ops {
		ops[i] = Op{
			Kind:   OpKind(op.Kind),
			Offset: op.Offset,
			Length: op.Length,
			Data:   append([]byte(nil), op.Data...),
		}
	}
	return &Bundle{
		BaseFingerprint: append([]byte(nil), b.Fingerprint...),
		TargetSize:      b.TargetSize,
		Ops:             ops,
	}
}

func hunksFromBuilt(b *delta.Built) []Hunk {
	out := make([]Hunk, len(b.Hunks))
	for i, h := range b.Hunks {
		out[i] = Hunk{
			Index:     h.Index,
			Kind:      h.Kind,
			BaseStart: h.BaseStart,
			Length:    h.Length,
			InsertLen: h.InsertLen,
		}
	}
	return out
}

func (p *Packer) BuildDelta(ctx context.Context, jobID string, base, target []byte) (*Bundle, error) {
	if jobID == "" || len(base) == 0 && len(target) == 0 {
		return nil, ErrBadInput
	}
	built, err := delta.Build(ctx, base, target, p.opts.BlockSize)
	if err != nil {
		return nil, err
	}
	bundle := bundleFromBuilt(built)
	job := Job{
		ID:         jobID,
		Status:     JobBuilt,
		BaseSize:   len(base),
		TargetSize: len(target),
		Hunks:      hunksFromBuilt(built),
	}
	if err := p.store.Put(jobID, jobToStore(job), bundleToStore(bundle)); err != nil {
		if err == jobstore.ErrStoreFull {
			return nil, ErrBadInput
		}
		return nil, err
	}
	if p.audit != nil {
		p.audit.Printf("build %s base=%d target=%d ops=%d", jobID, len(base), len(target), len(bundle.Ops))
	}
	return bundle, nil
}

func (p *Packer) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.closed = true
	if p.audit != nil {
		return p.audit.Close()
	}
	return nil
}
