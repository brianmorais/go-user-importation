package models_test

import (
	"testing"

	benefitModel "github.com/brianmorais/go-user-importation/domain/models/benefit"
	"github.com/brianmorais/go-user-importation/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMustDeserializeBenefit(t *testing.T) {
	jsonContent := mocks.GetJsonBenefits()

	_, err := benefitModel.DeserializeBenefit(jsonContent)

	assert.Nilf(t, err, "Erro ao descerializar json de benefícios")
}

func TestMustGetBenefitsByRoleId(t *testing.T) {
	benefits := mocks.GetBenefits()
	roleId := int64(1)

	res := benefits.GetBenefitsByRoleId(roleId)

	assert.GreaterOrEqualf(t, len(res), 0, "Array de benefícios vazio")
}

func TestMustSerializeBenefit(t *testing.T) {
	benefits := mocks.GetBenefits()

	_, err := benefits.SerializeBenefit()

	assert.Nilf(t, err, "Erro ao serializar array de benefícios")
}
