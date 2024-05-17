package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	t.Run("valid format", func(t *testing.T) {
		var formats = []string{"00:22", "01:01:02:30"}
		for _, format := range formats {
			_, err := Parse(format)
			assert.NoError(t, err)
		}
	})

	t.Run("invalid format", func(t *testing.T) {
		var formats = []string{"22", "01:01:91:92:01", "01:01:91 00:92:01", ":22"}
		for _, format := range formats {
			dur, err := Parse(format)
			assert.ErrorIs(t, err, ErrInvalidDurationFormat)
			assert.Equal(t, time.Duration(0), dur)
		}
	})

	t.Run("verify duration", func(t *testing.T) {
		type Test struct {
			f  string
			ex float64
		}

		var tests = []Test{
			{"00:22", 22},
			{"03:12", 192},
			{"01:01:02:30", 90150},
		}

		for _, test := range tests {
			dur, err := Parse(test.f)
			assert.NoError(t, err)
			assert.Equal(t, test.ex, dur.Seconds())
		}
	})
}
