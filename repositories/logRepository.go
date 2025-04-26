package repositories

import (
	"time"
)

type LogRepository struct{}

func (LogRepository) SaveLog(logMessage string, errLog error, registrationNumber string, importationNumber string, modifiedUser string) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		INSERT INTO ImportationLog (
			ExceptionMessage,
			DateOfImportation,
			RegistrationNumber,
			ImportationNumber,
			UserName,
			ManualImportation,
			DetailsMessage
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

	res, err := conn.Exec(
		queryText,
		logMessage,
		time.Now(),
		registrationNumber,
		importationNumber,
		modifiedUser,
		false,
		errLog.Error(),
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
