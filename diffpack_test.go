package diffpack_test

import (
	"context"
	"testing"

	dp "github.com/LYH2263/go-diffpack"
)

func newPacker(t *testing.T) *dp.Packer {
	t.Helper()
	p, err := dp.New(dp.Options{MaxJobs: 32})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = p.Close() })
	return p
}

func mustBuild(t *testing.T, p *dp.Packer, id string, base, target []byte) *dp.Bundle {
	t.Helper()
	b, err := p.BuildDelta(context.Background(), id, base, target)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	return b
}
