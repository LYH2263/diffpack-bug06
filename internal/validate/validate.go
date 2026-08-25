package validate

import "errors"

var ErrBadID = errors.New("validate: bad job id")

func JobID(id string) error {
	if id == "" || len(id) > 128 {
		return ErrBadID
	}
	for _, c := range id {
		if c == '/' || c == '\\' {
			return ErrBadID
		}
	}
	return nil
}
