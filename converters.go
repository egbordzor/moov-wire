// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"regexp"
	"strconv"
	"strings"
)

// converters handles golang to WIRE type Converters
type converters struct{}

func (c *converters) parseNumField(r string) (s int) {
	s, _ = strconv.Atoi(strings.TrimSpace(r))
	return s
}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}

func (c *converters) truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// numericStringField right-justified zero filled
func (c *converters) numericStringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

// cleanupDelimiters removes non-necessary extra "*" from the end of a string, and keeps only one
func (c *converters) cleanupDelimiters(line string) string {
	re := regexp.MustCompile(`\*{2,}$`)
	return re.ReplaceAllString(line, "*")
}

func (c *converters) prettyMessage(lines []*string, sep string) string {
	var rv []string
	for _, l := range lines {
		if l == nil || strings.TrimSpace(*l) == "" {
			continue
		}
		rv = append(rv, *l)
	}
	return strings.Join(rv, sep)
}
