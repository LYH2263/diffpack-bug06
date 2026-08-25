package diffpack

import (
	"context"

	"github.com/LYH2263/go-diffpack/internal/validate"
)

func (p *Packer) AcceptJob(ctx context.Context, jobID string, base, target []byte) (*Bundle, error) {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		return nil, ErrClosed
	}
	p.mu.RUnlock()
	if err := validate.JobID(jobID); err != nil {
		return nil, ErrBadInput
	}
	return p.BuildDelta(ctx, jobID, base, target)
}
