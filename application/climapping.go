package application

import (
	"databaselineservice/controller/clicontroller"
	"databaselineservice/domain/common"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

func cliMapping(action string, layer string, csvFilePath string) (string, error) {
	if !common.Actions[action] {
		actionsString := ""
		for key := range common.Actions {
			actionsString += fmt.Sprintf("\n %s", key)
		}
		return "", fmt.Errorf(fmt.Sprintf("expected action to be one of %s but recieved: %s",
			actionsString,
			action))
	}

	csvLines := make([][]string, 2)
	keysMap := map[string]int{}
	keysMap["initialValue"] = 0

	result := ""
	// keysLine := make([]string, 1)

	if action != common.DeleteAction {
		csvFile, err := os.Open(csvFilePath)
		if err != nil {
			return "error opening file", err
		}

		defer csvFile.Close()

		csvLines, err = csv.NewReader(csvFile).ReadAll()
		if err != nil {
			return "", err
		}
		keysLine := csvLines[0]

		for i, key := range keysLine {
			key = strings.TrimSpace(key)
			keysMap[key] = i
		}
	}

	csvLines = csvLines[1:]
	var err error
	switch layer {
	case "city":
		result, err = clicontroller.City(csvLines, keysMap, action)
	case "area":
		result, err = clicontroller.Area(csvLines, keysMap, action)
	case "irrController":
		result, err = clicontroller.HunterController(csvLines, keysMap, action)
	case "irrStation":
		result, err = clicontroller.HunterStation(csvLines, keysMap, action)
	case "weatherStation":
		result, err = clicontroller.WeatherStation(csvLines, keysMap, action)
	default:
		return "unsupported layer name", errors.New("wrong layer name")
	}

	return result, err
}
