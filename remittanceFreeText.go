// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// RemittanceFreeText is the remittance free text
type RemittanceFreeText struct {
	// tag
	tag string
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceFreeText returns a new RemittanceFreeText
func NewRemittanceFreeText() *RemittanceFreeText {
	rft := &RemittanceFreeText{
		tag: TagRemittanceFreeText,
	}
	return rft
}

// Parse takes the input string and parses the RemittanceFreeText values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rft *RemittanceFreeText) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 429 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 429] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	rft.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		rft.LineOne = rft.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		rft.LineTwo = rft.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		rft.LineThree = rft.parseStringField(optionalFields[2])
	}
	return nil
}

func (rft *RemittanceFreeText) UnmarshalJSON(data []byte) error {
	type Alias RemittanceFreeText
	aux := struct {
		*Alias
	}{
		(*Alias)(rft),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rft.tag = TagRemittanceFreeText
	return nil
}

// String writes RemittanceFreeText
func (rft *RemittanceFreeText) String() string {
	var buf strings.Builder
	buf.Grow(426)
	buf.WriteString(rft.tag)
	buf.WriteString(strings.TrimSpace(rft.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(rft.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(rft.LineThreeField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceFreeText and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rft *RemittanceFreeText) Validate() error {
	if rft.tag != TagRemittanceFreeText {
		return fieldError("tag", ErrValidTagForType, rft.tag)
	}
	if err := rft.isAlphanumeric(rft.LineOne); err != nil {
		return fieldError("LineOne", err, rft.LineOne)
	}
	if err := rft.isAlphanumeric(rft.LineTwo); err != nil {
		return fieldError("LineTwo", err, rft.LineTwo)
	}
	if err := rft.isAlphanumeric(rft.LineThree); err != nil {
		return fieldError("LineThree", err, rft.LineThree)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (rft *RemittanceFreeText) LineOneField() string {
	return rft.alphaField(rft.LineOne, 140)
}

// LineTwoField gets a string of the LineTwo field
func (rft *RemittanceFreeText) LineTwoField() string {
	return rft.alphaField(rft.LineTwo, 140)
}

// LineThreeField gets a string of the LineThree field
func (rft *RemittanceFreeText) LineThreeField() string {
	return rft.alphaField(rft.LineThree, 140)
}
