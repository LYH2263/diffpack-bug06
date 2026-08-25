package diffpack

func (p *Packer) ListHunks(jobID string) ([]Hunk, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.closed {
		return nil, ErrClosed
	}
	rec, ok := p.store.Get(jobID)
	if !ok {
		return nil, ErrNotFound
	}
	job := jobFromStore(rec.Job)
	out := make([]Hunk, len(job.Hunks))
	copy(out, job.Hunks)
	return out, nil
}

func (p *Packer) GetJob(jobID string) (Job, error) {
	rec, ok := p.store.Get(jobID)
	if !ok {
		return Job{}, ErrNotFound
	}
	return jobFromStore(rec.Job), nil
}
