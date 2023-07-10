package application

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	cli "github.com/jawher/mow.cli"
)

func StartBaselineApp() {
	// choose which service to start
	startCliInterface()
}

func StartApiApp() {
	startServer()
}

func startCliInterface() {
	helpString := `
A command line tool to baseline data between cervello IOT platform and maps in apps GIS application.
examples:
1- merge sensors from csv file:  merge Sensor /home/userName/databaselinefiles/sensors.csv  
2- validate the file data is correct: validate Sensor /home/userName/databaselinefiles/sensors.csv
3- compare the file with existing data in cervello : compare Sensor /home/userName/databaselinefiles/sensors.csv

actions:
1- merge: will create non existing devices and update existing devices from csv file
2- validate: will validate that the csv file has acceptable data format and value
3- compare: like merge but will delete the data from the system that doesnt exist in csv file
4- createOnly: will read the csv file and only create the values that doesnt exist in cervello database
5- updateOnly: will read the csv file and update the data if the same id exists in cervello database
6- delete : delete is used to delete all data of a certain layer its usage: delete [layer name] all
7- deleteCSV: delete only the data in certain csv, usage: deleteCSV [layer name] [csv path]
8- deleteOthers: delete the data that in the database but not in a certain csv, usage: deleteOthers [layer name] [csv path]


notes:
it is recommended to run validate before running any command to monitor any issues in the csv files, 
which means it is recommended to run merge or compare if the validate gives you this message ( valid document )  

examples of errors in csv files:
areaId doesnt exist in database for line 2                         : this error means that the parent area doesnt exist in cervello data
controllerId doesnt exist in database for line 1                   : this error means that the parent controller doesnt exist in cervello data
integrationId column is missing in the table                       : this error means that the table doesnt have a column named ( integrationId ) 
2022/06/16 14:58:40 validation finished                            : this means end of validation
2022/06/16 14:58:40 not valid document                             : this is validation result

available LayerNames :-
area - irrController - irrStation
`

	app := cli.App("hunterbaseline", helpString)
	action := app.StringArg("ACTION", "", "merge or validate or compare")
	layer := app.StringArg("LAYERNAME", "", "GIS layer name")
	csvFilePath := app.StringArg("CSVFILE", "", "a path to CSV file contain layer data")
	app.Action = func() {
		result, err := cliMapping(*action, *layer, *csvFilePath)
		if err != nil {
			log.Print(result)
			log.Print(err.Error())
			return
		}
		log.Print(result)
	}
	app.Run(os.Args)

	//app.Bool()
}

func startServer() {
	log.Println("starting the server")
	server := gin.Default()

	server.POST("/synch", CORSMiddleware(), AuthMiddleWare(AUTHORIZATION_METHOD), ApiMapping)
	_ = server.Run(":5050")
}
