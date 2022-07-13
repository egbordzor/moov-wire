// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// FIToFI is financial institution to financial institution
type FIToFI struct {
	converters
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

func (ff *FIToFI) AllLines() []*string {
	return []*string{
		&ff.LineOne,
		&ff.LineTwo,
		&ff.LineThree,
		&ff.LineFour,
		&ff.LineFive,
		&ff.LineSix,
	}
}

func (ff *FIToFI) FullText(sep string) string {
	return ff.prettyMessage(ff.AllLines(), sep)
}
