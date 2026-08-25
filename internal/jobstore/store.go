package jobstore

import (
	"errors"
	"sync"
	"time"
)

var ErrStoreFull = errors.New("jobstore: full")

type JobStatus string

const (
	StatusPending  JobStatus = "pending"
	StatusBuilt    JobStatus = "built"
	StatusVerified JobStatus = "verified"
	StatusFailed   JobStatus = "failed"
)

type Hunk struct {
	Index     int
	Kind      string
	BaseStart int64
	Length    int32
	InsertLen int
}

type Job struct {
	ID         string
	Status     JobStatus
	BaseSize   int
	TargetSize int
	Hunks      []Hunk
	Created    time.Time
	Verified   bool
}

type OpKind uint8

const (
	OpCopy   OpKind = 1
	OpInsert OpKind = 2
)

type Op struct {
	Kind   OpKind
	Offset int64
	Length int32
	Data   []byte
}

type Bundle struct {
	BaseFingerprint []byte
	TargetSize      int64
	Ops             []Op
}

type Record struct {
	Job    Job
	Bundle *Bundle
}

type Store struct {
	mu      sync.RWMutex
	jobs    map[string]*Record
	maxJobs int
	bytes   int64
}

func New(max int) *Store {
	return &Store{jobs: make(map[string]*Record), maxJobs: max}
}

func (s *Store) Put(id string, job Job, bundle *Bundle) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.jobs) >= s.maxJobs && s.jobs[id] == nil {
		return ErrStoreFull
	}
	if job.Created.IsZero() {
		job.Created = time.Now()
	}
	s.jobs[id] = &Record{Job: job, Bundle: bundle}
	if bundle != nil {
		s.bytes += bundle.TargetSize
	}
	return nil
}

func (s *Store) Get(id string) (*Record, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.jobs[id]
	return r, ok
}

func (s *Store) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.jobs[id]
	if !ok {
		return false
	}
	if r.Bundle != nil {
		s.bytes -= r.Bundle.TargetSize
	}
	delete(s.jobs, id)
	return true
}

func (s *Store) List() []Job {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Job, 0, len(s.jobs))
	for _, r := range s.jobs {
		out = append(out, r.Job)
	}
	return out
}

func (s *Store) Stats() (total, verified int, bytes int64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, r := range s.jobs {
		if r.Job.Verified {
			verified++
		}
	}
	return len(s.jobs), verified, s.bytes
}
