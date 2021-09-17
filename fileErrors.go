// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"errors"
	"fmt"
)

var (
	// ErrFileTooLong is the error given when a file exceeds the maximum possible length
	ErrFileTooLong = errors.New("file exceeds maximum possible number of lines")
)

// ErrInvalidTag is the error given when a tag is invalid
type ErrInvalidTag struct {
	Message string
	Type    string
}

// NewErrInvalidTag creates a new error of the ErrInvalidTag type
func NewErrInvalidTag(tag string) ErrInvalidTag {
	return ErrInvalidTag{
		Message: fmt.Sprintf("%s is an invalid tag", tag),
		Type:    tag,
	}
}

func (e ErrInvalidTag) Error() string {
	return e.Message
}
