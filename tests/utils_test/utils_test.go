package utils_test

import (
	"testing"

	"github.com/brianmorais/go-user-importation/domain/utils"
	"github.com/stretchr/testify/assert"
)

func TestMustCleanString(t *testing.T) {
	str := "    áã êúç   ñóòp  "
	expectedValue := "AA EUC NOOP"

	cleanString := utils.CleanString(str)

	assert.Equal(t, expectedValue, cleanString)
}
