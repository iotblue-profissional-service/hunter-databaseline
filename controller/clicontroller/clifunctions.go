package clicontroller

import (
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/domain/irrigationdomain"
)

func City(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	return infrastructuredomain.UseCaseCities(csvLines, keysMap, action)
}

func Area(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	return infrastructuredomain.UseCaseAreas(csvLines, keysMap, action)
}

func HunterController(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	return irrigationdomain.UseCaseHunterController(csvLines, keysMap, action)
}

func HunterStation(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	return irrigationdomain.UseCaseHunterStation(csvLines, keysMap, action)
}
