package diffpack

import "errors"

var (
    ErrClosed       = errors.New("diffpack: packer closed")
    ErrNotFound     = errors.New("diffpack: job not found")
    ErrInvalidBundle = errors.New("diffpack: invalid bundle")
    ErrJobVerified  = errors.New("diffpack: job already verified")
    ErrBadInput     = errors.New("diffpack: bad input")
)
