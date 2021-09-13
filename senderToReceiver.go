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

// SenderToReceiver is the remittance information
type SenderToReceiver struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderToReceiver returns a new SenderToReceiver
func NewSenderToReceiver() *SenderToReceiver {
	str := &SenderToReceiver{
		tag: TagSenderToReceiver,
	}
	return str
}

// Parse takes the input string and parses the SenderToReceiver values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (str *SenderToReceiver) Parse(record string) error {
	dataLen := utf8.RuneCountInString(record)
	if dataLen < 8 || dataLen > 228 {
		return TagWrongLengthErr{
			Message: fmt.Sprintf("must be [8, 228] characters and found %d", dataLen),
			Length:  dataLen,
		}
	}
	str.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		str.CoverPayment.SwiftFieldTag = str.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		str.CoverPayment.SwiftLineOne = str.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		str.CoverPayment.SwiftLineTwo = str.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		str.CoverPayment.SwiftLineThree = str.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		str.CoverPayment.SwiftLineFour = str.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		str.CoverPayment.SwiftLineFive = str.parseStringField(optionalFields[5])
	}
	if len(optionalFields) >= 7 {
		str.CoverPayment.SwiftLineSix = str.parseStringField(optionalFields[6])
	}
	return nil
}

func (str *SenderToReceiver) UnmarshalJSON(data []byte) error {
	type Alias SenderToReceiver
	aux := struct {
		*Alias
	}{
		(*Alias)(str),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	str.tag = TagSenderToReceiver
	return nil
}

// String writes SenderToReceiver
func (str *SenderToReceiver) String() string {
	var buf strings.Builder
	buf.Grow(221)
	buf.WriteString(str.tag)
	buf.WriteString(strings.TrimSpace(str.SwiftFieldTagField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(str.SwiftLineSixField()) + "*")
	return str.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on SenderToReceiver and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (str *SenderToReceiver) Validate() error {
	if str.tag != TagSenderToReceiver {
		return fieldError("tag", ErrValidTagForType, str.tag)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, str.CoverPayment.SwiftFieldTag)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, str.CoverPayment.SwiftLineOne)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, str.CoverPayment.SwiftLineTwo)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, str.CoverPayment.SwiftLineThree)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, str.CoverPayment.SwiftLineFour)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, str.CoverPayment.SwiftLineFive)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineSix); err != nil {
		return fieldError("SwiftLineSix", err, str.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (str *SenderToReceiver) SwiftFieldTagField() string {
	return str.alphaField(str.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (str *SenderToReceiver) SwiftLineOneField() string {
	return str.alphaField(str.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (str *SenderToReceiver) SwiftLineTwoField() string {
	return str.alphaField(str.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (str *SenderToReceiver) SwiftLineThreeField() string {
	return str.alphaField(str.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (str *SenderToReceiver) SwiftLineFourField() string {
	return str.alphaField(str.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (str *SenderToReceiver) SwiftLineFiveField() string {
	return str.alphaField(str.CoverPayment.SwiftLineFive, 35)
}

// SwiftLineSixField gets a string of the SwiftLineSix field
func (str *SenderToReceiver) SwiftLineSixField() string {
	return str.alphaField(str.CoverPayment.SwiftLineSix, 35)
}
