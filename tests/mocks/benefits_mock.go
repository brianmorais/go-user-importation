package mocks

import benefitModel "github.com/brianmorais/go-user-importation/domain/models/benefit"

func GetBenefits() benefitModel.Benefits {
	jsonBenefits := GetJsonBenefits()
	content, _ := benefitModel.DeserializeBenefit(jsonBenefits)

	return content
}

func GetJsonBenefits() string {
	return `[
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		}
	  ]`
}
