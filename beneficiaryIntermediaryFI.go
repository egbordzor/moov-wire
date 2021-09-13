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

// BeneficiaryIntermediaryFI {4000}
type BeneficiaryIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryIntermediaryFI returns a new BeneficiaryIntermediaryFI
func NewBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	bifi := &BeneficiaryIntermediaryFI{
		tag: TagBeneficiaryIntermediaryFI,
	}
	return bifi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bifi *BeneficiaryIntermediaryFI) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 12 || dataLen > 186 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [12, 186] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	bifi.tag = record[:6]
	bifi.FinancialInstitution.IdentificationCode = bifi.parseStringField(record[6:7])

	optionalFields := strings.Split(record[7:], "*")
	if len(optionalFields) >= 1 {
		bifi.FinancialInstitution.Identifier = bifi.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		bifi.FinancialInstitution.Name = bifi.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		bifi.FinancialInstitution.Address.AddressLineOne = bifi.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		bifi.FinancialInstitution.Address.AddressLineTwo = bifi.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		bifi.FinancialInstitution.Address.AddressLineThree = bifi.parseStringField(optionalFields[4])
	}
	return nil
}

func (bifi *BeneficiaryIntermediaryFI) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryIntermediaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(bifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bifi.tag = TagBeneficiaryIntermediaryFI
	return nil
}

// String writes BeneficiaryIntermediaryFI
func (bifi *BeneficiaryIntermediaryFI) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(bifi.tag)
	buf.WriteString(bifi.IdentificationCodeField())
	buf.WriteString(strings.TrimSpace(bifi.IdentifierField()) + "*")
	buf.WriteString(strings.TrimSpace(bifi.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(bifi.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(bifi.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(bifi.AddressLineThreeField()) + "*")
	return bifi.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on BeneficiaryIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (bifi *BeneficiaryIntermediaryFI) Validate() error {
	if err := bifi.fieldInclusion(); err != nil {
		return err
	}
	if bifi.tag != TagBeneficiaryIntermediaryFI {
		return fieldError("tag", ErrValidTagForType, bifi.tag)
	}
	if err := bifi.isIdentificationCode(bifi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, bifi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch bifi.FinancialInstitution.IdentificationCode {
	case
		"B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, bifi.FinancialInstitution.Identifier)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, bifi.FinancialInstitution.Name)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, bifi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, bifi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, bifi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bifi *BeneficiaryIntermediaryFI) fieldInclusion() error {
	if bifi.FinancialInstitution.IdentificationCode != "" && bifi.FinancialInstitution.Identifier == "" {
		return fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.Identifier", ErrFieldRequired)
	}
	if bifi.FinancialInstitution.IdentificationCode == "" && bifi.FinancialInstitution.Identifier != "" {
		return fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.IdentificationCode", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (bifi *BeneficiaryIntermediaryFI) IdentificationCodeField() string {
	return bifi.alphaField(bifi.FinancialInstitution.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (bifi *BeneficiaryIntermediaryFI) IdentifierField() string {
	return bifi.alphaField(bifi.FinancialInstitution.Identifier, 34)
}

// NameField gets a string of the Name field
func (bifi *BeneficiaryIntermediaryFI) NameField() string {
	return bifi.alphaField(bifi.FinancialInstitution.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (bifi *BeneficiaryIntermediaryFI) AddressLineOneField() string {
	return bifi.alphaField(bifi.FinancialInstitution.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (bifi *BeneficiaryIntermediaryFI) AddressLineTwoField() string {
	return bifi.alphaField(bifi.FinancialInstitution.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (bifi *BeneficiaryIntermediaryFI) AddressLineThreeField() string {
	return bifi.alphaField(bifi.FinancialInstitution.Address.AddressLineThree, 35)
}
