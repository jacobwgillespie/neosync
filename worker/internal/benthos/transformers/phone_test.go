package neosync_transformers

import (
	"strings"
	"testing"

	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePhoneNumberPreserveLengthTrue(t *testing.T) {

	pn := "1838492832"
	expectedLength := 10

	res, err := GeneratePhoneNumberPreserveLengthNoHyphensNotE164(pn)

	assert.NoError(t, err)
	assert.Equal(t, len(res), expectedLength, "The length of the output phone number should be the same as the input phone number")

}

func TestGeneratePhoneNumberPreserveLengthTrueHyphens(t *testing.T) {

	// we strip the hyphens when we Generate the phone number and the include hyphens param is set to false so the return val will not include hyphens
	pn := "183-849-2838"
	expectedLength := 10

	res, err := GeneratePhoneNumber(pn, true, false, false)

	assert.NoError(t, err)
	assert.Equal(t, len(res), expectedLength, "The length of the output phone number should be the same as the input phone number")

}

func TestGeneratePhoneNumberPreserveLengthFalseHyphens(t *testing.T) {

	pn := "183-849-2831"

	res, err := GeneratePhoneNumber(pn, false, true, true)

	assert.NoError(t, err)
	assert.False(t, strings.Contains(res, "-"), "The output int phone number should not contain hyphens and may not be the same length as the input")

}

func TestGeneratePhoneNumberPreserveLengthFalseNoHyphens(t *testing.T) {

	res, err := GenerateRandomPhoneNumberWithNoHyphens()

	assert.NoError(t, err)
	assert.False(t, strings.Contains(res, "-"), "The output int phone number should not contain hyphens and may not be the same length as the input")
	assert.Equal(t, len(res), 10)
}

func TestGeneratePhoneNumberPreserveLengthFalseIncludeHyphensTrue(t *testing.T) {

	expectedLength := 12

	res, err := GenerateRandomPhoneNumberWithHyphens()

	assert.NoError(t, err)
	assert.Equal(t, len(res), expectedLength, "The length of the output phone number should be the same as the input phone number")

}

func TestGeneratePhoneNumberPreserveLengthTrueIncludeHyphensTrueError(t *testing.T) {

	pn := "183-849-2839"
	_, err := GeneratePhoneNumber(pn, true, false, true)

	assert.Error(t, err, "The include hyphens param can only be used by itself, all other params must be false")

}

func TestGeneratePhoneNumberE164Format(t *testing.T) {

	pn := "+2393573894"
	expectedLength := 11

	res, err := GenerateE164FormatPhoneNumber()

	assert.NoError(t, err)
	assert.Equal(t, ValidateE164(res), ValidateE164("+2393573894"))
	assert.Equal(t, len(pn), expectedLength, "The length of the output phone number should be the same as the input phone number")

}

func TestGeneratePhoneNumberE164FormatPreserveLength(t *testing.T) {

	pn := "+2393573894"
	expectedLength := 11

	res, err := GenerateE164FormatPhoneNumberPreserveLength(pn)

	assert.NoError(t, err)
	assert.Equal(t, ValidateE164(res), ValidateE164("+2393573894"))
	assert.Equal(t, len(pn), expectedLength, "The length of the output phone number should be the same as the input phone number")

}

func TestPhoneNumberTransformer(t *testing.T) {
	mapping := `root = this.phonetransformer(true, false, false)`
	ex, err := bloblang.Parse(mapping)
	assert.NoError(t, err, "failed to parse the phone transformer")

	testVal := "6183849282"

	res, err := ex.Query(testVal)
	assert.NoError(t, err)

	assert.Len(t, res.(string), len(testVal), "Generated phone number must be the same length as the input phone number")
}

func TestValidateE164True(t *testing.T) {

	val := "+6272636472"

	res := ValidateE164(val)

	assert.Equal(t, res, true, "The e164 number should have a plus sign at the 0th index and be 10 < x < 15 characters long.")
}

func TestValidateE164FalseTooLong(t *testing.T) {

	val := "627263647278439"

	res := ValidateE164(val)

	assert.Equal(t, res, false, "The e164 number should  be x < 15 characters long.")
}

func TestValidateE164FalseNoPlusSign(t *testing.T) {

	val := "6272636472784"

	res := ValidateE164(val)

	assert.Equal(t, res, false, "The e164 number should have a plus sign at the beginning.")
}

func TestValidateE164FalseTooshort(t *testing.T) {

	val := "627263"

	res := ValidateE164(val)

	assert.Equal(t, res, false, "The e164 number should  be 10 < x")
}
