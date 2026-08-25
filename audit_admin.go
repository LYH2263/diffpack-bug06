package diffpack

func (p *Packer) RotateAudit() error {
	if p.audit == nil {
		return nil
	}
	return p.audit.Rotate()
}
