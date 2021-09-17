// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// FIIntermediaryFI is the financial institution intermediary financial institution
type FIIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFI returns a new FIIntermediaryFI
func NewFIIntermediaryFI() *FIIntermediaryFI {
	fiifi := &FIIntermediaryFI{
		tag: TagFIIntermediaryFI,
	}
	return fiifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifi *FIIntermediaryFI) Parse(record string) error {
	fiifi.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		fiifi.FIToFI.LineOne = fiifi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		fiifi.FIToFI.LineTwo = fiifi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		fiifi.FIToFI.LineThree = fiifi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		fiifi.FIToFI.LineFour = fiifi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		fiifi.FIToFI.LineFive = fiifi.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		fiifi.FIToFI.LineSix = fiifi.parseStringField(optionalFields[5])
	}
	return nil
}

func (fiifi *FIIntermediaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIIntermediaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fiifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fiifi.tag = TagFIIntermediaryFI
	return nil
}

// String writes FIIntermediaryFI
func (fiifi *FIIntermediaryFI) String() string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(fiifi.tag)
	buf.WriteString(strings.TrimSpace(fiifi.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifi.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifi.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifi.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifi.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifi.LineSixField()) + "*")
	return fiifi.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on FIIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifi *FIIntermediaryFI) Validate() error {
	if fiifi.tag != TagFIIntermediaryFI {
		return fieldError("tag", ErrValidTagForType, fiifi.tag)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fiifi.FIToFI.LineOne)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifi.FIToFI.LineTwo)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fiifi.FIToFI.LineThree)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fiifi.FIToFI.LineFour)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fiifi.FIToFI.LineFive)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fiifi.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fiifi *FIIntermediaryFI) LineOneField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fiifi *FIIntermediaryFI) LineTwoField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fiifi *FIIntermediaryFI) LineThreeField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fiifi *FIIntermediaryFI) LineFourField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fiifi *FIIntermediaryFI) LineFiveField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fiifi *FIIntermediaryFI) LineSixField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineSix, 33)
}
