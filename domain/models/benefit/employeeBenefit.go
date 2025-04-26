package benefitModel

import (
	"database/sql"
	"time"
)

type EmployeeBenefit struct {
	EmployeeId   sql.NullInt64
	CycleId      sql.NullInt64
	RoleId       sql.NullInt32
	Benefit      sql.NullString
	CreatedDate  sql.NullTime
	ModifiedDate sql.NullTime
	ModifiedUser sql.NullString
}

func (benefit *EmployeeBenefit) SetEmployeeId(employeeId int64) {
	benefit.EmployeeId = sql.NullInt64{Int64: employeeId, Valid: true}
}

func (benefit *EmployeeBenefit) SetCycleId(cycleId int64) {
	benefit.CycleId = sql.NullInt64{Int64: cycleId, Valid: true}
}

func (benefit *EmployeeBenefit) SetRoleId(roleId int32) {
	benefit.RoleId = sql.NullInt32{Int32: roleId, Valid: true}
}

func (benefit *EmployeeBenefit) SetBenefit(benefitSetting string) {
	benefit.Benefit = sql.NullString{String: benefitSetting, Valid: true}
}

func (benefit *EmployeeBenefit) SetCreatedDate(date time.Time) {
	benefit.CreatedDate = sql.NullTime{Time: date, Valid: true}
}

func (benefit *EmployeeBenefit) SetModifiedDate(date time.Time) {
	benefit.ModifiedDate = sql.NullTime{Time: date, Valid: true}
}

func (benefit *EmployeeBenefit) SetModifiedUser(user string) {
	benefit.ModifiedUser = sql.NullString{String: user, Valid: true}
}

func (benefit *EmployeeBenefit) IsValid() bool {
	return benefit.CycleId.Valid && (benefit != &EmployeeBenefit{})
}
