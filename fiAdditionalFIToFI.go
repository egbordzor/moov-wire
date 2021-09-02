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

// FIAdditionalFIToFI is the financial institution beneficiary financial institution
type FIAdditionalFIToFI struct {
	// tag
	tag string
	// AdditionalFiToFi is additional financial institution to financial institution information
	AdditionalFIToFI AdditionalFIToFI `json:"additionalFiToFi,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIAdditionalFIToFI returns a new FIAdditionalFIToFI
func NewFIAdditionalFIToFI() *FIAdditionalFIToFI {
	fifi := &FIAdditionalFIToFI{
		tag: TagFIAdditionalFIToFI,
	}
	return fifi
}

// Parse takes the input string and parses the FIAdditionalFIToFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fifi *FIAdditionalFIToFI) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 222 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 222] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	fifi.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		fifi.AdditionalFIToFI.LineOne = fifi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		fifi.AdditionalFIToFI.LineTwo = fifi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		fifi.AdditionalFIToFI.LineThree = fifi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		fifi.AdditionalFIToFI.LineFour = fifi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		fifi.AdditionalFIToFI.LineFive = fifi.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		fifi.AdditionalFIToFI.LineSix = fifi.parseStringField(optionalFields[5])
	}
	return nil
}

func (fifi *FIAdditionalFIToFI) UnmarshalJSON(data []byte) error {
	type Alias FIAdditionalFIToFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fifi.tag = TagFIAdditionalFIToFI
	return nil
}

// String writes FIAdditionalFIToFI
func (fifi *FIAdditionalFIToFI) String() string {
	var buf strings.Builder
	buf.Grow(216)
	buf.WriteString(fifi.tag)
	buf.WriteString(strings.TrimSpace(fifi.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(fifi.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(fifi.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(fifi.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(fifi.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(fifi.LineSixField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on FIAdditionalFIToFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fifi *FIAdditionalFIToFI) Validate() error {
	if fifi.tag != TagFIAdditionalFIToFI {
		return fieldError("tag", ErrValidTagForType, fifi.tag)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fifi.AdditionalFIToFI.LineOne)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fifi.AdditionalFIToFI.LineTwo)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fifi.AdditionalFIToFI.LineThree)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fifi.AdditionalFIToFI.LineFour)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fifi.AdditionalFIToFI.LineFive)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fifi.AdditionalFIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fifi *FIAdditionalFIToFI) LineOneField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (fifi *FIAdditionalFIToFI) LineTwoField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (fifi *FIAdditionalFIToFI) LineThreeField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (fifi *FIAdditionalFIToFI) LineFourField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (fifi *FIAdditionalFIToFI) LineFiveField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (fifi *FIAdditionalFIToFI) LineSixField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineSix, 35)
}
