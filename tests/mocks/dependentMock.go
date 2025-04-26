package mocks

import (
	"database/sql"
	"time"

	"github.com/brianmorais/go-user-importation/domain/enums/kinshipType"
	dependentModel "github.com/brianmorais/go-user-importation/domain/models/dependent"
)

func GetDependentsMock() dependentModel.Dependents {
	return dependentModel.Dependents{
		dependentModel.Dependent{
			DependentId:  sql.NullInt64{Int64: 1},
			FirstName:    sql.NullString{String: "Antonio"},
			LastName:     sql.NullString{String: "Carlos"},
			DateOfBirth:  sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Gender:       sql.NullString{String: "M"},
			Active:       sql.NullBool{Bool: true},
			CreatedDate:  sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser: sql.NullString{String: "Admin"},
			EmployeeId:   sql.NullInt64{Int64: 1},
			Kinship:      sql.NullInt32{Int32: kinshipType.Parent},
			Cpf:          sql.NullString{String: "123456789"},
		},
		dependentModel.Dependent{
			DependentId:  sql.NullInt64{Int64: 2},
			FirstName:    sql.NullString{String: "Marcos"},
			LastName:     sql.NullString{String: "Rubens"},
			DateOfBirth:  sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Gender:       sql.NullString{String: "M"},
			Active:       sql.NullBool{Bool: true},
			CreatedDate:  sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser: sql.NullString{String: "Admin"},
			EmployeeId:   sql.NullInt64{Int64: 2},
			Kinship:      sql.NullInt32{Int32: kinshipType.Parent},
			Cpf:          sql.NullString{String: "987456321"},
		},
		dependentModel.Dependent{
			DependentId:  sql.NullInt64{Int64: 3},
			FirstName:    sql.NullString{String: "Maria"},
			LastName:     sql.NullString{String: "Clara"},
			DateOfBirth:  sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Gender:       sql.NullString{String: "F"},
			Active:       sql.NullBool{Bool: true},
			CreatedDate:  sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser: sql.NullString{String: "Admin"},
			EmployeeId:   sql.NullInt64{Int64: 3},
			Kinship:      sql.NullInt32{Int32: kinshipType.Spouse},
			Cpf:          sql.NullString{String: "123987456"},
		},
	}
}

func GetDependentsViewMock() dependentModel.DependentsView {
	return dependentModel.DependentsView{
		dependentModel.DependentView{
			Re:                    sql.NullInt64{Int64: 1},
			IdDependent:           sql.NullInt64{Int64: 1},
			PrimeirioNome:         sql.NullString{String: "Antonio"},
			Sobrenome:             sql.NullString{String: "Carlos"},
			NomeCompleto:          sql.NullString{String: "Antonio Carlos"},
			CdiLigacaoPessoa:      sql.NullInt32{Int32: 1},
			DataNascimento:        sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Sexo:                  sql.NullString{String: "M"},
			CdiSituacaoDependente: sql.NullInt32{Int32: 1},
			Cpf:                   sql.NullString{String: "987456321"},
		},
		dependentModel.DependentView{
			Re:                    sql.NullInt64{Int64: 2},
			IdDependent:           sql.NullInt64{Int64: 2},
			PrimeirioNome:         sql.NullString{String: "Marcos"},
			Sobrenome:             sql.NullString{String: "Rubens"},
			NomeCompleto:          sql.NullString{String: "Marcos Rubens"},
			CdiLigacaoPessoa:      sql.NullInt32{Int32: 1},
			DataNascimento:        sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Sexo:                  sql.NullString{String: "M"},
			CdiSituacaoDependente: sql.NullInt32{Int32: 1},
			Cpf:                   sql.NullString{String: "123456789"},
		},
		dependentModel.DependentView{
			Re:                    sql.NullInt64{Int64: 3},
			IdDependent:           sql.NullInt64{Int64: 3},
			PrimeirioNome:         sql.NullString{String: "Maria"},
			Sobrenome:             sql.NullString{String: "Clara"},
			NomeCompleto:          sql.NullString{String: "Maria Clara"},
			CdiLigacaoPessoa:      sql.NullInt32{Int32: 1},
			DataNascimento:        sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
			Sexo:                  sql.NullString{String: "F"},
			CdiSituacaoDependente: sql.NullInt32{Int32: 1},
			Cpf:                   sql.NullString{String: "123987456"},
		},
	}
}

func GetDependent() dependentModel.Dependent {
	return dependentModel.Dependent{
		DependentId:  sql.NullInt64{Int64: 1},
		FirstName:    sql.NullString{String: "Antonio"},
		LastName:     sql.NullString{String: "Carlos"},
		DateOfBirth:  sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
		Gender:       sql.NullString{String: "M"},
		Active:       sql.NullBool{Bool: true},
		CreatedDate:  sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
		ModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
		ModifiedUser: sql.NullString{String: "Admin"},
		EmployeeId:   sql.NullInt64{Int64: 1},
		Kinship:      sql.NullInt32{Int32: kinshipType.Parent},
		Cpf:          sql.NullString{String: "123456789"},
	}
}

func GetDependentView() dependentModel.DependentView {
	return dependentModel.DependentView{
		Re:                    sql.NullInt64{Int64: 1},
		IdDependent:           sql.NullInt64{Int64: 1},
		PrimeirioNome:         sql.NullString{String: "Antonio", Valid: true},
		Sobrenome:             sql.NullString{String: "Carlos", Valid: true},
		NomeCompleto:          sql.NullString{String: "Antonio Carlos"},
		CdiLigacaoPessoa:      sql.NullInt32{Int32: 1},
		DataNascimento:        sql.NullTime{Time: time.Date(2000, 10, 1, 12, 0, 0, 0, time.UTC)},
		Sexo:                  sql.NullString{String: "M"},
		CdiSituacaoDependente: sql.NullInt32{Int32: 1},
		Cpf:                   sql.NullString{String: "123456789"},
	}
}
