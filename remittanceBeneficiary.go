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

// RemittanceBeneficiary is remittance beneficiary
type RemittanceBeneficiary struct {
	// tag
	tag string
	// IdentificationType is identification type
	IdentificationType string `json:"identificationType,omitempty"`
	// IdentificationCode  Organization Identification Codes  * `BANK` - Bank Party Identification * `CUST` - Customer Number * `DUNS` - Data Universal Number System (Dun & Bradstreet) * `EMPL` - Employer Identification Number * `GS1G` - Global Location Number * `PROP` - Proprietary Identification Number * `SWBB` - SWIFT BIC or BEI * `TXID` - Tax Identification Number  Private Identification Codes  * `ARNU` - Alien Registration Number * `CCPT` - Passport Number * `CUST` - Customer Number * `DPOB` - Date & Place of Birth * `DRLC` - Driver’s License Number * `EMPL` - Employee Identification Number * `NIDN` - National Identity Number * `PROP` - Proprietary Identification Number * `SOSE` - Social Security Number * `TXID` - Tax Identification Number
	IdentificationCode string `json:"identificationCode,omitempty"`
	// IdentificationNumber
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// IdentificationNumberIssuer
	IdentificationNumberIssuer string `json:"identificationNumberIssuer,omitempty"`
	// RemittanceData
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceBeneficiary returns a new RemittanceBeneficiary
func NewRemittanceBeneficiary() *RemittanceBeneficiary {
	rb := &RemittanceBeneficiary{
		tag: TagRemittanceBeneficiary,
	}
	return rb
}

// Parse takes the input string and parses the RemittanceBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rb *RemittanceBeneficiary) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 1137 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 1137] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	rb.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		rb.RemittanceData.Name = rb.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		rb.IdentificationType = rb.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		rb.IdentificationCode = rb.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		rb.IdentificationNumber = rb.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		rb.IdentificationNumberIssuer = rb.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		rb.RemittanceData.DateBirthPlace = rb.parseStringField(optionalFields[5])
	}
	if len(optionalFields) >= 7 {
		rb.RemittanceData.AddressType = rb.parseStringField(optionalFields[6])
	}
	if len(optionalFields) >= 8 {
		rb.RemittanceData.Department = rb.parseStringField(optionalFields[7])
	}
	if len(optionalFields) >= 9 {
		rb.RemittanceData.SubDepartment = rb.parseStringField(optionalFields[8])
	}
	if len(optionalFields) >= 10 {
		rb.RemittanceData.StreetName = rb.parseStringField(optionalFields[9])
	}
	if len(optionalFields) >= 11 {
		rb.RemittanceData.BuildingNumber = rb.parseStringField(optionalFields[10])
	}
	if len(optionalFields) >= 12 {
		rb.RemittanceData.PostCode = rb.parseStringField(optionalFields[11])
	}
	if len(optionalFields) >= 13 {
		rb.RemittanceData.TownName = rb.parseStringField(optionalFields[12])
	}
	if len(optionalFields) >= 14 {
		rb.RemittanceData.CountrySubDivisionState = rb.parseStringField(optionalFields[13])
	}
	if len(optionalFields) >= 15 {
		rb.RemittanceData.Country = rb.parseStringField(optionalFields[14])
	}
	if len(optionalFields) >= 16 {
		rb.RemittanceData.AddressLineOne = rb.parseStringField(optionalFields[15])
	}
	if len(optionalFields) >= 17 {
		rb.RemittanceData.AddressLineTwo = rb.parseStringField(optionalFields[16])
	}
	if len(optionalFields) >= 18 {
		rb.RemittanceData.AddressLineThree = rb.parseStringField(optionalFields[17])
	}
	if len(optionalFields) >= 19 {
		rb.RemittanceData.AddressLineFour = rb.parseStringField(optionalFields[18])
	}
	if len(optionalFields) >= 20 {
		rb.RemittanceData.AddressLineFive = rb.parseStringField(optionalFields[19])
	}
	if len(optionalFields) >= 21 {
		rb.RemittanceData.AddressLineSix = rb.parseStringField(optionalFields[20])
	}
	if len(optionalFields) >= 22 {
		rb.RemittanceData.AddressLineSeven = rb.parseStringField(optionalFields[21])
	}
	if len(optionalFields) >= 23 {
		rb.RemittanceData.CountryOfResidence = rb.parseStringField(optionalFields[22])
	}
	return nil
}

func (rb *RemittanceBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias RemittanceBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(rb),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rb.tag = TagRemittanceBeneficiary
	return nil
}

