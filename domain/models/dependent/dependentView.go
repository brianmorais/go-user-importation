package dependentModel

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/brianmorais/go-user-importation/domain/enums/kinshipType"
	"github.com/brianmorais/go-user-importation/domain/utils"
)

type DependentView struct {
	Re                    sql.NullInt64
	IdDependent           sql.NullInt64
	PrimeirioNome         sql.NullString
	Sobrenome             sql.NullString
	NomeCompleto          sql.NullString
	CdiLigacaoPessoa      sql.NullInt32
	DataNascimento        sql.NullTime
	Sexo                  sql.NullString
	CdiSituacaoDependente sql.NullInt32
	Cpf                   sql.NullString
}

type DependentsView []DependentView

func (dependent *DependentView) GetViewDataKindship() int32 {
	switch dependent.CdiLigacaoPessoa.Int32 {
	case 2, 1, 8:
		return kinshipType.Child
	case 3, 4:
		return kinshipType.Spouse
	case 5, 21:
		return kinshipType.Parent
	default:
		return kinshipType.Other
	}
}

func (dependents *DependentsView) FindDependentByRegistrationNumberAndName(registrationNumber int64, fullName string) DependentView {
	for i := range *dependents {
		name := fmt.Sprintf("%v %v", utils.CleanString((*dependents)[i].GetPrimeiroNome()), utils.CleanString((*dependents)[i].GetSobrenome()))
		if (*dependents)[i].Re.Int64 == registrationNumber && name == fullName {
			return (*dependents)[i]
		}
	}

	return DependentView{}
}

func (dependent *DependentView) NameIsValid() bool {
	return dependent.PrimeirioNome.Valid && dependent.Sobrenome.Valid && dependent.GetPrimeiroNome() != "" && dependent.GetSobrenome() != ""
}

func (dependent DependentView) RegistrationNumberToString() string {
	return strconv.FormatInt(dependent.Re.Int64, 10)
}

func (dependent *DependentView) GetPrimeiroNome() string {
	return dependent.PrimeirioNome.String
}

func (dependent *DependentView) GetSobrenome() string {
	return dependent.Sobrenome.String
}

func (dependent *DependentView) GetNomeCompleto() string {
	return dependent.NomeCompleto.String
}

func (dependent *DependentView) GetCdiSituacaoDependente() int32 {
	return dependent.CdiSituacaoDependente.Int32
}

func (dependent *DependentView) GetDataNascimento() time.Time {
	return dependent.DataNascimento.Time
}

func (dependent *DependentView) GetSexo() string {
	return dependent.Sexo.String
}

func (dependent *DependentView) GetCpf() string {
	return dependent.Cpf.String
}
