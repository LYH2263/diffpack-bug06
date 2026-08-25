package bundle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

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

type Wire struct {
	BaseFingerprint []byte
	TargetSize      int64
	Ops             []Op
}

var magic = []byte("DPK1")

func Encode(w io.Writer, b *Wire) error {
	if b == nil {
		return fmt.Errorf("bundle: nil")
	}
	if _, err := w.Write(magic); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint16(1)); err != nil {
		return err
	}
	fp := b.BaseFingerprint
	if err := binary.Write(w, binary.LittleEndian, uint16(len(fp))); err != nil {
		return err
	}
	if _, err := w.Write(fp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, b.TargetSize); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint32(len(b.Ops))); err != nil {
		return err
	}
	for _, op := range b.Ops {
		if err := binary.Write(w, binary.LittleEndian, uint8(op.Kind)); err != nil {
			return err
		}
		switch op.Kind {
		case OpCopy:
			if err := binary.Write(w, binary.LittleEndian, op.Offset); err != nil {
				return err
			}
			if err := binary.Write(w, binary.LittleEndian, op.Length); err != nil {
				return err
			}
		case OpInsert:
			if err := binary.Write(w, binary.LittleEndian, uint32(len(op.Data))); err != nil {
				return err
			}
			if _, err := w.Write(op.Data); err != nil {
				return err
			}
		default:
			return fmt.Errorf("bundle: bad op %d", op.Kind)
		}
	}
	return nil
}

func Decode(r io.Reader) (*Wire, error) {
	hdr := make([]byte, 4)
	if _, err := io.ReadFull(r, hdr); err != nil {
		return nil, err
	}
	if !bytes.Equal(hdr, magic) {
		return nil, fmt.Errorf("bundle: bad magic")
	}
	var ver uint16
	if err := binary.Read(r, binary.LittleEndian, &ver); err != nil {
		return nil, err
	}
	var fpLen uint16
	if err := binary.Read(r, binary.LittleEndian, &fpLen); err != nil {
		return nil, err
	}
	fp := make([]byte, fpLen)
	if _, err := io.ReadFull(r, fp); err != nil {
		return nil, err
	}
	var targetSize int64
	if err := binary.Read(r, binary.LittleEndian, &targetSize); err != nil {
		return nil, err
	}
	var opCount uint32
	if err := binary.Read(r, binary.LittleEndian, &opCount); err != nil {
		return nil, err
	}
	ops := make([]Op, opCount)
	for i := 0; i < int(opCount); i++ {
		var kind uint8
		if err := binary.Read(r, binary.LittleEndian, &kind); err != nil {
			return nil, err
		}
		switch OpKind(kind) {
		case OpCopy:
			var off int64
			var ln int32
			if err := binary.Read(r, binary.LittleEndian, &off); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &ln); err != nil {
				return nil, err
			}
			ops[i] = Op{Kind: OpCopy, Offset: off, Length: ln}
		case OpInsert:
			var dlen uint32
			if err := binary.Read(r, binary.LittleEndian, &dlen); err != nil {
				return nil, err
			}
			data := make([]byte, dlen)
			if _, err := io.ReadFull(r, data); err != nil {
				return nil, err
			}
			ops[i] = Op{Kind: OpInsert, Data: data}
		default:
			return nil, fmt.Errorf("bundle: bad op kind")
		}
	}
	return &Wire{
		BaseFingerprint: fp,
		TargetSize:      targetSize,
		Ops:             ops,
	}, nil
}

func Marshal(b *Wire) ([]byte, error) {
	var buf bytes.Buffer
	if err := Encode(&buf, b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(p []byte) (*Wire, error) {
	return Decode(bytes.NewReader(p))
}
