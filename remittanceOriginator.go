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

// RemittanceOriginator is remittance originator
type RemittanceOriginator struct {
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
	// ContactName
	ContactName string `json:"contactName,omitempty"`
	// ContactPhoneNumber
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	// ContactMobileNumber
	ContactMobileNumber string `json:"contactMobileNumber,omitempty"`
	// ContactFaxNumber
	ContactFaxNumber string `json:"contactFaxNumber,omitempty"`
	// ContactElectronicAddress ( i.e., E-mail or URL address)
	ContactElectronicAddress string `json:"contactElectronicAddress,omitempty"`
	// ContactOther
	ContactOther string `json:"contactOther,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceOriginator returns a new RemittanceOriginator
func NewRemittanceOriginator() *RemittanceOriginator {
	ro := &RemittanceOriginator{
		tag: TagRemittanceOriginator,
	}
	return ro
}

// Parse takes the input string and parses the RemittanceOriginator values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ro *RemittanceOriginator) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 14 || dataLen > 3469 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [14, 3469] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}

	ro.tag = record[:6]
	ro.IdentificationType = ro.parseStringField(record[6:8])
	ro.IdentificationCode = ro.parseStringField(record[8:12])

	optionalFields := strings.Split(record[12:], "*")
	if len(optionalFields) >= 1 {
		ro.RemittanceData.Name = ro.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		ro.IdentificationNumber = ro.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		ro.IdentificationNumberIssuer = ro.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		ro.RemittanceData.DateBirthPlace = ro.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		ro.RemittanceData.AddressType = ro.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		ro.RemittanceData.Department = ro.parseStringField(optionalFields[5])
	}
	if len(optionalFields) >= 7 {
		ro.RemittanceData.SubDepartment = ro.parseStringField(optionalFields[6])
	}
	if len(optionalFields) >= 8 {
		ro.RemittanceData.StreetName = ro.parseStringField(optionalFields[7])
	}
	if len(optionalFields) >= 9 {
		ro.RemittanceData.BuildingNumber = ro.parseStringField(optionalFields[8])
	}
	if len(optionalFields) >= 10 {
		ro.RemittanceData.PostCode = ro.parseStringField(optionalFields[9])
	}
	if len(optionalFields) >= 11 {
		ro.RemittanceData.TownName = ro.parseStringField(optionalFields[10])
	}
	if len(optionalFields) >= 12 {
		ro.RemittanceData.CountrySubDivisionState = ro.parseStringField(optionalFields[11])
	}
	if len(optionalFields) >= 13 {
		ro.RemittanceData.Country = ro.parseStringField(optionalFields[12])
	}
	if len(optionalFields) >= 14 {
		ro.RemittanceData.AddressLineOne = ro.parseStringField(optionalFields[13])
	}
	if len(optionalFields) >= 15 {
		ro.RemittanceData.AddressLineTwo = ro.parseStringField(optionalFields[14])
	}
	if len(optionalFields) >= 16 {
		ro.RemittanceData.AddressLineThree = ro.parseStringField(optionalFields[15])
	}
	if len(optionalFields) >= 17 {
		ro.RemittanceData.AddressLineFour = ro.parseStringField(optionalFields[16])
	}
	if len(optionalFields) >= 18 {
		ro.RemittanceData.AddressLineFive = ro.parseStringField(optionalFields[17])
	}
	if len(optionalFields) >= 19 {
		ro.RemittanceData.AddressLineSix = ro.parseStringField(optionalFields[18])
	}
	if len(optionalFields) >= 20 {
		ro.RemittanceData.AddressLineSeven = ro.parseStringField(optionalFields[19])
	}
	if len(optionalFields) >= 21 {
		ro.RemittanceData.CountryOfResidence = ro.parseStringField(optionalFields[20])
	}
	if len(optionalFields) >= 22 {
		ro.ContactName = ro.parseStringField(optionalFields[21])
	}
	if len(optionalFields) >= 23 {
		ro.ContactPhoneNumber = ro.parseStringField(optionalFields[22])
	}
	if len(optionalFields) >= 24 {
		ro.ContactMobileNumber = ro.parseStringField(optionalFields[23])
	}
	if len(optionalFields) >= 25 {
		ro.ContactFaxNumber = ro.parseStringField(optionalFields[24])
	}
	if len(optionalFields) >= 26 {
		ro.ContactElectronicAddress = ro.parseStringField(optionalFields[25])
	}
	if len(optionalFields) >= 27 {
		ro.ContactOther = ro.parseStringField(optionalFields[26])
	}
	return nil
}

func (ro *RemittanceOriginator) UnmarshalJSON(data []byte) error {
	type Alias RemittanceOriginator
	aux := struct {
		*Alias
	}{
		(*Alias)(ro),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ro.tag = TagRemittanceOriginator
	return nil
}

// String writes RemittanceOriginator
func (ro *RemittanceOriginator) String() string {
	var buf strings.Builder
	buf.Grow(3442)
	buf.WriteString(ro.tag)
	buf.WriteString(ro.IdentificationTypeField())
	buf.WriteString(ro.IdentificationCodeField())
	buf.WriteString(strings.TrimSpace(ro.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.IdentificationNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.IdentificationNumberIssuerField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.DateBirthPlaceField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressTypeField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.DepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.SubDepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.StreetNameField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.BuildingNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.PostCodeField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.TownNameField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.CountrySubDivisionStateField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.CountryField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineSixField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.AddressLineSevenField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.CountryOfResidenceField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactNameField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactPhoneNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactMobileNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactFaxNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactElectronicAddressField()) + "*")
	buf.WriteString(strings.TrimSpace(ro.ContactOtherField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceOriginator and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// * Identification Type, Identification Code and Name are mandatory.
// * Identification Number is mandatory for all Identification Codes except PICDateBirthPlace.
// * Identification Number is not permitted for Identification Code PICDateBirthPlace.
// * Identification Number Issuer is not permitted for Identification Code OICSWIFTBICORBEI and PICDateBirthPlace.
// * Date & Place of Birth is only permitted for Identification Code PICDateBirthPlace.
func (ro *RemittanceOriginator) Validate() error { //nolint:gocyclo
	if err := ro.fieldInclusion(); err != nil {
		return err
	}
	if ro.tag != TagRemittanceOriginator {
		return fieldError("tag", ErrValidTagForType, ro.tag)
	}
	if err := ro.isIdentificationType(ro.IdentificationType); err != nil {
		return fieldError("IdentificationType", err, ro.IdentificationType)
	}

	switch ro.IdentificationType {
	case OrganizationID:
		if err := ro.isOrganizationIdentificationCode(ro.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, ro.IdentificationCode)
		}

	case PrivateID:
		if err := ro.isPrivateIdentificationCode(ro.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, ro.IdentificationCode)
		}
	}

	if err := ro.isAlphanumeric(ro.IdentificationNumber); err != nil {
		return fieldError("IdentificationNumber", err, ro.IdentificationNumber)
	}
	if err := ro.isAlphanumeric(ro.IdentificationNumberIssuer); err != nil {
		return fieldError("IdentificationNumberIssuer", err, ro.IdentificationNumberIssuer)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Name); err != nil {
		return fieldError("Name", err, ro.RemittanceData.Name)
	}
	if err := ro.isAddressType(ro.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, ro.RemittanceData.AddressType)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Department); err != nil {
		return fieldError("Department", err, ro.RemittanceData.Department)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, ro.RemittanceData.SubDepartment)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, ro.RemittanceData.StreetName)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, ro.RemittanceData.BuildingNumber)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, ro.RemittanceData.PostCode)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, ro.RemittanceData.TownName)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, ro.RemittanceData.CountrySubDivisionState)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Country); err != nil {
		return fieldError("Country", err, ro.RemittanceData.Country)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ro.RemittanceData.AddressLineOne)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ro.RemittanceData.AddressLineTwo)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ro.RemittanceData.AddressLineThree)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, ro.RemittanceData.AddressLineFour)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, ro.RemittanceData.AddressLineFive)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, ro.RemittanceData.AddressLineSix)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, ro.RemittanceData.AddressLineSeven)
	}

	if err := ro.isAlphanumeric(ro.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, ro.RemittanceData.CountryOfResidence)
	}
	if err := ro.isAlphanumeric(ro.ContactName); err != nil {
		return fieldError("ContactName", err, ro.ContactName)
	}
	if err := ro.isAlphanumeric(ro.ContactPhoneNumber); err != nil {
		return fieldError("ContactPhoneNumber", err, ro.ContactPhoneNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactMobileNumber); err != nil {
		return fieldError("ContactMobileNumber", err, ro.ContactMobileNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactFaxNumber); err != nil {
		return fieldError("ContactFaxNumber", err, ro.ContactFaxNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactElectronicAddress); err != nil {
		return fieldError("ContactElectronicAddress", err, ro.ContactElectronicAddress)
	}
	if err := ro.isAlphanumeric(ro.ContactOther); err != nil {
		return fieldError("ContactOther", err, ro.ContactOther)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ro *RemittanceOriginator) fieldInclusion() error {
	if ro.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}
	if ro.IdentificationCode == PICDateBirthPlace {
		if ro.IdentificationNumber != "" {
			return fieldError("IdentificationNumber", ErrInvalidProperty, ro.IdentificationNumber)
		}
	}

	if ro.IdentificationNumber == "" || ro.IdentificationCode == OICSWIFTBICORBEI ||
		ro.IdentificationCode == PICDateBirthPlace {
		if ro.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty, ro.IdentificationNumberIssuer)
		}
	}
	if ro.IdentificationCode != PICDateBirthPlace {
		if ro.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty, ro.RemittanceData.DateBirthPlace)
		}
	}
	return nil
}

// IdentificationTypeField gets a string of the IdentificationType field
func (ro *RemittanceOriginator) IdentificationTypeField() string {
	return ro.alphaField(ro.IdentificationType, 2)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ro *RemittanceOriginator) IdentificationCodeField() string {
	return ro.alphaField(ro.IdentificationCode, 4)
}

// NameField gets a string of the Name field
func (ro *RemittanceOriginator) NameField() string {
	return ro.alphaField(ro.RemittanceData.Name, 140)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (ro *RemittanceOriginator) IdentificationNumberField() string {
	return ro.alphaField(ro.IdentificationNumber, 35)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (ro *RemittanceOriginator) IdentificationNumberIssuerField() string {
	return ro.alphaField(ro.IdentificationNumberIssuer, 35)
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (ro *RemittanceOriginator) DateBirthPlaceField() string {
	return ro.alphaField(ro.RemittanceData.DateBirthPlace, 82)
}

// AddressTypeField gets a string of the AddressType field
func (ro *RemittanceOriginator) AddressTypeField() string {
	return ro.alphaField(ro.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (ro *RemittanceOriginator) DepartmentField() string {
	return ro.alphaField(ro.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (ro *RemittanceOriginator) SubDepartmentField() string {
	return ro.alphaField(ro.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (ro *RemittanceOriginator) StreetNameField() string {
	return ro.alphaField(ro.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (ro *RemittanceOriginator) BuildingNumberField() string {
	return ro.alphaField(ro.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (ro *RemittanceOriginator) PostCodeField() string {
	return ro.alphaField(ro.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (ro *RemittanceOriginator) TownNameField() string {
	return ro.alphaField(ro.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (ro *RemittanceOriginator) CountrySubDivisionStateField() string {
	return ro.alphaField(ro.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (ro *RemittanceOriginator) CountryField() string {
	return ro.alphaField(ro.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (ro *RemittanceOriginator) AddressLineOneField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (ro *RemittanceOriginator) AddressLineTwoField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (ro *RemittanceOriginator) AddressLineThreeField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (ro *RemittanceOriginator) AddressLineFourField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (ro *RemittanceOriginator) AddressLineFiveField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (ro *RemittanceOriginator) AddressLineSixField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (ro *RemittanceOriginator) AddressLineSevenField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineSeven, 70)
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (ro *RemittanceOriginator) CountryOfResidenceField() string {
	return ro.alphaField(ro.RemittanceData.CountryOfResidence, 2)
}

// ContactNameField gets a string of the ContactName field
func (ro *RemittanceOriginator) ContactNameField() string {
	return ro.alphaField(ro.ContactName, 140)
}

// ContactPhoneNumberField gets a string of the ContactPhoneNumber field
func (ro *RemittanceOriginator) ContactPhoneNumberField() string {
	return ro.alphaField(ro.ContactPhoneNumber, 35)
}

// ContactMobileNumberField gets a string of the ContactMobileNumber field
func (ro *RemittanceOriginator) ContactMobileNumberField() string {
	return ro.alphaField(ro.ContactMobileNumber, 35)
}

// ContactFaxNumberField gets a string of the ContactFaxNumber field
func (ro *RemittanceOriginator) ContactFaxNumberField() string {
	return ro.alphaField(ro.ContactFaxNumber, 35)
}

// ContactElectronicAddressField gets a string of the ContactElectronicAddress field
func (ro *RemittanceOriginator) ContactElectronicAddressField() string {
	return ro.alphaField(ro.ContactElectronicAddress, 2048)
}

// ContactOtherField gets a string of the ContactOther field
func (ro *RemittanceOriginator) ContactOtherField() string {
	return ro.alphaField(ro.ContactOther, 35)
}
