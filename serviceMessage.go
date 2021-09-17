// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// ServiceMessage is the ServiceMessage of the wire
type ServiceMessage struct {
	// tag
	tag string
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`
	// LineFive
	LineFive string `json:"lineFive,omitempty"`
	// LineSix
	LineSix string `json:"lineSix,omitempty"`
	// LineSeven
	LineSeven string `json:"lineSeven,omitempty"`
	// LineEight
	LineEight string `json:"lineEight,omitempty"`
	// LineNine
	LineNine string `json:"lineNine,omitempty"`
	// LineTen
	LineTen string `json:"lineTen,omitempty"`
	// LineEleven
	LineEleven string `json:"lineEleven,omitempty"`
	// LineTwelve
	LineTwelve string `json:"lineTwelve,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewServiceMessage returns a new ServiceMessage
func NewServiceMessage() *ServiceMessage {
	sm := &ServiceMessage{
		tag: TagServiceMessage,
	}
	return sm
}

// Parse takes the input string and parses the ServiceMessage values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sm *ServiceMessage) Parse(record string) error {
	sm.tag = record[:6]

	optionalFields := strings.Split(record[6:], "*")
	if len(optionalFields) >= 1 {
		sm.LineOne = sm.parseStringField(optionalFields[0])
	}
	if len(optionalFields) >= 2 {
		sm.LineTwo = sm.parseStringField(optionalFields[1])
	}
	if len(optionalFields) >= 3 {
		sm.LineThree = sm.parseStringField(optionalFields[2])
	}
	if len(optionalFields) >= 4 {
		sm.LineFour = sm.parseStringField(optionalFields[3])
	}
	if len(optionalFields) >= 5 {
		sm.LineFive = sm.parseStringField(optionalFields[4])
	}
	if len(optionalFields) >= 6 {
		sm.LineSix = sm.parseStringField(optionalFields[5])
	}
	if len(optionalFields) >= 7 {
		sm.LineSeven = sm.parseStringField(optionalFields[6])
	}
	if len(optionalFields) >= 8 {
		sm.LineEight = sm.parseStringField(optionalFields[7])
	}
	if len(optionalFields) >= 9 {
		sm.LineNine = sm.parseStringField(optionalFields[8])
	}
	if len(optionalFields) >= 10 {
		sm.LineTen = sm.parseStringField(optionalFields[9])
	}
	if len(optionalFields) >= 11 {
		sm.LineEleven = sm.parseStringField(optionalFields[10])
	}
	if len(optionalFields) >= 12 {
		sm.LineTwelve = sm.parseStringField(optionalFields[11])
	}
	return nil
}

func (sm *ServiceMessage) UnmarshalJSON(data []byte) error {
	type Alias ServiceMessage
	aux := struct {
		*Alias
	}{
		(*Alias)(sm),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	sm.tag = TagServiceMessage
	return nil
}

// String writes ServiceMessage
func (sm *ServiceMessage) String() string {
	var buf strings.Builder
	buf.Grow(426)
	buf.WriteString(sm.tag)
	buf.WriteString(strings.TrimSpace(sm.LineOneField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineTwoField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineThreeField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineFourField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineFiveField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineSixField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineSevenField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineEightField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineNineField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineTenField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineElevenField()) + "*")
	buf.WriteString(strings.TrimSpace(sm.LineTwelveField()) + "*")
	return sm.cleanupDelimiters(buf.String())
}

// Validate performs WIRE format rule checks on ServiceMessage and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sm *ServiceMessage) Validate() error {
	if err := sm.fieldInclusion(); err != nil {
		return err
	}
	if sm.tag != TagServiceMessage {
		return fieldError("tag", ErrValidTagForType, sm.tag)
	}
	if err := sm.isAlphanumeric(sm.LineOne); err != nil {
		return fieldError("LineOne", err, sm.LineOne)
	}
	if err := sm.isAlphanumeric(sm.LineTwo); err != nil {
		return fieldError("LineTwo", err, sm.LineTwo)
	}
	if err := sm.isAlphanumeric(sm.LineThree); err != nil {
		return fieldError("LineThree", err, sm.LineThree)
	}
	if err := sm.isAlphanumeric(sm.LineFour); err != nil {
		return fieldError("LineFour", err, sm.LineFour)
	}
	if err := sm.isAlphanumeric(sm.LineFive); err != nil {
		return fieldError("LineFive", err, sm.LineFive)
	}
	if err := sm.isAlphanumeric(sm.LineSix); err != nil {
		return fieldError("LineSix", err, sm.LineSix)
	}
	if err := sm.isAlphanumeric(sm.LineSeven); err != nil {
		return fieldError("LineSeven", err, sm.LineSeven)
	}
	if err := sm.isAlphanumeric(sm.LineEight); err != nil {
		return fieldError("LineEight", err, sm.LineEight)
	}
	if err := sm.isAlphanumeric(sm.LineNine); err != nil {
		return fieldError("LineNine", err, sm.LineNine)
	}
	if err := sm.isAlphanumeric(sm.LineTen); err != nil {
		return fieldError("LineTen", err, sm.LineTen)
	}
	if err := sm.isAlphanumeric(sm.LineEleven); err != nil {
		return fieldError("LineEleven", err, sm.LineEleven)
	}
	if err := sm.isAlphanumeric(sm.LineTwelve); err != nil {
		return fieldError("LineTwelve", err, sm.LineTwelve)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sm *ServiceMessage) fieldInclusion() error {
	// If ServiceMessage is defined, LineOne is required
	if sm.LineOne == "" {
		return fieldError("LineOne", ErrFieldRequired)
	}

	return nil
}

// LineOneField gets a string of the LineOne field
func (sm *ServiceMessage) LineOneField() string {
	return sm.alphaField(sm.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (sm *ServiceMessage) LineTwoField() string {
	return sm.alphaField(sm.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (sm *ServiceMessage) LineThreeField() string {
	return sm.alphaField(sm.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (sm *ServiceMessage) LineFourField() string {
	return sm.alphaField(sm.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (sm *ServiceMessage) LineFiveField() string {
	return sm.alphaField(sm.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (sm *ServiceMessage) LineSixField() string {
	return sm.alphaField(sm.LineSix, 35)
}

// LineSevenField gets a string of the LineSeven field
func (sm *ServiceMessage) LineSevenField() string {
	return sm.alphaField(sm.LineSeven, 35)
}

// LineEightField gets a string of the LineEight field
func (sm *ServiceMessage) LineEightField() string {
	return sm.alphaField(sm.LineEight, 35)
}

// LineNineField gets a string of the LineNine field
func (sm *ServiceMessage) LineNineField() string {
	return sm.alphaField(sm.LineNine, 35)
}

// LineTenField gets a string of the LineTen field
func (sm *ServiceMessage) LineTenField() string {
	return sm.alphaField(sm.LineTen, 35)
}

// LineElevenField gets a string of the LineEleven field
func (sm *ServiceMessage) LineElevenField() string {
	return sm.alphaField(sm.LineEleven, 35)
}

// LineTwelveField gets a string of the LineTwelve field
func (sm *ServiceMessage) LineTwelveField() string {
	return sm.alphaField(sm.LineTwelve, 35)
}
