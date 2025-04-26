package interfaces

type ILogRepository interface {
	SaveLog(logMessage string, errLog error, registrationNumber string, importationNumber string, modifiedUser string) (int64, error)
}
