package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockErrorWire creates a ErrorWire
func mockErrorWire() *ErrorWire {
	ew := NewErrorWire()
	ew.ErrorCategory = "E"
	ew.ErrorCode = "XYZ"
	ew.ErrorDescription = "Data Error"
	return ew
}

// TestMockErrorWire validates mockErrorWire
func TestMockErrorWire(t *testing.T) {
	ew := mockErrorWire()

	require.NoError(t, ew.Validate(), "mockErrorWire does not validate and will break other tests")
}

// TestParseErrorWire parses a known ErrorWire  record string
func TestParseErrorWire(t *testing.T) {
	var line = "{1130}1XYZINVLD CYCLE DT/MISSING/INVLD {1520}*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseErrorWire())
	record := r.currentFEDWireMessage.ErrorWire

	assert.Equal(t, "1", record.ErrorCategory)
	assert.Equal(t, "XYZ", record.ErrorCode)
	assert.Equal(t, "INVLD CYCLE DT/MISSING/INVLD {1520}", record.ErrorDescription)
	assert.Equal(t, line, record.String())
}

func TestParseErrorWireEmptyDescription(t *testing.T) {
	var line = "{1130}1XYZ*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseErrorWire())
	record := r.currentFEDWireMessage.ErrorWire

	assert.Equal(t, "1", record.ErrorCategory)
	assert.Equal(t, "XYZ", record.ErrorCode)
	assert.Equal(t, "", record.ErrorDescription)
	assert.Equal(t, line, record.String())
}
