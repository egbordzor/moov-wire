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

// RelatedRemittance is related remittance
type RelatedRemittance struct {
	// tag
	tag string
	// RemittanceIdentification is remittance identification
	RemittanceIdentification string `json:"remittanceIdentification,omitempty"`
	// RemittanceLocationMethod is  remittance location method
	RemittanceLocationMethod string `json:"remittanceLocationMethod,omitempty"`
	// RemittanceLocationElectronicAddress (E-mail or URL address)
	RemittanceLocationElectronicAddress string `json:"remittanceLocationElctronicAddress,omitempty"`
	// RemittanceData is RemittanceData
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRelatedRemittance returns a new RelatedRemittance
func NewRelatedRemittance() *RelatedRemittance {
	rr := &RelatedRemittance{
		tag: TagRelatedRemittance,
	}
	return rr
}

// Parse takes the input string and parses the RelatedRemittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rr *RelatedRemittance) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 7 || dataLen > 3061 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [7, 3061] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	rr.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		rr.RemittanceIdentification = rr.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		rr.RemittanceLocationMethod = rr.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		rr.RemittanceLocationElectronicAddress = rr.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		rr.RemittanceData.Name = rr.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		rr.RemittanceData.AddressType = rr.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		rr.RemittanceData.Department = rr.parseStringField(optionalFields[5])
	}
	if len(optionalFields) >= 7 {
		rr.RemittanceData.SubDepartment = rr.parseStringField(optionalFields[6])
	}
	if len(optionalFields) >= 8 {
		rr.RemittanceData.StreetName = rr.parseStringField(optionalFields[7])
	}
	if len(optionalFields) >= 9 {
		rr.RemittanceData.BuildingNumber = rr.parseStringField(optionalFields[8])
	}
	if len(optionalFields) >= 10 {
		rr.RemittanceData.PostCode = rr.parseStringField(optionalFields[9])
	}
	if len(optionalFields) >= 11 {
		rr.RemittanceData.TownName = rr.parseStringField(optionalFields[10])
	}
	if len(optionalFields) >= 12 {
		rr.RemittanceData.CountrySubDivisionState = rr.parseStringField(optionalFields[11])
	}
	if len(optionalFields) >= 13 {
		rr.RemittanceData.Country = rr.parseStringField(optionalFields[12])
	}
	if len(optionalFields) >= 14 {
		rr.RemittanceData.AddressLineOne = rr.parseStringField(optionalFields[13])
	}
	if len(optionalFields) >= 15 {
		rr.RemittanceData.AddressLineTwo = rr.parseStringField(optionalFields[14])
	}
	if len(optionalFields) >= 16 {
		rr.RemittanceData.AddressLineThree = rr.parseStringField(optionalFields[15])
	}
	if len(optionalFields) >= 17 {
		rr.RemittanceData.AddressLineFour = rr.parseStringField(optionalFields[16])
	}
	if len(optionalFields) >= 18 {
		rr.RemittanceData.AddressLineFive = rr.parseStringField(optionalFields[17])
	}
	if len(optionalFields) >= 19 {
		rr.RemittanceData.AddressLineSix = rr.parseStringField(optionalFields[18])
	}
	if len(optionalFields) >= 20 {
		rr.RemittanceData.AddressLineSeven = rr.parseStringField(optionalFields[19])
	}
	return nil
}

