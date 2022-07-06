// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// AdditionalFIToFI is additional financial institution to financial institution information
type AdditionalFIToFI struct {
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`
	// LineFive
	LineFive string `json:"lineFive,omitempty"`
	// LineSix
	LineSix string `json:"lineSix,omitempty"`
}

func (ff *AdditionalFIToFI) String() string {
	return strings.TrimSpace(strings.Join([]string{ff.LineOne, ff.LineTwo, ff.LineThree, ff.LineFour, ff.LineFive, ff.LineSix}, ""))
}
