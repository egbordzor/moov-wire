// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// FIReceiverFI is the financial institution receiver financial institution
type FIReceiverFI struct {
	// tag
	tag string
	// FIToFI is financial institution to financial institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIReceiverFI returns a new FIReceiverFI
func NewFIReceiverFI() *FIReceiverFI {
	firfi := &FIReceiverFI{
		tag: TagFIReceiverFI,
	}
	return firfi
}

// Parse takes the input string and parses the FIReceiverFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (firfi *FIReceiverFI) Parse(record string) error {
	firfi.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		firfi.FIToFI.LineOne = firfi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		firfi.FIToFI.LineTwo = firfi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		firfi.FIToFI.LineThree = firfi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		firfi.FIToFI.LineFour = firfi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		firfi.FIToFI.LineFive = firfi.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		firfi.FIToFI.LineSix = firfi.parseStringField(optionalFields[5])
	}
	return nil
}

func (firfi *FIReceiverFI) UnmarshalJSON(data []byte) error {
	type Alias FIReceiverFI
	aux := struct {
		*Alias
	}{
		(*Alias)(firfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	firfi.tag = TagFIReceiverFI
	return nil
}

// String writes FIReceiverFI
func (firfi *FIReceiverFI) String() string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(firfi.tag)
	buf.WriteString(strings.TrimSpace(firfi.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(firfi.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(firfi.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(firfi.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(firfi.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(firfi.LineSixField()) + "*")
	return firfi.cleanupDelimiters(buf.String())
}

func (firfi *FIReceiverFI) FullText(sep string) string {
	return firfi.FIToFI.FullText(sep)
}

// Validate performs WIRE format rule checks on FIReceiverFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (firfi *FIReceiverFI) Validate() error {
	if firfi.tag != TagFIReceiverFI {
		return fieldError("tag", ErrValidTagForType, firfi.tag)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, firfi.FIToFI.LineOne)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, firfi.FIToFI.LineTwo)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, firfi.FIToFI.LineThree)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, firfi.FIToFI.LineFour)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, firfi.FIToFI.LineFive)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, firfi.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (firfi *FIReceiverFI) LineOneField() string {
	return firfi.alphaField(firfi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (firfi *FIReceiverFI) LineTwoField() string {
	return firfi.alphaField(firfi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (firfi *FIReceiverFI) LineThreeField() string {
	return firfi.alphaField(firfi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (firfi *FIReceiverFI) LineFourField() string {
	return firfi.alphaField(firfi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (firfi *FIReceiverFI) LineFiveField() string {
	return firfi.alphaField(firfi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (firfi *FIReceiverFI) LineSixField() string {
	return firfi.alphaField(firfi.FIToFI.LineSix, 33)
}
