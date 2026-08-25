package diffpack

type Options struct {
    NodeID     string
    BlockSize  int
    MaxJobs    int
    AuditPath  string
}

func (o Options) withDefaults() Options {
    if o.NodeID == "" {
        o.NodeID = "local"
    }
    if o.BlockSize <= 0 {
        o.BlockSize = 64
    }
    if o.MaxJobs <= 0 {
        o.MaxJobs = 256
    }
    return o
}