func (rr *RelatedRemittance) UnmarshalJSON(data []byte) error {
	type Alias RelatedRemittance
	aux := struct {
		*Alias
	}{
		(*Alias)(rr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rr.tag = TagRelatedRemittance
	return nil
}

// String writes RelatedRemittance
func (rr *RelatedRemittance) String() string {
	var buf strings.Builder
	buf.Grow(3041)
	buf.WriteString(rr.tag)
	buf.WriteString(strings.TrimSpace(rr.RemittanceIdentificationField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.RemittanceLocationMethodField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.RemittanceLocationElectronicAddressField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.NameField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressTypeField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.DepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.SubDepartmentField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.StreetNameField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.BuildingNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.PostCodeField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.TownNameField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.CountrySubDivisionStateField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.CountryField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineSixField()) + "*")
	buf.WriteString(strings.TrimSpace(rr.AddressLineSevenField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on RelatedRemittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rr *RelatedRemittance) Validate() error {
	if rr.tag != TagRelatedRemittance {
		return fieldError("tag", ErrValidTagForType, rr.tag)
	}
	if err := rr.fieldInclusion(); err != nil {
		return err
	}
	if err := rr.isAlphanumeric(rr.RemittanceIdentification); err != nil {
		return fieldError("RemittanceIdentification", err, rr.RemittanceIdentification)
	}
	if err := rr.isRemittanceLocationMethod(rr.RemittanceLocationMethod); err != nil {
		return fieldError("RemittanceLocationMethod", err, rr.RemittanceLocationMethod)
	}
	if err := rr.isAlphanumeric(rr.RemittanceLocationElectronicAddress); err != nil {
		return fieldError("RemittanceLocationElectronicAddress", err, rr.RemittanceLocationElectronicAddress)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Name); err != nil {
		return fieldError("Name", err, rr.RemittanceData.Name)
	}
	if err := rr.isAddressType(rr.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, rr.RemittanceData.AddressType)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Department); err != nil {
		return fieldError("Department", err, rr.RemittanceData.Department)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, rr.RemittanceData.SubDepartment)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, rr.RemittanceData.StreetName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, rr.RemittanceData.BuildingNumber)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, rr.RemittanceData.PostCode)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, rr.RemittanceData.TownName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, rr.RemittanceData.CountrySubDivisionState)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Country); err != nil {
		return fieldError("Country", err, rr.RemittanceData.Country)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, rr.RemittanceData.AddressLineOne)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, rr.RemittanceData.AddressLineTwo)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, rr.RemittanceData.AddressLineThree)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, rr.RemittanceData.AddressLineFour)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, rr.RemittanceData.AddressLineFive)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, rr.RemittanceData.AddressLineSix)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, rr.RemittanceData.AddressLineSeven)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, rr.RemittanceData.CountryOfResidence)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rr *RelatedRemittance) fieldInclusion() error {
	return nil
}

// RemittanceIdentificationField gets a string of the RemittanceIdentification field
func (rr *RelatedRemittance) RemittanceIdentificationField() string {
	return rr.alphaField(rr.RemittanceIdentification, 35)
}

// RemittanceLocationMethodField gets a string of the RemittanceLocationMethod field
func (rr *RelatedRemittance) RemittanceLocationMethodField() string {
	return rr.alphaField(rr.RemittanceLocationMethod, 4)
}

// RemittanceLocationElectronicAddressField gets a string of the RemittanceLocationElectronicAddress field
func (rr *RelatedRemittance) RemittanceLocationElectronicAddressField() string {
	return rr.alphaField(rr.RemittanceLocationElectronicAddress, 2048)
}

// NameField gets a string of the Name field
func (rr *RelatedRemittance) NameField() string {
	return rr.alphaField(rr.RemittanceData.Name, 140)
}

// AddressTypeField gets a string of the AddressType field
func (rr *RelatedRemittance) AddressTypeField() string {
	return rr.alphaField(rr.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (rr *RelatedRemittance) DepartmentField() string {
	return rr.alphaField(rr.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (rr *RelatedRemittance) SubDepartmentField() string {
	return rr.alphaField(rr.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (rr *RelatedRemittance) StreetNameField() string {
	return rr.alphaField(rr.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rr *RelatedRemittance) BuildingNumberField() string {
	return rr.alphaField(rr.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (rr *RelatedRemittance) PostCodeField() string {
	return rr.alphaField(rr.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (rr *RelatedRemittance) TownNameField() string {
	return rr.alphaField(rr.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rr *RelatedRemittance) CountrySubDivisionStateField() string {
	return rr.alphaField(rr.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (rr *RelatedRemittance) CountryField() string {
	return rr.alphaField(rr.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rr *RelatedRemittance) AddressLineOneField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rr *RelatedRemittance) AddressLineTwoField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rr *RelatedRemittance) AddressLineThreeField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rr *RelatedRemittance) AddressLineFourField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rr *RelatedRemittance) AddressLineFiveField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rr *RelatedRemittance) AddressLineSixField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rr *RelatedRemittance) AddressLineSevenField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineSeven, 70)
}
