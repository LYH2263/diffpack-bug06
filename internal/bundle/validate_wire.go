package bundle

import "fmt"

func ValidateWire(w *Wire) error {
	if w == nil {
		return fmt.Errorf("bundle: nil wire")
	}
	if len(w.BaseFingerprint) == 0 {
		return fmt.Errorf("bundle: empty fingerprint")
	}
	var size int64
	for _, op := range w.Ops {
		switch op.Kind {
		case OpCopy:
			if op.Length <= 0 {
				return fmt.Errorf("bundle: bad copy len")
			}
			size += int64(op.Length)
		case OpInsert:
			if len(op.Data) == 0 {
				return fmt.Errorf("bundle: empty insert")
			}
			size += int64(len(op.Data))
		default:
			return fmt.Errorf("bundle: unknown op")
		}
	}
	if size != w.TargetSize {
		return fmt.Errorf("bundle: size mismatch")
	}
	return nil
}
