// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// InstructingFI is the instructing financial institution
type InstructingFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstructingFI returns a new InstructingFI
func NewInstructingFI() *InstructingFI {
	ifi := &InstructingFI{
		tag: TagInstructingFI,
	}
	return ifi
}

// Parse takes the input string and parses the InstructingFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ifi *InstructingFI) Parse(record string) error {
	ifi.tag = record[:6]
	ifi.FinancialInstitution.IdentificationCode = ifi.parseStringField(record[6:7])

	optionalFields := strings.Split(record[7:], "*")
	if len(optionalFields) >= 1 {
		ifi.FinancialInstitution.Identifier = ifi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		ifi.FinancialInstitution.Name = ifi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		ifi.FinancialInstitution.Address.AddressLineOne = ifi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		ifi.FinancialInstitution.Address.AddressLineTwo = ifi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		ifi.FinancialInstitution.Address.AddressLineThree = ifi.parseStringField(optionalFields[4])
	}
	return nil
}

func (ifi *InstructingFI) UnmarshalJSON(data []byte) error {
	type Alias InstructingFI
	aux := struct {
		*Alias
	}{
		(*Alias)(ifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ifi.tag = TagInstructingFI
	return nil
}

// String writes InstructingFI
func (ifi *InstructingFI) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(ifi.tag)
	buf.WriteString(ifi.IdentificationCodeField())
	buf.WriteString(strings.TrimSpace(ifi.IdentifierField()) + "*")
	buf.WriteString(strings.TrimSpace(ifi.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(ifi.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(ifi.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(ifi.AddressLineThreeField()) + "*")
	return ifi.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on InstructingFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (ifi *InstructingFI) Validate() error {
	if ifi.tag != TagInstructingFI {
		return fieldError("tag", ErrValidTagForType, ifi.tag)
	}

	if err := ifi.fieldInclusion(); err != nil {
		return err
	}

	if err := ifi.isIdentificationCode(ifi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, ifi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch ifi.FinancialInstitution.IdentificationCode {
	case
		"", "B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, ifi.FinancialInstitution.IdentificationCode)
	}
	if err := ifi.isAlphanumeric(ifi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, ifi.FinancialInstitution.Identifier)
	}

	if err := ifi.isAlphanumeric(ifi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, ifi.FinancialInstitution.Name)
	}
	if err := ifi.isAlphanumeric(ifi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ifi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := ifi.isAlphanumeric(ifi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ifi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := ifi.isAlphanumeric(ifi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ifi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifi *InstructingFI) fieldInclusion() error {
	if ifi.FinancialInstitution.IdentificationCode != "" && ifi.FinancialInstitution.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if ifi.FinancialInstitution.IdentificationCode == "" && ifi.FinancialInstitution.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ifi *InstructingFI) IdentificationCodeField() string {
	return ifi.alphaField(ifi.FinancialInstitution.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (ifi *InstructingFI) IdentifierField() string {
	return ifi.alphaField(ifi.FinancialInstitution.Identifier, 34)
}

// NameField gets a string of the Name field
func (ifi *InstructingFI) NameField() string {
	return ifi.alphaField(ifi.FinancialInstitution.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (ifi *InstructingFI) AddressLineOneField() string {
	return ifi.alphaField(ifi.FinancialInstitution.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (ifi *InstructingFI) AddressLineTwoField() string {
	return ifi.alphaField(ifi.FinancialInstitution.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (ifi *InstructingFI) AddressLineThreeField() string {
	return ifi.alphaField(ifi.FinancialInstitution.Address.AddressLineThree, 35)
}
