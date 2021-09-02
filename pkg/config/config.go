package config

import (
	"strconv"
	mylog "uploader/pkg/log"

	"github.com/sasbury/mini"
)

var log = mylog.Log
var config *mini.Config
var err error

//Init configs from file
func Init(filepath string) {

	config, err = mini.LoadConfiguration(filepath)
	if err != nil {
		log.Error("error not get conf >>", err)
	}
}

//LoadHTTPConfigs  ....
func LoadHTTPConfigs() string {
	SectionName := "HTTP_SERVER"
	return loadStringFromSection(SectionName, config, "PORT", "9000")
}

//GetLogLevel ...
func GetLogLevel() string {
	return config.String("LOG_LEVEL", "")
}

// loadIntFromSection load int paparameter and log err
func loadIntFromSection(sectionName string, pgcfg *mini.Config, name string, defval string) int {
	strVal := pgcfg.StringFromSection(sectionName, name, defval)
	if defval == "" && strVal == "" {
		log.Error("5007", "Missing mandatory: Section, Parameter", sectionName, name)
		return 0
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		log.Error("5005", "Incorrect integer: Section, Parameter, Value", err, sectionName, name, strVal)
		return 0
	}
	// только положительные параметры
	if intVal < 0 {
		log.Error("5005", "Negative integer is not allowed: Section, Parameter, Value", sectionName, name, strVal)
		return 0
	}

	log.InfoDepth("Load config parameter: Section, Parameter, Value", 1, sectionName, name, intVal)

	return intVal
}

// loadStringFromSection load str paparameter and log err
func loadStringFromSection(sectionName string, pgcfg *mini.Config, name string, defval string) string {
	strVal := pgcfg.StringFromSection(sectionName, name, defval)
	if defval == "" && strVal == "" {
		log.Error("5007", "Missing mandatory: Section, Parameter", sectionName, name)
		return ""
	}
	log.InfoDepth("Load config parameter: Section, Parameter, Value", 1, sectionName, name, strVal)

	return strVal
}

// loadBoolFromSection load bool paparameter and log err
func loadBoolFromSection(sectionName string, pgcfg *mini.Config, name string, defval string) bool {
	var boolVal bool
	strVal := pgcfg.StringFromSection(sectionName, name, defval)
	if defval == "" && strVal == "" {
		log.ErrorDepth("loadBoolFromSection", 1, "5007", "Missing mandatory: Section, Parameter", sectionName, name)
		return false
	}

	if strVal != "" {
		switch strVal {
		case "true":
			boolVal = true
		case "false":
			boolVal = false
		default:
			log.ErrorDepth("loadBoolFromSection", 1, "5014", "Incorrect boolean, оnly avaliable: 'true', 'false': Section, Parameter, Value", sectionName, name, strVal)
			return false
		}
	}

	log.InfoDepth("Load config parameter: Section, Parameter, Value", 1, sectionName, name, boolVal)

	return boolVal
}
