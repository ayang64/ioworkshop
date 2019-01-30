package kvmap_test

import (
	"io"
	"strings"
	"testing"

	"github.com/ayang64/ioworkshop/applications/kvmap"
)

func TestNewDecoder(t *testing.T) {

	tests := []struct {
		Name          string
		NillError     bool
		ExpectedError error
		Text          string
	}{
		{
			Name:          "Properly Formed",
			NillError:     true,
			ExpectedError: nil,
			Text: `
ipaddress		10.0.0.10
netmask			255.255.255.0
router			10.0.0.1
hostname		bogoparse.nyx.net
`,
		}, {
			Name:          "Missing Value",
			NillError:     false,
			ExpectedError: io.ErrUnexpectedEOF,
			Text: `
ipaddress		10.0.0.10
netmask			255.255.255.0
router			10.0.0.1
hostname
`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			d := kvmap.NewDecoder(strings.NewReader(test.Text))
			config := make(map[string]string)

			err := d.Unmarshal(config)

			if err != test.ExpectedError {
				t.Logf("Unmarshal() returned %v; expected %v", err, test.ExpectedError)
				t.Fail()
			}
		})
	}
}
