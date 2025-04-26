package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Databases    Databases  `json:"databases"`
	GoRoutines   GoRoutines `json:"goRoutines"`
	ModifiedUser string     `json:"modifiedUser"`
	EmailDomain  string     `json:"emailDomain"`
}

type Database struct {
	Server       string `json:"server"`
	Port         int    `json:"port"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
	User         string `json:"user"`
}

type Databases struct {
	WriteDatabase Database `json:"writeDatabase"`
	ReadDatabase  Database `json:"readDatabase"`
}

type GoRoutines struct {
	GoRoutinesCount int  `json:"goRoutinesCount"`
	UseGoRoutines   bool `json:"useGoRoutines"`
}

func GetSettings() *Settings {
	file, err := os.ReadFile("settings.json")

	if err != nil {
		fmt.Println("Erro ao carregar arquivo de configuração:", err.Error())
	}

	var settings Settings

	json.Unmarshal(file, &settings)

	return &settings
}
