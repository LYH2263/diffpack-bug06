package diffpack

import "time"

type OpKind uint8

const (
    OpCopy OpKind = 1
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

type Hunk struct {
    Index     int
    Kind      string
    BaseStart int64
    Length    int32
    InsertLen int
}

type JobStatus string

const (
    JobPending   JobStatus = "pending"
    JobBuilt     JobStatus = "built"
    JobVerified  JobStatus = "verified"
    JobFailed    JobStatus = "failed"
)

type Job struct {
    ID        string
    Status    JobStatus
    BaseSize  int
    TargetSize int
    Hunks     []Hunk
    Created   time.Time
    Verified  bool
}

type Stats struct {
    JobsTotal    int
    JobsVerified int
    BytesDelta   int64
}

type Health struct {
    OK      bool
    Closed  bool
    NodeID  string
}
