package parser

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Errors
var (
	ErrInvalidDurationFormat = errors.New("invalid duration format")
)

// Parse parses the duration string
func Parse(d string) (time.Duration, error) {
	// Look for ":"
	// Should atleast contain one ":"
	t := strings.Split(d, ":")

	if len(t) <= 1 {
		return 0, ErrInvalidDurationFormat
	}

	if len(t) > 4 {
		return 0, ErrInvalidDurationFormat
	}

	i := len(t) - 1
	j := 0
	dur := 0
	for ; i >= 0; i -= 1 {
		switch j {
		case 0:
			p, err := strconv.ParseInt(t[i], 10, 64)
			if err != nil {
				return 0, errors.Join(ErrInvalidDurationFormat, err)
			}

			if p < 0 {
				return 0, errors.New("seconds cannot be negative")
			}

			if p > 59 {
				return 0, errors.New("seconds cannot be greater than 59")
			}

			// add seconds
			dur += int(p)

		case 1:
			p, err := strconv.ParseInt(t[i], 10, 64)
			if err != nil {
				return 0, errors.Join(ErrInvalidDurationFormat, err)
			}

			if p < 0 {
				return 0, errors.New("minutes cannot be negative")
			}

			if p > 59 {
				return 0, errors.New("minutes cannot be greater than 59")
			}

			// add minutes
			dur += int(p) * 60

		case 2:
			p, err := strconv.ParseInt(t[i], 10, 64)
			if err != nil {
				return 0, errors.Join(ErrInvalidDurationFormat, err)
			}

			if p < 0 {
				return 0, errors.New("hours cannot be negative")
			}

			if p > 23 {
				return 0, errors.New("hours cannot be greater than 23")
			}

			// add hours
			dur += int(p) * 60 * 60

		case 3:
			p, err := strconv.ParseInt(t[i], 10, 64)
			if err != nil {
				return 0, errors.Join(ErrInvalidDurationFormat, err)
			}

			if p < 0 {
				return 0, errors.New("days cannot be negative")
			}

			// add days
			dur += int(p) * 60 * 60 * 24

		}
		j += 1
	}

	return time.Duration(dur) * time.Second, nil
}