// String writes RemittanceBeneficiary
func (rb *RemittanceBeneficiary) String() string {
	var buf strings.Builder
	buf.Grow(1114)
	buf.WriteString(rb.tag)
	buf.WriteString(strings.TrimSpace(rb.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.IdentificationTypeField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.IdentificationCodeField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.IdentificationNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.IdentificationNumberIssuerField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.DateBirthPlaceField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressTypeField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.DepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.SubDepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.StreetNameField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.BuildingNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.PostCodeField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.TownNameField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.CountrySubDivisionStateField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.CountryField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineSixField()) + "*")
	buf.WriteString(strings.TrimSpace(rb.AddressLineSevenField()) + "*")
	if rb.RemittanceData.CountryOfResidence != "" {
		buf.WriteString(strings.TrimSpace(rb.CountryOfResidenceField()) + "*")
	}
	return rb.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on RemittanceBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// * Name is mandatory.
// * Identification Number
//   * Not permitted unless Identification Type and Identification Code are present.
//   * Not permitted for Identification Code PICDateBirthPlace.
// * Identification Number Issuer
//   * Not permitted unless Identification Type, Identification Code and Identification Number are present.
//   * Not permitted for Identification Code SWBB and PICDateBirthPlace.
// * Date & Place of Birth is only permitted for Identification Code PICDateBirthPlace.
func (rb *RemittanceBeneficiary) Validate() error {
	if err := rb.fieldInclusion(); err != nil {
		return err
	}
	if rb.tag != TagRemittanceBeneficiary {
		return fieldError("tag", ErrValidTagForType, rb.tag)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Name); err != nil {
		return fieldError("Name", err, rb.RemittanceData.Name)
	}
	if err := rb.isIdentificationType(rb.IdentificationType); err != nil {
		return fieldError("IdentificationType", err, rb.IdentificationType)
	}
	switch rb.IdentificationType {
	case OrganizationID:
		if err := rb.isOrganizationIdentificationCode(rb.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, rb.IdentificationCode)
		}
	case PrivateID:
		if err := rb.isPrivateIdentificationCode(rb.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, rb.IdentificationCode)
		}
	}
	if err := rb.isAlphanumeric(rb.IdentificationNumber); err != nil {
		return fieldError("IdentificationNumber", err, rb.IdentificationNumber)
	}
	if err := rb.isAlphanumeric(rb.IdentificationNumberIssuer); err != nil {
		return fieldError("IdentificationNumberIssuer", err, rb.IdentificationNumberIssuer)
	}
	if err := rb.isAddressType(rb.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, rb.RemittanceData.AddressType)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Department); err != nil {
		return fieldError("Department", err, rb.RemittanceData.Department)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, rb.RemittanceData.SubDepartment)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, rb.RemittanceData.StreetName)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, rb.RemittanceData.BuildingNumber)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, rb.RemittanceData.PostCode)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, rb.RemittanceData.TownName)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, rb.RemittanceData.CountrySubDivisionState)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Country); err != nil {
		return fieldError("Country", err, rb.RemittanceData.Country)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, rb.RemittanceData.AddressLineOne)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, rb.RemittanceData.AddressLineTwo)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, rb.RemittanceData.AddressLineThree)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, rb.RemittanceData.AddressLineFour)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, rb.RemittanceData.AddressLineFive)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, rb.RemittanceData.AddressLineSix)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, rb.RemittanceData.AddressLineSeven)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, rb.RemittanceData.CountryOfResidence)
	}

	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rb *RemittanceBeneficiary) fieldInclusion() error {
	if rb.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}

	if rb.IdentificationCode == PICDateBirthPlace {
		if rb.IdentificationNumber != "" {
			return fieldError("IdentificationNumber", ErrInvalidProperty, rb.IdentificationNumber)
		}
	}
	if rb.IdentificationNumber == "" || rb.IdentificationCode == OICSWIFTBICORBEI ||
		rb.IdentificationCode == PICDateBirthPlace {
		if rb.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty, rb.IdentificationNumberIssuer)
		}
	}
	if rb.IdentificationCode != PICDateBirthPlace {
		if rb.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty, rb.RemittanceData.DateBirthPlace)
		}
	}

	return nil
}

// NameField gets a string of the Name field
func (rb *RemittanceBeneficiary) NameField() string {
	return rb.alphaField(rb.RemittanceData.Name, 140)
}

// IdentificationTypeField gets a string of the IdentificationType field
func (rb *RemittanceBeneficiary) IdentificationTypeField() string {
	return rb.alphaField(rb.IdentificationType, 2)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (rb *RemittanceBeneficiary) IdentificationCodeField() string {
	return rb.alphaField(rb.IdentificationCode, 4)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (rb *RemittanceBeneficiary) IdentificationNumberField() string {
	return rb.alphaField(rb.IdentificationNumber, 35)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (rb *RemittanceBeneficiary) IdentificationNumberIssuerField() string {
	return rb.alphaField(rb.IdentificationNumberIssuer, 35)
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (rb *RemittanceBeneficiary) DateBirthPlaceField() string {
	return rb.alphaField(rb.RemittanceData.DateBirthPlace, 82)
}

// AddressTypeField gets a string of the AddressType field
func (rb *RemittanceBeneficiary) AddressTypeField() string {
	return rb.alphaField(rb.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (rb *RemittanceBeneficiary) DepartmentField() string {
	return rb.alphaField(rb.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (rb *RemittanceBeneficiary) SubDepartmentField() string {
	return rb.alphaField(rb.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (rb *RemittanceBeneficiary) StreetNameField() string {
	return rb.alphaField(rb.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rb *RemittanceBeneficiary) BuildingNumberField() string {
	return rb.alphaField(rb.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (rb *RemittanceBeneficiary) PostCodeField() string {
	return rb.alphaField(rb.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (rb *RemittanceBeneficiary) TownNameField() string {
	return rb.alphaField(rb.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rb *RemittanceBeneficiary) CountrySubDivisionStateField() string {
	return rb.alphaField(rb.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (rb *RemittanceBeneficiary) CountryField() string {
	return rb.alphaField(rb.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rb *RemittanceBeneficiary) AddressLineOneField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rb *RemittanceBeneficiary) AddressLineTwoField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rb *RemittanceBeneficiary) AddressLineThreeField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rb *RemittanceBeneficiary) AddressLineFourField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rb *RemittanceBeneficiary) AddressLineFiveField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rb *RemittanceBeneficiary) AddressLineSixField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rb *RemittanceBeneficiary) AddressLineSevenField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineSeven, 70)
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (rb *RemittanceBeneficiary) CountryOfResidenceField() string {
	return rb.alphaField(rb.RemittanceData.CountryOfResidence, 2)
}
