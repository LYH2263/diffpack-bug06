package diffpack

import "github.com/LYH2263/go-diffpack/internal/jobstore"

func (p *Packer) DeleteJob(jobID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return ErrClosed
	}
	if !p.store.Delete(jobID) {
		return ErrNotFound
	}
	return nil
}

func (p *Packer) MarkVerified(jobID string) error {
	rec, ok := p.store.Get(jobID)
	if !ok {
		return ErrNotFound
	}
	if rec.Job.Verified {
		return ErrJobVerified
	}
	rec.Job.Verified = true
	rec.Job.Status = jobstore.JobStatus(JobVerified)
	return p.store.Put(jobID, rec.Job, rec.Bundle)
}

func (p *Packer) ListJobs() []Job {
	stored := p.store.List()
	out := make([]Job, len(stored))
	for i, j := range stored {
		out[i] = jobFromStore(j)
	}
	return out
}
