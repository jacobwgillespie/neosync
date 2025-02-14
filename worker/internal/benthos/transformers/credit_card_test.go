package neosync_transformers

import (
	"strconv"
	"testing"

	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/stretchr/testify/assert"
)

func TestGenerateValidLuhnCreditCard(t *testing.T) {

	val, err := GenerateValidVLuhnCheckCreditCard()

	assert.NoError(t, err)
	assert.Len(t, strconv.FormatInt(val, 10), 16, "The output credit card should be 16 characteres long")
	assert.Equal(t, isValidLuhn(val), true)
}

func TestGenerateRandomLuhnCreditCard(t *testing.T) {

	val, err := GenerateCreditCard(false)

	assert.NoError(t, err)
	assert.Len(t, strconv.FormatInt(val, 10), 16, "The output credit card should be 16 characteres long")
	assert.Equal(t, isValidLuhn(val), false)
}

func TestGenerateCreditCardTransformer(t *testing.T) {
	mapping := `root = creditcardtransformer(true)`
	ex, err := bloblang.Parse(mapping)
	assert.NoError(t, err, "failed to parse the random int transformer")

	res, err := ex.Query(nil)
	assert.NoError(t, err)

	assert.Len(t, strconv.FormatInt(res.(int64), 10), 16, "The output credit card should be 16 characteres long")
	assert.Equal(t, isValidLuhn(res.(int64)), true)
}

func isValidLuhn(cc int64) bool {

	return (cc%10+checksum(cc/10))%10 == 0

}

func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur *= 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number /= 10
	}
	return luhn % 10
}
