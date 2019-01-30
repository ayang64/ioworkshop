package repeater

import "errors"

type Repeater struct {
	bytes []byte
	pos   int
}

func (r *Repeater) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i, r.pos = i+1, (r.pos+1)%len(r.bytes) {
		b[i] = r.bytes[r.pos]
	}
	return len(b), nil
}

func New(b []byte) (*Repeater, error) {
	if len(b) == 0 {
		return nil, errors.New("lengh of repeated data must be greater than zero")
	}
	return &Repeater{bytes: b, pos: 0}, nil
}
