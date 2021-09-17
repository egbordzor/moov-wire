// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// FIBeneficiaryFI is the financial institution beneficiary financial institution
type FIBeneficiaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFI returns a new FIBeneficiaryFI
func NewFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := &FIBeneficiaryFI{
		tag: TagFIBeneficiaryFI,
	}
	return fibfi
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfi *FIBeneficiaryFI) Parse(record string) error {
	fibfi.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		fibfi.FIToFI.LineOne = fibfi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		fibfi.FIToFI.LineTwo = fibfi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		fibfi.FIToFI.LineThree = fibfi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		fibfi.FIToFI.LineFour = fibfi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		fibfi.FIToFI.LineFive = fibfi.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		fibfi.FIToFI.LineSix = fibfi.parseStringField(optionalFields[5])
	}
	return nil
}

func (fibfi *FIBeneficiaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfi.tag = TagFIBeneficiaryFI
	return nil
}

// String writes FIBeneficiaryFI
func (fibfi *FIBeneficiaryFI) String() string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(fibfi.tag)
	buf.WriteString(strings.TrimSpace(fibfi.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfi.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfi.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfi.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfi.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(fibfi.LineSixField()) + "*")
	return fibfi.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on FIBeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfi *FIBeneficiaryFI) Validate() error {
	if fibfi.tag != TagFIBeneficiaryFI {
		return fieldError("tag", ErrValidTagForType, fibfi.tag)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fibfi.FIToFI.LineOne)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfi.FIToFI.LineTwo)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fibfi.FIToFI.LineThree)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fibfi.FIToFI.LineFour)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fibfi.FIToFI.LineFive)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fibfi.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fibfi *FIBeneficiaryFI) LineOneField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fibfi *FIBeneficiaryFI) LineTwoField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fibfi *FIBeneficiaryFI) LineThreeField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fibfi *FIBeneficiaryFI) LineFourField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fibfi *FIBeneficiaryFI) LineFiveField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fibfi *FIBeneficiaryFI) LineSixField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineSix, 33)
}
