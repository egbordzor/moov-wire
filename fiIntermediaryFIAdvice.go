// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// FIIntermediaryFIAdvice is the financial institution intermediary financial institution
type FIIntermediaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFIAdvice returns a new FIIntermediaryFIAdvice
func NewFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	fiifia := &FIIntermediaryFIAdvice{
		tag: TagFIIntermediaryFIAdvice,
	}
	return fiifia
}

// Parse takes the input string and parses the FIIntermediaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifia *FIIntermediaryFIAdvice) Parse(record string) error {
	fiifia.tag = record[:6]
	fiifia.Advice.AdviceCode = fiifia.parseStringField(record[6:9])

	optionalFields := strings.Split(record[9:], "*")
	fiifia.Advice.LineOne = fiifia.parseStringField(optionalFields[0])
	if len(optionalFields) >= 2 {
		fiifia.Advice.LineTwo = fiifia.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		fiifia.Advice.LineThree = fiifia.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		fiifia.Advice.LineFour = fiifia.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		fiifia.Advice.LineFive = fiifia.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		fiifia.Advice.LineSix = fiifia.parseStringField(optionalFields[5])
	}
	return nil
}

func (fiifia *FIIntermediaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIIntermediaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fiifia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fiifia.tag = TagFIIntermediaryFIAdvice
	return nil
}

// String writes FIIntermediaryFIAdvice
func (fiifia *FIIntermediaryFIAdvice) String() string {
	var buf strings.Builder
	buf.Grow(200)
	buf.WriteString(fiifia.tag)
	buf.WriteString(fiifia.AdviceCodeField())
	buf.WriteString(strings.TrimSpace(fiifia.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifia.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifia.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifia.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifia.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(fiifia.LineSixField()) + "*")
	return fiifia.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on FIIntermediaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifia *FIIntermediaryFIAdvice) Validate() error {
	if fiifia.tag != TagFIIntermediaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fiifia.tag)
	}
	if err := fiifia.isAdviceCode(fiifia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fiifia.Advice.AdviceCode)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fiifia.Advice.LineOne)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifia.Advice.LineTwo)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fiifia.Advice.LineThree)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fiifia.Advice.LineFour)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fiifia.Advice.LineFive)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fiifia.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fiifia *FIIntermediaryFIAdvice) AdviceCodeField() string {
	return fiifia.alphaField(fiifia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fiifia *FIIntermediaryFIAdvice) LineOneField() string {
	return fiifia.alphaField(fiifia.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (fiifia *FIIntermediaryFIAdvice) LineTwoField() string {
	return fiifia.alphaField(fiifia.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fiifia *FIIntermediaryFIAdvice) LineThreeField() string {
	return fiifia.alphaField(fiifia.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fiifia *FIIntermediaryFIAdvice) LineFourField() string {
	return fiifia.alphaField(fiifia.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fiifia *FIIntermediaryFIAdvice) LineFiveField() string {
	return fiifia.alphaField(fiifia.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fiifia *FIIntermediaryFIAdvice) LineSixField() string {
	return fiifia.alphaField(fiifia.Advice.LineSix, 33)
}
