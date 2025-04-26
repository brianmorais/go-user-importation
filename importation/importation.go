package importation

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golobby/container/v3"
	"github.com/brianmorais/go-user-importation/configuration"
	"github.com/brianmorais/go-user-importation/domain/interfaces"
	"github.com/brianmorais/go-user-importation/ioc"
)

var settings *configuration.Settings
var dependentRepository interfaces.IDependentRepository
var logRepository interfaces.ILogRepository
var employeeRepository interfaces.IEmployeeRepository
var userRepository interfaces.IUserRepository
var benefitRepository interfaces.IBenefitRepository
var importationNumber *string

func Run() {
	ioc.SetDependencies()
	setRepositories()
	generateImportationNumber()
	saveMessageLog(fmt.Sprintf("Início da importação - %v", *importationNumber))
	readEmployeeViewData()
	readDependentViewData()
	saveMessageLog(fmt.Sprintf("Fim da importação - %v", *importationNumber))
}

func generateImportationNumber() {
	layout := "060102150405"
	formated := time.Now().Format(layout)
	importationNumber = &formated
}

func setRepositories() {
	var err error

	if err = container.Resolve(&settings); err != nil {
		log.Panic(err)
	}

	if err = container.Resolve(&employeeRepository); err != nil {
		log.Panic(err)
	}

	if err = container.Resolve(&dependentRepository); err != nil {
		log.Panic(err)
	}

	if err = container.Resolve(&benefitRepository); err != nil {
		log.Panic(err)
	}

	if err = container.Resolve(&userRepository); err != nil {
		log.Panic(err)
	}

	if err = container.Resolve(&logRepository); err != nil {
		log.Panic(err)
	}
}

func saveMessageLog(message string) {
	saveErrorLog(message, errors.New(""), "")
}

func saveErrorLog(message string, err error, registrationNumber string) {
	logRepository.SaveLog(message, err, registrationNumber, *importationNumber, settings.ModifiedUser)
	if err != nil && err.Error() != "" {
		log.Printf("%v: %v", message, err.Error())
	} else {
		log.Println(message)
	}
}
