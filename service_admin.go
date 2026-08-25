package diffpack

func (p *Packer) Health() Health {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return Health{OK: !p.closed, Closed: p.closed, NodeID: p.opts.NodeID}
}

func (p *Packer) Stats() Stats {
	total, verified, bytes := p.store.Stats()
	return Stats{
		JobsTotal:    total,
		JobsVerified: verified,
		BytesDelta:   bytes,
	}
}
