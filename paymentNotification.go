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

// PaymentNotification is the PaymentNotification of the wire
type PaymentNotification struct {
	// tag
	tag string
	// PaymentNotificationIndicator
	// * `0 - 6` - Reserved for market practice conventions.
	// * `7 - 9` - Reserved for bilateral agreements between Fedwire senders and receivers.
	PaymentNotificationIndicator string `json:"paymentNotificationIndicator,omitempty"`
	// ContactNotificationElectronicAddress
	ContactNotificationElectronicAddress string `json:"contactNotificationElectronicAddress,omitempty"`
	// ContactName
	ContactName string `json:"contactName,omitempty"`
	// ContactPhoneNumber
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	// ContactMobileNumber
	ContactMobileNumber string `json:"contactMobileNumber,omitempty"`
	// FaxNumber
	ContactFaxNumber string `json:"faxNumber,omitempty"`
	// EndToEndIdentification
	EndToEndIdentification string `json:"endToEndIdentification,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPaymentNotification returns a new PaymentNotification
func NewPaymentNotification() *PaymentNotification {
	pn := &PaymentNotification{
		tag: TagPaymentNotification,
	}
	return pn
}

// Parse takes the input string and parses the PaymentNotification values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pn *PaymentNotification) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 13 || dataLen > 2341 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [13, 2341] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	pn.tag = record[:6]
	pn.PaymentNotificationIndicator = pn.parseStringField(record[6:7])

	optionalFields := strings.Split(record[7:], "*")
	if len(optionalFields) >= 1 {
		pn.ContactNotificationElectronicAddress = pn.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		pn.ContactName = pn.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		pn.ContactPhoneNumber = pn.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		pn.ContactMobileNumber = pn.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		pn.ContactFaxNumber = pn.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		pn.EndToEndIdentification = pn.parseStringField(optionalFields[5])
	}
	return nil
}

func (pn *PaymentNotification) UnmarshalJSON(data []byte) error {
	type Alias PaymentNotification
	aux := struct {
		*Alias
	}{
		(*Alias)(pn),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	pn.tag = TagPaymentNotification
	return nil
}

// String writes PaymentNotification
func (pn *PaymentNotification) String() string {
	var buf strings.Builder
	buf.Grow(2335)
	buf.WriteString(pn.tag)
	buf.WriteString(pn.PaymentNotificationIndicatorField())
	buf.WriteString(strings.TrimSpace(pn.ContactNotificationElectronicAddressField()) + "*")
	buf.WriteString(strings.TrimSpace(pn.ContactNameField()) + "*")
	buf.WriteString(strings.TrimSpace(pn.ContactPhoneNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(pn.ContactMobileNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(pn.ContactFaxNumberField()) + "*")
	buf.WriteString(strings.TrimSpace(pn.EndToEndIdentificationField()) + "*")
	return buf.String()
}

// Validate performs WIRE format rule checks on PaymentNotification and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pn *PaymentNotification) Validate() error {
	if pn.tag != TagPaymentNotification {
		return fieldError("tag", ErrValidTagForType, pn.tag)
	}
	if err := pn.isNumeric(pn.PaymentNotificationIndicator); err != nil {
		return fieldError("PaymentNotificationIndicator", err, pn.PaymentNotificationIndicator)
	}
	if err := pn.isAlphanumeric(pn.ContactNotificationElectronicAddress); err != nil {
		return fieldError("ContactNotificationElectronicAddress", err, pn.ContactNotificationElectronicAddress)
	}
	if err := pn.isAlphanumeric(pn.ContactName); err != nil {
		return fieldError("ContactName", err, pn.ContactName)
	}
	if err := pn.isAlphanumeric(pn.ContactPhoneNumber); err != nil {
		return fieldError("ContactPhoneNumber", err, pn.ContactPhoneNumber)
	}
	if err := pn.isAlphanumeric(pn.ContactMobileNumber); err != nil {
		return fieldError("ContactMobileNumber", err, pn.ContactMobileNumber)
	}
	if err := pn.isAlphanumeric(pn.ContactFaxNumber); err != nil {
		return fieldError("FaxNumber", err, pn.ContactFaxNumber)
	}
	if err := pn.isAlphanumeric(pn.EndToEndIdentification); err != nil {
		return fieldError("EndToEndIdentification", err, pn.EndToEndIdentification)
	}
	return nil
}

// PaymentNotificationIndicatorField gets a string of PaymentNotificationIndicator field
func (pn *PaymentNotification) PaymentNotificationIndicatorField() string {
	return pn.alphaField(pn.PaymentNotificationIndicator, 1)
}

// ContactNotificationElectronicAddressField gets a string of ContactNotificationElectronicAddress field
func (pn *PaymentNotification) ContactNotificationElectronicAddressField() string {
	return pn.alphaField(pn.ContactNotificationElectronicAddress, 2048)
}

// ContactNameField gets a string of ContactName field
func (pn *PaymentNotification) ContactNameField() string {
	return pn.alphaField(pn.ContactName, 140)
}

// ContactPhoneNumberField gets a string of ContactPhoneNumberField field
func (pn *PaymentNotification) ContactPhoneNumberField() string {
	return pn.alphaField(pn.ContactPhoneNumber, 35)
}

// ContactMobileNumberField gets a string of ContactMobileNumber field
func (pn *PaymentNotification) ContactMobileNumberField() string {
	return pn.alphaField(pn.ContactMobileNumber, 35)
}

// ContactFaxNumberField gets a string of FaxNumber field
func (pn *PaymentNotification) ContactFaxNumberField() string {
	return pn.alphaField(pn.ContactFaxNumber, 35)
}

// EndToEndIdentificationField gets a string of EndToEndIdentification field
func (pn *PaymentNotification) EndToEndIdentificationField() string {
	return pn.alphaField(pn.EndToEndIdentification, 35)
}
