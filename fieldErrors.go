// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"errors"
	"fmt"
)

var (
	// Errors specific to validation

	// ErrValidTagForType is returned when a type has an invalid tag
	ErrValidTagForType = errors.New("is an invalid tag")
	// ErrNonNumeric is returned when a field has non-numeric characters
	ErrNonNumeric = errors.New("has non numeric characters")
	// ErrNonAlphanumeric is returned when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
	// ErrNonAmount is returned for an incorrect wire amount format
	ErrNonAmount = errors.New("is an incorrect amount format")
	// ErrNonCurrencyCode is returned for an incorrect currency code
	ErrNonCurrencyCode = errors.New("is not a recognized currency code")
	// ErrUpperAlpha is returned when a field is not in uppercase
	ErrUpperAlpha = errors.New("is not uppercase A-Z or 0-9")
	// ErrFieldInclusion is returned when a field is mandatory and has a default value
	ErrFieldInclusion = errors.New("is a mandatory field and has a default value")
	// ErrConstructor is returned when there's a mandatory field is not initialized correctly, and prompts to use the constructor
	ErrConstructor = errors.New("is a mandatory field and has a default value. Use the constructor")
	// ErrFieldRequired is returned when a field is required
	ErrFieldRequired = errors.New("is a required field")
	// ErrNotPermitted is returned when a field is included, but not permitted in combination with the other tags present in the message
	ErrNotPermitted = errors.New("is not permitted in this context")
	// ErrValidMonth is returned for an invalid month
	ErrValidMonth = errors.New("is an invalid month")
	// ErrValidDay is returned for an invalid day
	ErrValidDay = errors.New("is an invalid day")
	// ErrValidYear is returned for an invalid year
	ErrValidYear = errors.New("is an invalid year")
	// ErrValidCentury is returned for an invalid century
	ErrValidCentury = errors.New("is an invalid century")
	// ErrValidDate is returned for an invalid date
	ErrValidDate = errors.New("is an invalid date format")
	// ErrInvalidProperty is returned for an invalid type property
	ErrInvalidProperty = errors.New("is an invalid property")

	// SenderSupplied Tag {1500}

	// ErrFormatVersion is returned for an invalid an invalid FormatVersion
	ErrFormatVersion = errors.New("is not 30")
	// ErrTestProductionCode is returned for an invalid TestProductionCode
	ErrTestProductionCode = errors.New("is an invalid test production code")
	// ErrMessageDuplicationCode is returned for an invalid MessageDuplicationCode
	ErrMessageDuplicationCode = errors.New("is an invalid message duplication code")

	// TypeSubType Tag {1510}

	// ErrTypeCode is returned for an invalid TypeCode tag
	ErrTypeCode = errors.New("is an invalid type code")
	// ErrSubTypeCode is returned when there's an invalid SubTypeCode tag
	ErrSubTypeCode = errors.New("is an invalid sub type Code")

	// BusinessFunctionCode Tag {3600}

	// ErrBusinessFunctionCode is returned for an invalid business function code
	ErrBusinessFunctionCode = errors.New("is an invalid business function code")
	// ErrTransactionTypeCode is returned for an invalid transaction type code
	ErrTransactionTypeCode = errors.New("is an invalid transaction type code")

	// ErrLocalInstrumentNotPermitted is returned when LocalInstrument is included and BusinessFunctionCode is NOT CustomerTransferPlus
	ErrLocalInstrumentNotPermitted = errors.New("is only permitted for business function code CTP")
	// ErrLocalInstrumentCode is returned for an invalid local instrument code tag {3610}
	ErrLocalInstrumentCode = errors.New("is an invalid local instrument Code")
	// ErrPaymentNotificationIndicator is returned for an invalid payment notification indicator {3620}
	ErrPaymentNotificationIndicator = errors.New("is an invalid payment notification indicator")

	// Charges Tag {3700}

	// ErrChargeDetails is returned for an invalid charge details for charges
	ErrChargeDetails = errors.New("is an invalid charge detail")

	// Beneficiary {4000}

	// ErrIdentificationCode is returned for an invalid identification code
	ErrIdentificationCode = errors.New("is an invalid identification code")

	// ErrAdviceCode is returned for an invalid advice code
	ErrAdviceCode = errors.New("is an invalid advice code")

	// Related Remittance Information {8250}

	// ErrRemittanceLocationMethod is returned for an invalid remittance location method
	ErrRemittanceLocationMethod = errors.New("is an invalid remittance location method")

	// ErrAddressType is returned for an invalid address type
	ErrAddressType = errors.New("is an invalid address type")

	// ErrIdentificationType is returned for an invalid remittance Identification Typ
	ErrIdentificationType = errors.New("is an invalid remittance identification type")

	// ErrOrganizationIdentificationCode is returned for an invalid organization identification code
	ErrOrganizationIdentificationCode = errors.New("is an invalid organization identification code")

	// ErrPrivateIdentificationCode is returned for an invalid private identification code
	ErrPrivateIdentificationCode = errors.New("is an invalid private identification code")

	// ErrDocumentTypeCode is returned for an invalid document type code
	ErrDocumentTypeCode = errors.New("is an invalid document type code")

	// ErrCreditDebitIndicator is returned for an invalid credit or debit indicator
	ErrCreditDebitIndicator = errors.New("is an invalid credit or debit indicator")

	// ErrAdjustmentReasonCode is returned for an invalid adjustment reason code
	ErrAdjustmentReasonCode = errors.New("is an invalid adjustment reason code")

	// ErrPartyIdentifier is returned for an invalid party identifier
	ErrPartyIdentifier = errors.New("is an invalid party identifier")

	// ErrOptionFLine is returned for an invalid line for OriginatorOptionF
	ErrOptionFLine = errors.New("is an invalid line for originator optionF")

	// ErrOptionFName is returned for an invalid name for OriginatorOptionF
	ErrOptionFName = errors.New("is an invalid name for originator optionF")
)

