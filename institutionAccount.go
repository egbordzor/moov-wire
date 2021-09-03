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

// InstitutionAccount is the institution account
type InstitutionAccount struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstitutionAccount returns a new InstitutionAccount
func NewInstitutionAccount() *InstitutionAccount {
	iAccount := &InstitutionAccount{
		tag: TagInstitutionAccount,
	}
	return iAccount
}

// Parse takes the input string and parses the InstitutionAccount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (iAccount *InstitutionAccount) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 192 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 192] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	iAccount.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		iAccount.CoverPayment.SwiftFieldTag = iAccount.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		iAccount.CoverPayment.SwiftLineOne = iAccount.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		iAccount.CoverPayment.SwiftLineTwo = iAccount.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		iAccount.CoverPayment.SwiftLineThree = iAccount.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		iAccount.CoverPayment.SwiftLineFour = iAccount.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		iAccount.CoverPayment.SwiftLineFive = iAccount.parseStringField(optionalFields[5])
	}
	return nil
}

func (iAccount *InstitutionAccount) UnmarshalJSON(data []byte) error {
	type Alias InstitutionAccount
	aux := struct {
		*Alias
	}{
		(*Alias)(iAccount),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	iAccount.tag = TagInstitutionAccount
	return nil
}

// String writes InstitutionAccount
func (iAccount *InstitutionAccount) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(iAccount.tag)
	buf.WriteString(strings.TrimSpace(iAccount.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(iAccount.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(iAccount.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(iAccount.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(iAccount.SwiftLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(iAccount.SwiftLineFiveField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on InstitutionAccount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iAccount *InstitutionAccount) Validate() error {
	if err := iAccount.fieldInclusion(); err != nil {
		return err
	}
	if iAccount.tag != TagInstitutionAccount {
		return fieldError("tag", ErrValidTagForType, iAccount.tag)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, iAccount.CoverPayment.SwiftFieldTag)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, iAccount.CoverPayment.SwiftLineOne)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, iAccount.CoverPayment.SwiftLineTwo)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, iAccount.CoverPayment.SwiftLineThree)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, iAccount.CoverPayment.SwiftLineFour)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, iAccount.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (iAccount *InstitutionAccount) fieldInclusion() error {
	if iAccount.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, iAccount.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (iAccount *InstitutionAccount) SwiftFieldTagField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (iAccount *InstitutionAccount) SwiftLineOneField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (iAccount *InstitutionAccount) SwiftLineTwoField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (iAccount *InstitutionAccount) SwiftLineThreeField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (iAccount *InstitutionAccount) SwiftLineFourField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (iAccount *InstitutionAccount) SwiftLineFiveField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFive, 35)
}
