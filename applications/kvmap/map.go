package kvmap

import (
	"fmt"
	"io"
	"unicode"
)

type Decoder struct {
	tokenState int
	io.RuneScanner
}

func (d *Decoder) peek() (rune, error) {
	r, _, err := d.RuneScanner.ReadRune()
	if err != nil {
		return ' ', err
	}
	d.RuneScanner.UnreadRune()
	return r, nil
}

func (d *Decoder) eatWhiteSpace() error {
	for {
		// peek at next rune
		r, err := d.peek()

		if err != nil {
			return err
		}

		// if our rune is not white space, immediately return.
		if unicode.IsSpace(r) == false {
			return nil
		}

		// if our rune *is* white space, consume it.
		if _, _, err := d.RuneScanner.ReadRune(); err != nil {
			return err
		}
	}
}

type scannerState int

const (
	expectingKey = scannerState(iota)
	expectingValue
)

func (s scannerState) String() string {
	switch s {
	case 0:
		return "expecting-key"
	case 1:
		return "expecting-value"
	}
	return "*error-unknown-state*"
}

func (s scannerState) Next() scannerState {
	return (s + 1) % 2
}

func (d *Decoder) readAtom() (string, error) {
	rc := ""
	for {
		r, err := d.peek()

		if err != nil {
			return "", err
		}

		if unicode.IsSpace(r) {
			break
		}

		rc += string(r)

		if _, _, err := d.RuneScanner.ReadRune(); err != nil {
			return "", err
		}
	}
	return rc, nil
}

func (d *Decoder) Unmarshal(m map[string]string) error {
	for state, key := scannerState(0), ""; ; state = state.Next() {
		// ignore any errors here.  if there are any, they'll crop up in our
		// next attempt to read and we'll handle it there.
		d.eatWhiteSpace()

		atom, err := d.readAtom()

		if err == io.EOF {
			// we've received EOF before we have a complete k/v pair so return
			// with an Unexpected EOF error.
			if state == expectingValue {
				return io.ErrUnexpectedEOF
			}

			// we're in a coherent state.  if we've received an io.EOF then
			// we've completed the operation.
			break
		}

		if err != nil {
			return err
		}

		switch state {
		case expectingKey:
			key = atom
		case expectingValue:
			m[key] = atom
		default:
			return fmt.Errorf("scanner entered an unknown state %q (%d)", state, state)
		}
	}
	return nil
}

func NewDecoder(rs io.RuneScanner) *Decoder {
	return &Decoder{
		RuneScanner: rs,
	}
}