// FieldError is returned for errors at a field level in a tag
type FieldError struct {
	FieldName string      // field name where error happened
	Value     interface{} // value that cause error
	Err       error       // context of the error.
	Msg       string      // deprecated
}

// Error message is constructed

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %v %s", e.FieldName, e.Value, e.Err)
}

// Unwrap implements the base.UnwrappableError interface for FieldError
func (e *FieldError) Unwrap() error {
	return e.Err
}

func fieldError(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*FieldError); ok {
		return err
	}
	fe := FieldError{
		FieldName: field,
		Err:       err,
	}
	// only the first value counts
	if len(values) > 0 {
		fe.Value = values[0]
	}
	return &fe
}

// ErrBusinessFunctionCodeProperty is the error given when the observed check digit does not match the calculated one
type ErrBusinessFunctionCodeProperty struct {
	Message              string
	Property             string
	PropertyValue        string
	BusinessFunctionCode string
}

// NewErrBusinessFunctionCodeProperty creates a new error of the ErrBusinessFunctionCodeProperty type
func NewErrBusinessFunctionCodeProperty(property, propertyValue, businessFunctionCode string) ErrBusinessFunctionCodeProperty {
	return ErrBusinessFunctionCodeProperty{
		Message:              fmt.Sprintf("%v: %v is not valid for %v", property, propertyValue, businessFunctionCode),
		Property:             property,
		PropertyValue:        propertyValue,
		BusinessFunctionCode: businessFunctionCode,
	}
}

func (e ErrBusinessFunctionCodeProperty) Error() string {
	return e.Message
}

// ErrInvalidPropertyForProperty is the error given when the observed check digit does not match the calculated one
type ErrInvalidPropertyForProperty struct {
	Message             string
	Property            string
	PropertyValue       string
	SecondProperty      string
	SecondPropertyValue string
}

// NewErrInvalidPropertyForProperty creates a new error of the ErrInvalidPropertyForProperty type
func NewErrInvalidPropertyForProperty(property, propertyValue, secondProperty, secondPropertyValue string) ErrInvalidPropertyForProperty {
	return ErrInvalidPropertyForProperty{
		Message:             fmt.Sprintf("%v: %v is not valid for %v: %v", property, propertyValue, secondProperty, secondPropertyValue),
		Property:            property,
		PropertyValue:       propertyValue,
		SecondProperty:      secondProperty,
		SecondPropertyValue: secondPropertyValue,
	}
}

func (e ErrInvalidPropertyForProperty) Error() string {
	return e.Message
}

// FieldWrongMinLengthErr is the error given when a Field is the wrong length
type FieldWrongMinLengthErr struct {
	Message     string
	FieldLength int
	Length      int
}

// NewFieldWrongMinLengthErr creates a new error of the FieldWrongLengthErr type
func NewFieldWrongMinLengthErr(FieldLength, length int) FieldWrongMinLengthErr {
	return FieldWrongMinLengthErr{
		Message:     fmt.Sprintf("must be %d characters or more and found %d", FieldLength, length),
		FieldLength: FieldLength,
		Length:      length,
	}
}

func (e FieldWrongMinLengthErr) Error() string {
	return e.Message
}
