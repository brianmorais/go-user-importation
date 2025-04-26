package benefitModel

import "encoding/json"

type Benefit struct {
	RoleId              int64 `json:"roleId"`
	MarketTypeId        int64 `json:"marketTypeId"`
	BenefitTypeId       int64 `json:"benefitTypeId"`
	PassengerTypeId     int64 `json:"passengerTypeId"`
	AvailableJourneyQty int64 `json:"availableJourneyQtd"`
}

type Benefits []Benefit

func DeserializeBenefit(jsonBenefits string) (Benefits, error) {
	var benefits Benefits

	err := json.Unmarshal([]byte(jsonBenefits), &benefits)

	if err != nil {
		return Benefits{}, err
	}

	return benefits, nil
}

func (benefits Benefits) GetBenefitsByRoleId(roleId int64) Benefits {
	var foundBenefits Benefits

	for i := range benefits {
		if benefits[i].RoleId == roleId {
			foundBenefits = append(foundBenefits, benefits[i])
		}
	}

	return foundBenefits
}

func (benefits Benefits) SerializeBenefit() (string, error) {
	res, err := json.Marshal(benefits)

	if err != nil {
		return "", err
	}

	return string(res), nil
}
