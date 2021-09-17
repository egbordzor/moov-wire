package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//  CurrencyInstructedAmount creates a CurrencyInstructedAmount
func mockCurrencyInstructedAmount() *CurrencyInstructedAmount {
	cia := NewCurrencyInstructedAmount()
	cia.SwiftFieldTag = "Swift Field Tag"
	cia.CurrencyCode = "USD"
	cia.Amount = "1500,49"
	return cia
}

// TestMockCurrencyInstructedAmount validates mockCurrencyInstructedAmount
func TestMockCurrencyInstructedAmount(t *testing.T) {
	cia := mockCurrencyInstructedAmount()

	require.NoError(t, cia.Validate(), "mockCurrencyInstructedAmount does not validate and will break other tests")
}

// TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric validates CurrencyInstructedAmount SwiftFieldTag is alphanumeric
func TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.SwiftFieldTag = "®"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, cia.SwiftFieldTag).Error())
}

// TestCurrencyInstructedAmountValid validates CurrencyInstructedAmount Amount is valid
func TestCurrencyInstructedAmountValid(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.Amount = "1-0"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, cia.Amount).Error())
}

// TestParseCurrencyInstructedAmountReaderParseError parses a wrong CurrencyInstructedAmount reader parse error
func TestParseCurrencyInstructedAmountReaderParseError(t *testing.T) {
	var line = "{7033}Swift*USD00000Z001500,49*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCurrencyInstructedAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "00000Z001500,49")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "00000Z001500,49")).Error())
}

// TestCurrencyInstructedAmountTagError validates a CurrencyInstructedAmount tag
func TestCurrencyInstructedAmountTagError(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.tag = "{9999}"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, cia.tag).Error())
}
