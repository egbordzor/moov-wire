package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// UnstructuredAddenda creates a UnstructuredAddenda
func mockUnstructuredAddenda() *UnstructuredAddenda {
	ua := NewUnstructuredAddenda()
	ua.Addenda = "Unstructured Addenda<<>>"
	return ua
}

// TestMockUnstructuredAddenda validates mockUnstructuredAddenda
func TestMockUnstructuredAddenda(t *testing.T) {
	ua := mockUnstructuredAddenda()

	require.NoError(t, ua.Validate(), "mockUnstructuredAddenda does not validate and will break other tests")
}

// TestUnstructuredAddendaTagError validates a UnstructuredAddenda tag
func TestUnstructuredAddendaTagError(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.tag = "{9999}"

	require.EqualError(t, ua.Validate(), fieldError("tag", ErrValidTagForType, ua.tag).Error())
}

func TestUnstructuredAddendaISO20022(t *testing.T) {
	raw := `{8200}<!-- <Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"><FIToFICstmrCdtTrf><GrpHdr><MsgId>BCA2303235139631</MsgId><CreDtTm>2023-03-23T20:49:46+00:00</CreDtTm><NbOfTxs>1</NbOfTxs><SttlmInf><SttlmMtd>INDA</SttlmMtd><SttlmAcct><Id><Othr><Id>04464371</Id></Othr></Id></SttlmAcct></SttlmInf></GrpHdr><CdtTrfTxInf><PmtId><InstrId>WW23032358229581</InstrId><EndToEndId>WW23032358229581</EndToEndId><UETR>961076b0-d9fd-421d-8c89-9433e8e78cf9</UETR></PmtId><PmtTpInf><SvcLvl><Cd>G001</Cd></SvcLvl></PmtTpInf><IntrBkSttlmAmt Ccy="USD">10000</IntrBkSttlmAmt><IntrBkSttlmDt>2023-03-23</IntrBkSttlmDt><InstdAmt Ccy="USD">10000</InstdAmt><ChrgBr>CRED</ChrgBr><ChrgsInf><Amt Ccy="USD">0</Amt><Agt><FinInstnId><BICFI>BOFMCAM2</BICFI></FinInstnId></Agt></ChrgsInf><InstgAgt><FinInstnId><BICFI>BOFMCAM2</BICFI></FinInstnId></InstgAgt><InstdAgt><FinInstnId><BICFI>BKTRUS33XXX</BICFI></FinInstnId></InstdAgt><Dbtr><Nm>10X DIGITAL SOLUTIONS INC.</Nm><PstlAdr><AdrLine>1055 W GEORGIA ST SUITE 1750</AdrLine><AdrLine>VANCOUVER,BC,V6E 3P3 CA</AdrLine></PstlAdr></Dbtr><DbtrAcct><Id><Othr><Id>00044574064</Id></Othr></Id><Ccy>USD</Ccy></DbtrAcct><DbtrAgt><FinInstnId><BICFI>BOFMCAM2XXX</BICFI></FinInstnId></DbtrAgt><CdtrAgt><FinInstnId><ClrSysMmbId><ClrSysId><Cd>USABA</Cd></ClrSysId><MmbId>121145349</MmbId></ClrSysMmbId><Nm>Column National Association</Nm><PstlAdr><PstCd>94129</PstCd><TwnNm>San Francisco</TwnNm><Ctry>US</Ctry></PstlAdr></FinInstnId></CdtrAgt><Cdtr><Nm>LIONS DIGITAL SOLUTIONS USA INC.</Nm><PstlAdr><AdrLine>2055 LIMESTONE RD STE 200C</AdrLine><AdrLine>WILMINGTON DE 19808 US</AdrLine></PstlAdr></Cdtr><CdtrAcct><Id><Othr><Id>339414045812319</Id></Othr></Id></CdtrAcct><RmtInf><Ustrd>CREDIT CARD REIMBURSEMENT</Ustrd></RmtInf></CdtTrfTxInf></FIToFICstmrCdtTrf></Document> --><<>>`
	ua := &UnstructuredAddenda{}
	require.NoError(t, ua.Parse(raw))
	require.NoError(t, ua.Validate())
	assert.Equal(t, ua.Addenda, raw[6:])
}
