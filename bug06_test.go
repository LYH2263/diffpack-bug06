package diffpack_test

import (
	"context"
	"testing"
)

func TestBug06_ExportBundlePreservesOps(t *testing.T) {
	p := newPacker(t)
	_, _ = p.BuildDelta(context.Background(), "e1", []byte("hello"), []byte("hallo"))
	a, err := p.ExportBundle("e1")
	if err != nil {
		t.Fatal(err)
	}
	b, err := p.ExportBundle("e1")
	if err != nil {
		t.Fatal(err)
	}
	if len(a) != len(b) {
		t.Fatalf("len %d vs %d", len(a), len(b))
	}
}
