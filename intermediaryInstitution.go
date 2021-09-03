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

// IntermediaryInstitution is the intermediary institution
type IntermediaryInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewIntermediaryInstitution returns a new IntermediaryInstitution
func NewIntermediaryInstitution() *IntermediaryInstitution {
	ii := &IntermediaryInstitution{
		tag: TagIntermediaryInstitution,
	}
	return ii
}

// Parse takes the input string and parses the IntermediaryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ii *IntermediaryInstitution) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 192 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 192] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	ii.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		ii.CoverPayment.SwiftFieldTag = ii.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		ii.CoverPayment.SwiftLineOne = ii.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		ii.CoverPayment.SwiftLineTwo = ii.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		ii.CoverPayment.SwiftLineThree = ii.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		ii.CoverPayment.SwiftLineFour = ii.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		ii.CoverPayment.SwiftLineFive = ii.parseStringField(optionalFields[5])
	}
	return nil
}

func (ii *IntermediaryInstitution) UnmarshalJSON(data []byte) error {
	type Alias IntermediaryInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(ii),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ii.tag = TagIntermediaryInstitution
	return nil
}

// String writes IntermediaryInstitution
func (ii *IntermediaryInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(ii.tag)
	buf.WriteString(strings.TrimSpace(ii.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(ii.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(ii.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(ii.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(ii.SwiftLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(ii.SwiftLineFiveField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on IntermediaryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ii *IntermediaryInstitution) Validate() error {
	if err := ii.fieldInclusion(); err != nil {
		return err
	}
	if ii.tag != TagIntermediaryInstitution {
		return fieldError("tag", ErrValidTagForType, ii.tag)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ii.CoverPayment.SwiftFieldTag)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ii.CoverPayment.SwiftLineOne)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ii.CoverPayment.SwiftLineTwo)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ii.CoverPayment.SwiftLineThree)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ii.CoverPayment.SwiftLineFour)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, ii.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ii *IntermediaryInstitution) fieldInclusion() error {
	if ii.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, ii.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ii *IntermediaryInstitution) SwiftFieldTagField() string {
	return ii.alphaField(ii.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ii *IntermediaryInstitution) SwiftLineOneField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ii *IntermediaryInstitution) SwiftLineTwoField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ii *IntermediaryInstitution) SwiftLineThreeField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ii *IntermediaryInstitution) SwiftLineFourField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (ii *IntermediaryInstitution) SwiftLineFiveField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFive, 35)
}
