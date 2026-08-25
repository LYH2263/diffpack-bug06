package metrics

type Counter struct {
    name string
    n    int64
}

func NewCounter(name string) *Counter {
    return &Counter{name: name}
}

func (c *Counter) Add(v int64) {
    c.n += v
}

func (c *Counter) Value() int64 {
    return c.n
}
