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

// OrderingInstitution is the ordering institution
type OrderingInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingInstitution returns a new OrderingInstitution
func NewOrderingInstitution() *OrderingInstitution {
	oi := &OrderingInstitution{
		tag: TagOrderingInstitution,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 192 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 192] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	oi.tag = record[:6]
	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		oi.CoverPayment.SwiftFieldTag = oi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		oi.CoverPayment.SwiftLineOne = oi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		oi.CoverPayment.SwiftLineTwo = oi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		oi.CoverPayment.SwiftLineThree = oi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		oi.CoverPayment.SwiftLineFour = oi.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		oi.CoverPayment.SwiftLineFive = oi.parseStringField(optionalFields[5])
	}
	return nil
}

func (oi *OrderingInstitution) UnmarshalJSON(data []byte) error {
	type Alias OrderingInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(oi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	oi.tag = TagOrderingInstitution
	return nil
}

// String writes OrderingInstitution
func (oi *OrderingInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oi.tag)
	buf.WriteString(strings.TrimSpace(oi.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(oi.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(oi.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(oi.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(oi.SwiftLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(oi.SwiftLineFiveField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oi *OrderingInstitution) Validate() error {
	if err := oi.fieldInclusion(); err != nil {
		return err
	}
	if oi.tag != TagOrderingInstitution {
		return fieldError("tag", ErrValidTagForType, oi.tag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oi.CoverPayment.SwiftFieldTag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oi.CoverPayment.SwiftLineOne)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oi.CoverPayment.SwiftLineTwo)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oi.CoverPayment.SwiftLineThree)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oi.CoverPayment.SwiftLineFour)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oi.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oi *OrderingInstitution) fieldInclusion() error {
	if oi.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, oi.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (oi *OrderingInstitution) SwiftFieldTagField() string {
	return oi.alphaField(oi.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oi *OrderingInstitution) SwiftLineOneField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oi *OrderingInstitution) SwiftLineTwoField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oi *OrderingInstitution) SwiftLineThreeField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oi *OrderingInstitution) SwiftLineFourField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oi *OrderingInstitution) SwiftLineFiveField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineFive, 35)
}
