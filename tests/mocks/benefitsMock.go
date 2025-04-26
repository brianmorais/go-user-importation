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
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 1,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": 4
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": 4
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 2,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 3,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 4,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 5,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2,
		  "availableJourneyQty": -1
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3,
		  "availableJourneyQty": 8
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 6,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 1
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 2
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 1,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 2,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 1,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 2,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		},
		{
		  "roleId": 7,
		  "marketTypeId": 3,
		  "benefitTypeId": 3,
		  "passengerTypeId": 3
		}
	  ]`
}
