// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// OrderingCustomer is the ordering customer
type OrderingCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingCustomer returns a new OrderingCustomer
func NewOrderingCustomer() *OrderingCustomer {
	oc := &OrderingCustomer{
		tag: TagOrderingCustomer,
	}
	return oc
}

// Parse takes the input string and parses the OrderingCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oc *OrderingCustomer) Parse(record string) error {
	oc.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		oc.CoverPayment.SwiftFieldTag = oc.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		oc.CoverPayment.SwiftLineOne = oc.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		oc.CoverPayment.SwiftLineTwo = oc.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		oc.CoverPayment.SwiftLineThree = oc.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		oc.CoverPayment.SwiftLineFour = oc.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		oc.CoverPayment.SwiftLineFive = oc.parseStringField(optionalFields[5])
	}
	return nil
}

func (oc *OrderingCustomer) UnmarshalJSON(data []byte) error {
	type Alias OrderingCustomer
	aux := struct {
		*Alias
	}{
		(*Alias)(oc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	oc.tag = TagOrderingCustomer
	return nil
}

// String writes OrderingCustomer
func (oc *OrderingCustomer) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oc.tag)
	buf.WriteString(strings.TrimSpace(oc.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(oc.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(oc.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(oc.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(oc.SwiftLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(oc.SwiftLineFiveField()) + "*")
	return oc.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on OrderingCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oc *OrderingCustomer) Validate() error {
	if err := oc.fieldInclusion(); err != nil {
		return err
	}
	if oc.tag != TagOrderingCustomer {
		return fieldError("tag", ErrValidTagForType, oc.tag)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oc.CoverPayment.SwiftFieldTag)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oc.CoverPayment.SwiftLineOne)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oc.CoverPayment.SwiftLineTwo)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oc.CoverPayment.SwiftLineThree)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oc.CoverPayment.SwiftLineFour)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oc.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oc *OrderingCustomer) fieldInclusion() error {
	if oc.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, oc.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (oc *OrderingCustomer) SwiftFieldTagField() string {
	return oc.alphaField(oc.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oc *OrderingCustomer) SwiftLineOneField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oc *OrderingCustomer) SwiftLineTwoField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oc *OrderingCustomer) SwiftLineThreeField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oc *OrderingCustomer) SwiftLineFourField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oc *OrderingCustomer) SwiftLineFiveField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineFive, 35)
}
