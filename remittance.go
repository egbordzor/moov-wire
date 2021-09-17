// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// Remittance is the remittance information
type Remittance struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittance returns a new Remittance
func NewRemittance() *Remittance {
	ri := &Remittance{
		tag: TagRemittance,
	}
	return ri
}

// Parse takes the input string and parses the Remittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ri *Remittance) Parse(record string) error {
	ri.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		ri.CoverPayment.SwiftFieldTag = ri.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		ri.CoverPayment.SwiftLineOne = ri.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		ri.CoverPayment.SwiftLineTwo = ri.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		ri.CoverPayment.SwiftLineThree = ri.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		ri.CoverPayment.SwiftLineFour = ri.parseStringField(optionalFields[4])
	}
	return nil
}

func (ri *Remittance) UnmarshalJSON(data []byte) error {
	type Alias Remittance
	aux := struct {
		*Alias
	}{
		(*Alias)(ri),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ri.tag = TagRemittance
	return nil
}

// String writes Remittance
func (ri *Remittance) String() string {
	var buf strings.Builder
	buf.Grow(156)
	buf.WriteString(ri.tag)
	buf.WriteString(strings.TrimSpace(ri.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(ri.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(ri.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(ri.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(ri.SwiftLineFourField()) + "*")
	return ri.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on Remittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ri *Remittance) Validate() error {
	if err := ri.fieldInclusion(); err != nil {
		return err
	}
	if ri.tag != TagRemittance {
		return fieldError("tag", ErrValidTagForType, ri.tag)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ri.CoverPayment.SwiftFieldTag)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ri.CoverPayment.SwiftLineOne)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ri.CoverPayment.SwiftLineTwo)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ri.CoverPayment.SwiftLineThree)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ri.CoverPayment.SwiftLineFour)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ri *Remittance) fieldInclusion() error {
	if ri.CoverPayment.SwiftLineFive != "" {
		return fieldError("SwiftLineFive", ErrInvalidProperty, ri.CoverPayment.SwiftLineFive)
	}
	if ri.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, ri.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ri *Remittance) SwiftFieldTagField() string {
	return ri.alphaField(ri.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ri *Remittance) SwiftLineOneField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ri *Remittance) SwiftLineTwoField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ri *Remittance) SwiftLineThreeField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ri *Remittance) SwiftLineFourField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineFour, 35)
}
