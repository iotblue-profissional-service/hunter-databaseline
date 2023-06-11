package modbusConfig

import "databaselineservice/sdk/cervello"

var (
	WeatherStationConfig = [...]cervello.ModbusConfiguration{
		{
			Address: 56,
			Mapping: map[string]map[string]string{
				"56": {
					"key":  "ModType",
					"type": "TELEMETRY",
				},
				"57": {
					"key":  "Slot",
					"type": "TELEMETRY",
				},
				"58": {
					"key":  "ModFwVersion",
					"type": "TELEMETRY",
				},
				"59": {
					"key":  "ModEngRevision",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      4,
			Sequence:      1,
			OperationCode: 4,
		},
	}
)
