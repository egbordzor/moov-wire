// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
)

// UnstructuredAddenda is the unstructured addenda information
type UnstructuredAddenda struct {
	// tag
	tag string
	// Addenda
	Addenda string `json:"addenda,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewUnstructuredAddenda returns a new UnstructuredAddenda
func NewUnstructuredAddenda() *UnstructuredAddenda {
	ua := &UnstructuredAddenda{
		tag: TagUnstructuredAddenda,
	}
	return ua
}

// Parse takes the input string and parses the UnstructuredAddenda values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ua *UnstructuredAddenda) Parse(record string) error {
	ua.tag = record[:6]
	ua.Addenda = ua.parseStringField(record[6:])
	return nil
}

func (ua *UnstructuredAddenda) UnmarshalJSON(data []byte) error {
	type Alias UnstructuredAddenda
	aux := struct {
		*Alias
	}{
		(*Alias)(ua),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ua.tag = TagUnstructuredAddenda
	return nil
}

// String writes UnstructuredAddenda
func (ua *UnstructuredAddenda) String() string {
	var buf strings.Builder
	buf.Grow(len(ua.tag) + len(ua.Addenda))
	buf.WriteString(ua.tag)
	buf.WriteString(ua.AddendaField())
	return buf.String()
}

// Validate performs WIRE format rule checks on UnstructuredAddenda and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// AddendaLength must be numeric, padded with leading zeros if less than four characters and must equal
//  length of content in Addenda Information (e.g., if content of Addenda Information is 987 characters,
//  Addenda Length must be 0987).
func (ua *UnstructuredAddenda) Validate() error {
	if err := ua.fieldInclusion(); err != nil {
		return err
	}
	if ua.tag != TagUnstructuredAddenda {
		return fieldError("tag", ErrValidTagForType, ua.tag)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ua *UnstructuredAddenda) fieldInclusion() error {
	return nil
}

// AddendaField gets a string of the Addenda field
func (ua *UnstructuredAddenda) AddendaField() string {
	return ua.Addenda
}
