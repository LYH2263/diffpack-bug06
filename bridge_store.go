package diffpack

import "github.com/LYH2263/go-diffpack/internal/jobstore"

func jobToStore(j Job) jobstore.Job {
	hunks := make([]jobstore.Hunk, len(j.Hunks))
	for i, h := range j.Hunks {
		hunks[i] = jobstore.Hunk{
			Index:     h.Index,
			Kind:      h.Kind,
			BaseStart: h.BaseStart,
			Length:    h.Length,
			InsertLen: h.InsertLen,
		}
	}
	return jobstore.Job{
		ID:         j.ID,
		Status:     jobstore.JobStatus(j.Status),
		BaseSize:   j.BaseSize,
		TargetSize: j.TargetSize,
		Hunks:      hunks,
		Created:    j.Created,
		Verified:   j.Verified,
	}
}

func jobFromStore(j jobstore.Job) Job {
	hunks := make([]Hunk, len(j.Hunks))
	for i, h := range j.Hunks {
		hunks[i] = Hunk{
			Index:     h.Index,
			Kind:      h.Kind,
			BaseStart: h.BaseStart,
			Length:    h.Length,
			InsertLen: h.InsertLen,
		}
	}
	return Job{
		ID:         j.ID,
		Status:     JobStatus(j.Status),
		BaseSize:   j.BaseSize,
		TargetSize: j.TargetSize,
		Hunks:      hunks,
		Created:    j.Created,
		Verified:   j.Verified,
	}
}

func bundleToStore(b *Bundle) *jobstore.Bundle {
	if b == nil {
		return nil
	}
	ops := make([]jobstore.Op, len(b.Ops))
	for i, op := range b.Ops {
		ops[i] = jobstore.Op{
			Kind:   jobstore.OpKind(op.Kind),
			Offset: op.Offset,
			Length: op.Length,
			Data:   op.Data,
		}
	}
	return &jobstore.Bundle{
		BaseFingerprint: append([]byte(nil), b.BaseFingerprint...),
		TargetSize:      b.TargetSize,
		Ops:             ops,
	}
}

func bundleFromStore(b *jobstore.Bundle) *Bundle {
	if b == nil {
		return nil
	}
	ops := make([]Op, len(b.Ops))
	for i, op := range b.Ops {
		ops[i] = Op{
			Kind:   OpKind(op.Kind),
			Offset: op.Offset,
			Length: op.Length,
			Data:   op.Data,
		}
	}
	return &Bundle{
		BaseFingerprint: append([]byte(nil), b.BaseFingerprint...),
		TargetSize:      b.TargetSize,
		Ops:             ops,
	}
}
