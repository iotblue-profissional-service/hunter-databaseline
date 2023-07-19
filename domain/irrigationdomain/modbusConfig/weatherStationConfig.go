package modbusConfig

import "databaselineservice/sdk/cervello"

var (
	WeatherStationConfig = [...]cervello.ModbusConfiguration{
		{
			Address: 40001,
			Mapping: map[string]map[string]string{
				"40001": {
					"key":  "Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40003,
			Mapping: map[string]map[string]string{
				"40003": {
					"key":  "Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40005,
			Mapping: map[string]map[string]string{
				"40005": {
					"key":  "Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40007,
			Mapping: map[string]map[string]string{
				"40007": {
					"key":  "Hour",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40009,
			Mapping: map[string]map[string]string{
				"40009": {
					"key":  "Minute",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40011,
			Mapping: map[string]map[string]string{
				"40011": {
					"key":  "Second",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40013,
			Mapping: map[string]map[string]string{
				"40013": {
					"key":  "RecordNumber",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40015,
			Mapping: map[string]map[string]string{
				"40015": {
					"key":  "BatteryVoltage",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40017,
			Mapping: map[string]map[string]string{
				"40017": {
					"key":  "EnclosureRH",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40019,
			Mapping: map[string]map[string]string{
				"40019": {
					"key":  "PanelTemperature",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40021,
			Mapping: map[string]map[string]string{
				"40021": {
					"key":  "skippedScans",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40023,
			Mapping: map[string]map[string]string{
				"40023": {
					"key":  "voltageDroppedCount",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40025,
			Mapping: map[string]map[string]string{
				"40025": {
					"key":  "WatchdogErrors",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40027,
			Mapping: map[string]map[string]string{
				"40027": {
					"key":  "lithiumBatteryVoltage",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40029,
			Mapping: map[string]map[string]string{
				"40029": {
					"key":  "programErrorsCount",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40031,
			Mapping: map[string]map[string]string{
				"40031": {
					"key":  "airTemperature",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40033,
			Mapping: map[string]map[string]string{
				"40033": {
					"key":  "PD-AvgAirTemp",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40035,
			Mapping: map[string]map[string]string{
				"40035": {
					"key":  "PD-MaxAirTemp",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40037,
			Mapping: map[string]map[string]string{
				"40037": {
					"key":  "PD-MinAirTemp",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40039,
			Mapping: map[string]map[string]string{
				"40039": {
					"key":  "relativeHumidity",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40041,
			Mapping: map[string]map[string]string{
				"40041": {
					"key":  "PD-MaxRH",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40043,
			Mapping: map[string]map[string]string{
				"40043": {
					"key":  "PD-MinRH",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40045,
			Mapping: map[string]map[string]string{
				"40045": {
					"key":  "dewPoint",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40047,
			Mapping: map[string]map[string]string{
				"40047": {
					"key":  "PD-AvgDewPoint",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40049,
			Mapping: map[string]map[string]string{
				"40049": {
					"key":  "PD-MaxDewPoint",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40051,
			Mapping: map[string]map[string]string{
				"40051": {
					"key":  "PD-MinDewPoint",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40053,
			Mapping: map[string]map[string]string{
				"40053": {
					"key":  "windSpeed",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40055,
			Mapping: map[string]map[string]string{
				"40055": {
					"key":  "windGust",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40057,
			Mapping: map[string]map[string]string{
				"40057": {
					"key":  "windGustDirection",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40059,
			Mapping: map[string]map[string]string{
				"40059": {
					"key":  "PD-AvgWindSpeed",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40061,
			Mapping: map[string]map[string]string{
				"40061": {
					"key":  "windDirection",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40063,
			Mapping: map[string]map[string]string{
				"40063": {
					"key":  "windDirectionStandardDeviation",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40065,
			Mapping: map[string]map[string]string{
				"40065": {
					"key":  "PD-MeanWindVectorDirection",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40067,
			Mapping: map[string]map[string]string{
				"40067": {
					"key":  "totalRainFall",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40069,
			Mapping: map[string]map[string]string{
				"40069": {
					"key":  "PD-TotalRainfall",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40071,
			Mapping: map[string]map[string]string{
				"40071": {
					"key":  "solarRadiation",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40073,
			Mapping: map[string]map[string]string{
				"40073": {
					"key":  "totalSolarRadiation",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40075,
			Mapping: map[string]map[string]string{
				"40075": {
					"key":  "PD-TotalSolarRadiation",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40077,
			Mapping: map[string]map[string]string{
				"40077": {
					"key":  "PD-TotalPotentialSolarRadiation",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40079,
			Mapping: map[string]map[string]string{
				"40079": {
					"key":  "PD-SunShineHours",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40081,
			Mapping: map[string]map[string]string{
				"40081": {
					"key":  "LastHourETo",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40083,
			Mapping: map[string]map[string]string{
				"40083": {
					"key":  "ETo",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40085,
			Mapping: map[string]map[string]string{
				"40085": {
					"key":  "PrevDayETo",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40087,
			Mapping: map[string]map[string]string{
				"40087": {
					"key":  "PD-GDD",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 40089,
			Mapping: map[string]map[string]string{
				"40089": {
					"key":  "PD-ChillingUnits",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
	}
)
