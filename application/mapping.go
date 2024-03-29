package application

import (
	"databaselineservice/controller/clicontroller"
	"databaselineservice/domain/common"
	"databaselineservice/utils/httperror"
	"databaselineservice/utils/httpresponse"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
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
	result, err = mapToControllerFunctions(csvLines, keysMap, action, layer)

	return result, err
}

func ApiMapping(context *gin.Context) {
	var body APIBody
	if err := context.ShouldBindJSON(&body); err != nil {
		restError := httperror.NewBadRequestError(err.Error())
		context.JSON(
			restError.Code,
			restError,
		)
		return
	}

	if len(body.Data) < 1 {
		restError := httperror.NewBadRequestError("no data sent")
		context.JSON(
			restError.Code,
			restError,
		)
		return
	}
	keysMap := make(map[string]int)
	keysMap["initialValue"] = 0

	keysMapLength := 1
	for key := range body.Data[0] {
		keysMap[key] = keysMapLength
		keysMapLength += 1
	}

	csvLines := make([][]string, 0)

	for _, entity := range body.Data {
		line := make([]string, keysMapLength)
		for key, index := range keysMap {
			if index == 0 {
				continue
			}
			line[index] = entity[key].(string)
		}
		csvLines = append(csvLines, line)
	}

	message, err := mapToControllerFunctions(csvLines, keysMap, body.Action, body.LayerName)
	if err != nil {
		restError := httperror.NewBadRequestError(err.Error())
		context.JSON(
			restError.Code,
			restError,
		)
		return
	}

	result := httpresponse.NewBaseResponse(nil, message, "POST")

	context.JSON(
		200,
		result,
	)

}

func mapToControllerFunctions(lines [][]string, keysMap map[string]int, action string, layer string) (string, error) {
	switch layer {
	case "city":
		return clicontroller.City(lines, keysMap, action)
	case "area":
		return clicontroller.Area(lines, keysMap, action)
	case "irrController":
		return clicontroller.HunterController(lines, keysMap, action)
	case "irrStation":
		return clicontroller.HunterStation(lines, keysMap, action)
	case "weatherStation":
		return clicontroller.WeatherStation(lines, keysMap, action)
	case "flowZone":
		return clicontroller.FlowZone(lines, keysMap, action)

	default:
		return "unsupported layer name", errors.New("wrong layer name")
	}
}
