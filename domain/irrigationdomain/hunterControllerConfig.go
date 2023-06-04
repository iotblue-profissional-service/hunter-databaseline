package irrigationdomain

import "databaselineservice/sdk/cervello"

var (
	hunterControllerConfig = [...]cervello.ModbusConfiguration{
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
		{
			Address: 214,
			Mapping: map[string]map[string]string{
				"214": {
					"key":  "GblSeasAdj",
					"type": "TELEMETRY",
				},
				"215": {
					"key":  "Month1SeasAdj",
					"type": "TELEMETRY",
				},
				"216": {
					"key":  "Month2SeasAdj",
					"type": "TELEMETRY",
				},
				"217": {
					"key":  "Month3SeasAdj",
					"type": "TELEMETRY",
				},
				"218": {
					"key":  "Month4SeasAdj",
					"type": "TELEMETRY",
				},
				"219": {
					"key":  "Month5SeasAdj",
					"type": "TELEMETRY",
				},
				"220": {
					"key":  "Month6SeasAdj",
					"type": "TELEMETRY",
				},
				"221": {
					"key":  "Month7SeasAdj",
					"type": "TELEMETRY",
				},
				"222": {
					"key":  "Month8SeasAdj",
					"type": "TELEMETRY",
				},
				"223": {
					"key":  "Month9SeasAdj",
					"type": "TELEMETRY",
				},
				"224": {
					"key":  "Month10SeasAdj",
					"type": "TELEMETRY",
				},
				"225": {
					"key":  "Month11SeasAdj",
					"type": "TELEMETRY",
				},
				"226": {
					"key":  "Month12SeasAdj",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      13,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 418,
			Mapping: map[string]map[string]string{
				"418": {
					"key":  "XfmrCurDraw",
					"type": "TELEMETRY",
				},
				"419": {
					"key":  "DecMod1Current",
					"type": "TELEMETRY",
				},
				"420": {
					"key":  "DecMod2Current",
					"type": "TELEMETRY",
				},
				"421": {
					"key":  "DecMod3Current",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      4,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 1302,
			Mapping: map[string]map[string]string{
				"1302": {
					"key":  "TimeStamp",
					"type": "TELEMETRY",
				},
				"1303": {
					"key":  "activeStationsCount",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      2,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 68,
			Mapping: map[string]map[string]string{
				"68": {
					"key":  "Stn_1",
					"type": "TELEMETRY",
				},
				"69": {
					"key":  "Stn_2",
					"type": "TELEMETRY",
				},
				"70": {
					"key":  "Stn_3",
					"type": "TELEMETRY",
				},
				"71": {
					"key":  "Stn_4",
					"type": "TELEMETRY",
				},
				"72": {
					"key":  "Stn_5",
					"type": "TELEMETRY",
				},
				"73": {
					"key":  "Stn_6",
					"type": "TELEMETRY",
				},
				"74": {
					"key":  "Stn_7",
					"type": "TELEMETRY",
				},
				"75": {
					"key":  "Stn_8",
					"type": "TELEMETRY",
				},
				"76": {
					"key":  "Stn_9",
					"type": "TELEMETRY",
				},
				"77": {
					"key":  "Stn_10",
					"type": "TELEMETRY",
				},
				"78": {
					"key":  "Stn_11",
					"type": "TELEMETRY",
				},
				"79": {
					"key":  "Stn_12",
					"type": "TELEMETRY",
				},
				"80": {
					"key":  "Stn_13",
					"type": "TELEMETRY",
				},
				"81": {
					"key":  "Stn_14",
					"type": "TELEMETRY",
				},
				"82": {
					"key":  "Stn_15",
					"type": "TELEMETRY",
				},
				"83": {
					"key":  "Stn_16",
					"type": "TELEMETRY",
				},
				"84": {
					"key":  "Stn_17",
					"type": "TELEMETRY",
				},
				"85": {
					"key":  "Stn_18",
					"type": "TELEMETRY",
				},
				"86": {
					"key":  "Stn_19",
					"type": "TELEMETRY",
				},
				"87": {
					"key":  "Stn_20",
					"type": "TELEMETRY",
				},
				"88": {
					"key":  "Stn_21",
					"type": "TELEMETRY",
				},
				"89": {
					"key":  "Stn_22",
					"type": "TELEMETRY",
				},
				"90": {
					"key":  "Stn_23",
					"type": "TELEMETRY",
				},
				"91": {
					"key":  "Stn_24",
					"type": "TELEMETRY",
				},
				"92": {
					"key":  "Stn_25",
					"type": "TELEMETRY",
				},
				"93": {
					"key":  "Stn_26",
					"type": "TELEMETRY",
				},
				"94": {
					"key":  "Stn_27",
					"type": "TELEMETRY",
				},
				"95": {
					"key":  "Stn_28",
					"type": "TELEMETRY",
				},
				"96": {
					"key":  "Stn_29",
					"type": "TELEMETRY",
				},
				"97": {
					"key":  "Stn_30",
					"type": "TELEMETRY",
				},
				"98": {
					"key":  "Stn_31",
					"type": "TELEMETRY",
				},
				"99": {
					"key":  "Stn_32",
					"type": "TELEMETRY",
				},
				"100": {
					"key":  "Stn_33",
					"type": "TELEMETRY",
				},
				"101": {
					"key":  "Stn_34",
					"type": "TELEMETRY",
				},
				"102": {
					"key":  "Stn_35",
					"type": "TELEMETRY",
				},
				"103": {
					"key":  "Stn_36",
					"type": "TELEMETRY",
				},
				"104": {
					"key":  "Stn_37",
					"type": "TELEMETRY",
				},
				"105": {
					"key":  "Stn_38",
					"type": "TELEMETRY",
				},
				"106": {
					"key":  "Stn_39",
					"type": "TELEMETRY",
				},
				"107": {
					"key":  "Stn_40",
					"type": "TELEMETRY",
				},
				"108": {
					"key":  "Stn_41",
					"type": "TELEMETRY",
				},
				"109": {
					"key":  "Stn_42",
					"type": "TELEMETRY",
				},
				"110": {
					"key":  "Stn_43",
					"type": "TELEMETRY",
				},
				"111": {
					"key":  "Stn_44",
					"type": "TELEMETRY",
				},
				"112": {
					"key":  "Stn_45",
					"type": "TELEMETRY",
				},
				"113": {
					"key":  "Stn_46",
					"type": "TELEMETRY",
				},
				"114": {
					"key":  "Stn_47",
					"type": "TELEMETRY",
				},
				"115": {
					"key":  "Stn_48",
					"type": "TELEMETRY",
				},
				"116": {
					"key":  "Stn_49",
					"type": "TELEMETRY",
				},
				"117": {
					"key":  "Stn_50",
					"type": "TELEMETRY",
				},
				"118": {
					"key":  "Stn_51",
					"type": "TELEMETRY",
				},
				"119": {
					"key":  "Stn_52",
					"type": "TELEMETRY",
				},
				"120": {
					"key":  "Stn_53",
					"type": "TELEMETRY",
				},
				"121": {
					"key":  "Stn_54",
					"type": "TELEMETRY",
				},
				"122": {
					"key":  "Stn_55",
					"type": "TELEMETRY",
				},
				"123": {
					"key":  "Stn_56",
					"type": "TELEMETRY",
				},
				"124": {
					"key":  "Stn_57",
					"type": "TELEMETRY",
				},
				"125": {
					"key":  "Stn_58",
					"type": "TELEMETRY",
				},
				"126": {
					"key":  "Stn_59",
					"type": "TELEMETRY",
				},
				"127": {
					"key":  "Stn_60",
					"type": "TELEMETRY",
				},
				"128": {
					"key":  "Stn_61",
					"type": "TELEMETRY",
				},
				"129": {
					"key":  "Stn_62",
					"type": "TELEMETRY",
				},
				"130": {
					"key":  "Stn_63",
					"type": "TELEMETRY",
				},
				"131": {
					"key":  "Stn_64",
					"type": "TELEMETRY",
				},
				"132": {
					"key":  "Stn_65",
					"type": "TELEMETRY",
				},
				"133": {
					"key":  "Stn_66",
					"type": "TELEMETRY",
				},
				"134": {
					"key":  "Stn_67",
					"type": "TELEMETRY",
				},
				"135": {
					"key":  "Stn_68",
					"type": "TELEMETRY",
				},
				"136": {
					"key":  "Stn_69",
					"type": "TELEMETRY",
				},
				"137": {
					"key":  "Stn_70",
					"type": "TELEMETRY",
				},
				"138": {
					"key":  "Stn_71",
					"type": "TELEMETRY",
				},
				"139": {
					"key":  "Stn_72",
					"type": "TELEMETRY",
				},
				"140": {
					"key":  "Stn_73",
					"type": "TELEMETRY",
				},
				"141": {
					"key":  "Stn_74",
					"type": "TELEMETRY",
				},
				"142": {
					"key":  "Stn_75",
					"type": "TELEMETRY",
				},
				"143": {
					"key":  "Stn_76",
					"type": "TELEMETRY",
				},
				"144": {
					"key":  "Stn_77",
					"type": "TELEMETRY",
				},
				"145": {
					"key":  "Stn_78",
					"type": "TELEMETRY",
				},
				"146": {
					"key":  "Stn_79",
					"type": "TELEMETRY",
				},
				"147": {
					"key":  "Stn_80",
					"type": "TELEMETRY",
				},
				"148": {
					"key":  "Stn_81",
					"type": "TELEMETRY",
				},
				"149": {
					"key":  "Stn_82",
					"type": "TELEMETRY",
				},
				"150": {
					"key":  "Stn_83",
					"type": "TELEMETRY",
				},
				"151": {
					"key":  "Stn_84",
					"type": "TELEMETRY",
				},
				"152": {
					"key":  "Stn_85",
					"type": "TELEMETRY",
				},
				"153": {
					"key":  "Stn_86",
					"type": "TELEMETRY",
				},
				"154": {
					"key":  "Stn_87",
					"type": "TELEMETRY",
				},
				"155": {
					"key":  "Stn_88",
					"type": "TELEMETRY",
				},
				"156": {
					"key":  "Stn_89",
					"type": "TELEMETRY",
				},
				"157": {
					"key":  "Stn_90",
					"type": "TELEMETRY",
				},
				"158": {
					"key":  "Stn_91",
					"type": "TELEMETRY",
				},
				"159": {
					"key":  "Stn_92",
					"type": "TELEMETRY",
				},
				"160": {
					"key":  "Stn_93",
					"type": "TELEMETRY",
				},
				"161": {
					"key":  "Stn_94",
					"type": "TELEMETRY",
				},
				"162": {
					"key":  "Stn_95",
					"type": "TELEMETRY",
				},
				"163": {
					"key":  "Stn_96",
					"type": "TELEMETRY",
				},
				"164": {
					"key":  "Stn_97",
					"type": "TELEMETRY",
				},
				"165": {
					"key":  "Stn_98",
					"type": "TELEMETRY",
				},
				"166": {
					"key":  "Stn_99",
					"type": "TELEMETRY",
				},
				"167": {
					"key":  "Stn_100",
					"type": "TELEMETRY",
				},
				"168": {
					"key":  "Stn_101",
					"type": "TELEMETRY",
				},
				"169": {
					"key":  "Stn_102",
					"type": "TELEMETRY",
				},
				"170": {
					"key":  "Stn_103",
					"type": "TELEMETRY",
				},
				"171": {
					"key":  "Stn_104",
					"type": "TELEMETRY",
				},
				"172": {
					"key":  "Stn_105",
					"type": "TELEMETRY",
				},
				"173": {
					"key":  "Stn_106",
					"type": "TELEMETRY",
				},
				"174": {
					"key":  "Stn_107",
					"type": "TELEMETRY",
				},
				"175": {
					"key":  "Stn_108",
					"type": "TELEMETRY",
				},
				"176": {
					"key":  "Stn_109",
					"type": "TELEMETRY",
				},
				"177": {
					"key":  "Stn_110",
					"type": "TELEMETRY",
				},
				"178": {
					"key":  "Stn_111",
					"type": "TELEMETRY",
				},
				"179": {
					"key":  "Stn_112",
					"type": "TELEMETRY",
				},
				"180": {
					"key":  "Stn_113",
					"type": "TELEMETRY",
				},
				"181": {
					"key":  "Stn_114",
					"type": "TELEMETRY",
				},
				"182": {
					"key":  "Stn_115",
					"type": "TELEMETRY",
				},
				"183": {
					"key":  "Stn_116",
					"type": "TELEMETRY",
				},
				"184": {
					"key":  "Stn_117",
					"type": "TELEMETRY",
				},
				"185": {
					"key":  "Stn_118",
					"type": "TELEMETRY",
				},
				"186": {
					"key":  "Stn_119",
					"type": "TELEMETRY",
				},
				"187": {
					"key":  "Stn_120",
					"type": "TELEMETRY",
				},
				"188": {
					"key":  "Stn_121",
					"type": "TELEMETRY",
				},
				"189": {
					"key":  "Stn_122",
					"type": "TELEMETRY",
				},
				"190": {
					"key":  "Stn_123",
					"type": "TELEMETRY",
				},
				"191": {
					"key":  "Stn_124",
					"type": "TELEMETRY",
				},
				"192": {
					"key":  "Stn_125",
					"type": "TELEMETRY",
				},
				"193": {
					"key":  "Stn_126",
					"type": "TELEMETRY",
				},
				"194": {
					"key":  "Stn_127",
					"type": "TELEMETRY",
				},
				"195": {
					"key":  "Stn_128",
					"type": "TELEMETRY",
				},
				"196": {
					"key":  "Stn_129",
					"type": "TELEMETRY",
				},
				"197": {
					"key":  "Stn_130",
					"type": "TELEMETRY",
				},
				"198": {
					"key":  "Stn_131",
					"type": "TELEMETRY",
				},
				"199": {
					"key":  "Stn_132",
					"type": "TELEMETRY",
				},
				"200": {
					"key":  "Stn_133",
					"type": "TELEMETRY",
				},
				"201": {
					"key":  "Stn_134",
					"type": "TELEMETRY",
				},
				"202": {
					"key":  "Stn_135",
					"type": "TELEMETRY",
				},
				"203": {
					"key":  "Stn_136",
					"type": "TELEMETRY",
				},
				"204": {
					"key":  "Stn_137",
					"type": "TELEMETRY",
				},
				"205": {
					"key":  "Stn_138",
					"type": "TELEMETRY",
				},
				"206": {
					"key":  "Stn_139",
					"type": "TELEMETRY",
				},
				"207": {
					"key":  "Stn_140",
					"type": "TELEMETRY",
				},
				"208": {
					"key":  "Stn_141",
					"type": "TELEMETRY",
				},
				"209": {
					"key":  "Stn_142",
					"type": "TELEMETRY",
				},
				"210": {
					"key":  "Stn_143",
					"type": "TELEMETRY",
				},
				"211": {
					"key":  "Stn_144",
					"type": "TELEMETRY",
				},
				"212": {
					"key":  "Stn_145",
					"type": "TELEMETRY",
				},
				"213": {
					"key":  "Stn_146",
					"type": "TELEMETRY",
				},
				"214": {
					"key":  "Stn_147",
					"type": "TELEMETRY",
				},
				"215": {
					"key":  "Stn_148",
					"type": "TELEMETRY",
				},
				"216": {
					"key":  "Stn_149",
					"type": "TELEMETRY",
				},
				"217": {
					"key":  "Stn_150",
					"type": "TELEMETRY",
				},
				"218": {
					"key":  "Stn_151",
					"type": "TELEMETRY",
				},
				"219": {
					"key":  "Stn_152",
					"type": "TELEMETRY",
				},
				"220": {
					"key":  "Stn_153",
					"type": "TELEMETRY",
				},
				"221": {
					"key":  "Stn_154",
					"type": "TELEMETRY",
				},
				"222": {
					"key":  "Stn_155",
					"type": "TELEMETRY",
				},
				"223": {
					"key":  "Stn_156",
					"type": "TELEMETRY",
				},
				"224": {
					"key":  "Stn_157",
					"type": "TELEMETRY",
				},
				"225": {
					"key":  "Stn_158",
					"type": "TELEMETRY",
				},
				"226": {
					"key":  "Stn_159",
					"type": "TELEMETRY",
				},
				"227": {
					"key":  "Stn_160",
					"type": "TELEMETRY",
				},
				"228": {
					"key":  "Stn_161",
					"type": "TELEMETRY",
				},
				"229": {
					"key":  "Stn_162",
					"type": "TELEMETRY",
				},
				"230": {
					"key":  "Stn_163",
					"type": "TELEMETRY",
				},
				"231": {
					"key":  "Stn_164",
					"type": "TELEMETRY",
				},
				"232": {
					"key":  "Stn_165",
					"type": "TELEMETRY",
				},
				"233": {
					"key":  "Stn_166",
					"type": "TELEMETRY",
				},
				"234": {
					"key":  "Stn_167",
					"type": "TELEMETRY",
				},
				"235": {
					"key":  "Stn_168",
					"type": "TELEMETRY",
				},
				"236": {
					"key":  "Stn_169",
					"type": "TELEMETRY",
				},
				"237": {
					"key":  "Stn_170",
					"type": "TELEMETRY",
				},
				"238": {
					"key":  "Stn_171",
					"type": "TELEMETRY",
				},
				"239": {
					"key":  "Stn_172",
					"type": "TELEMETRY",
				},
				"240": {
					"key":  "Stn_173",
					"type": "TELEMETRY",
				},
				"241": {
					"key":  "Stn_174",
					"type": "TELEMETRY",
				},
				"242": {
					"key":  "Stn_175",
					"type": "TELEMETRY",
				},
				"243": {
					"key":  "Stn_176",
					"type": "TELEMETRY",
				},
				"244": {
					"key":  "Stn_177",
					"type": "TELEMETRY",
				},
				"245": {
					"key":  "Stn_178",
					"type": "TELEMETRY",
				},
				"246": {
					"key":  "Stn_179",
					"type": "TELEMETRY",
				},
				"247": {
					"key":  "Stn_180",
					"type": "TELEMETRY",
				},
				"248": {
					"key":  "Stn_181",
					"type": "TELEMETRY",
				},
				"249": {
					"key":  "Stn_182",
					"type": "TELEMETRY",
				},
				"250": {
					"key":  "Stn_183",
					"type": "TELEMETRY",
				},
				"251": {
					"key":  "Stn_184",
					"type": "TELEMETRY",
				},
				"252": {
					"key":  "Stn_185",
					"type": "TELEMETRY",
				},
				"253": {
					"key":  "Stn_186",
					"type": "TELEMETRY",
				},
				"254": {
					"key":  "Stn_187",
					"type": "TELEMETRY",
				},
				"255": {
					"key":  "Stn_188",
					"type": "TELEMETRY",
				},
				"256": {
					"key":  "Stn_189",
					"type": "TELEMETRY",
				},
				"257": {
					"key":  "Stn_190",
					"type": "TELEMETRY",
				},
				"258": {
					"key":  "Stn_191",
					"type": "TELEMETRY",
				},
				"259": {
					"key":  "Stn_192",
					"type": "TELEMETRY",
				},
				"260": {
					"key":  "Stn_193",
					"type": "TELEMETRY",
				},
				"261": {
					"key":  "Stn_194",
					"type": "TELEMETRY",
				},
				"262": {
					"key":  "Stn_195",
					"type": "TELEMETRY",
				},
				"263": {
					"key":  "Stn_196",
					"type": "TELEMETRY",
				},
				"264": {
					"key":  "Stn_197",
					"type": "TELEMETRY",
				},
				"265": {
					"key":  "Stn_198",
					"type": "TELEMETRY",
				},
				"266": {
					"key":  "Stn_199",
					"type": "TELEMETRY",
				},
				"267": {
					"key":  "Stn_200",
					"type": "TELEMETRY",
				},
				"268": {
					"key":  "Stn_201",
					"type": "TELEMETRY",
				},
				"269": {
					"key":  "Stn_202",
					"type": "TELEMETRY",
				},
				"270": {
					"key":  "Stn_203",
					"type": "TELEMETRY",
				},
				"271": {
					"key":  "Stn_204",
					"type": "TELEMETRY",
				},
				"272": {
					"key":  "Stn_205",
					"type": "TELEMETRY",
				},
				"273": {
					"key":  "Stn_206",
					"type": "TELEMETRY",
				},
				"274": {
					"key":  "Stn_207",
					"type": "TELEMETRY",
				},
				"275": {
					"key":  "Stn_208",
					"type": "TELEMETRY",
				},
				"276": {
					"key":  "Stn_209",
					"type": "TELEMETRY",
				},
				"277": {
					"key":  "Stn_210",
					"type": "TELEMETRY",
				},
				"278": {
					"key":  "Stn_211",
					"type": "TELEMETRY",
				},
				"279": {
					"key":  "Stn_212",
					"type": "TELEMETRY",
				},
				"280": {
					"key":  "Stn_213",
					"type": "TELEMETRY",
				},
				"281": {
					"key":  "Stn_214",
					"type": "TELEMETRY",
				},
				"282": {
					"key":  "Stn_215",
					"type": "TELEMETRY",
				},
				"283": {
					"key":  "Stn_216",
					"type": "TELEMETRY",
				},
				"284": {
					"key":  "Stn_217",
					"type": "TELEMETRY",
				},
				"285": {
					"key":  "Stn_218",
					"type": "TELEMETRY",
				},
				"286": {
					"key":  "Stn_219",
					"type": "TELEMETRY",
				},
				"287": {
					"key":  "Stn_220",
					"type": "TELEMETRY",
				},
				"288": {
					"key":  "Stn_221",
					"type": "TELEMETRY",
				},
				"289": {
					"key":  "Stn_222",
					"type": "TELEMETRY",
				},
				"290": {
					"key":  "Stn_223",
					"type": "TELEMETRY",
				},
				"291": {
					"key":  "Stn_224",
					"type": "TELEMETRY",
				},
				"292": {
					"key":  "Stn_225",
					"type": "TELEMETRY",
				},
				"293": {
					"key":  "PMV_1",
					"type": "TELEMETRY",
				},
				"294": {
					"key":  "PMV_2",
					"type": "TELEMETRY",
				},
				"295": {
					"key":  "PMV_3",
					"type": "TELEMETRY",
				},
				"296": {
					"key":  "PMV_4",
					"type": "TELEMETRY",
				},
				"297": {
					"key":  "PMV_5",
					"type": "TELEMETRY",
				},
				"298": {
					"key":  "PMV_6",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      231,
			Sequence:      1,
			OperationCode: 2,
		},
		{
			Address: 1,
			Mapping: map[string]map[string]string{
				"1": {
					"key":  "programmableOffActive-Event",
					"type": "TELEMETRY",
				},
				"2": {
					"key":  "DataReset-Event",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      2,
			Sequence:      1,
			OperationCode: 2,
		},
		{
			Address: 4,
			Mapping: map[string]map[string]string{
				"4": {
					"key":  "configurationUpdated-Event",
					"type": "TELEMETRY",
				},
				"5": {
					"key":  "controllerIrrigating-Event",
					"type": "TELEMETRY",
				},
				"6": {
					"key":  "muteActive-Event",
					"type": "TELEMETRY",
				},
				"7": {
					"key":  "suspendActive-Event",
					"type": "TELEMETRY",
				},
				"8": {
					"key":  "pauseActive-Event",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      5,
			Sequence:      1,
			OperationCode: 2,
		},
		{
			Address: 12,
			Mapping: map[string]map[string]string{
				"12": {
					"key":  "shutdownActive-Event",
					"type": "TELEMETRY",
				},
				"13": {
					"key":  "stationSizeChanged-Event",
					"type": "TELEMETRY",
				},
				"14": {
					"key":  "timeDateUpdated-Event",
					"type": "TELEMETRY",
				},
				"15": {
					"key":  "activeEventListFull-Alarm",
					"type": "TELEMETRY",
				},
				"16": {
					"key":  "programStopped-Event",
					"type": "TELEMETRY",
				},
				"17": {
					"key":  "blockStopped-Event",
					"type": "TELEMETRY",
				},
				"18": {
					"key":  "stationStopped-Event",
					"type": "TELEMETRY",
				},
				"19": {
					"key":  "irrigationStopped-Event",
					"type": "TELEMETRY",
				},
				"20": {
					"key":  "conditionResponseStatementStarted-Event",
					"type": "TELEMETRY",
				},
				"21": {
					"key":  "decoderWireTestModeActive-Event",
					"type": "TELEMETRY",
				},
				"22": {
					"key":  "ClikSensorActive-Event",
					"type": "TELEMETRY",
				},
				"23": {
					"key":  "weatherSensorActive-Event",
					"type": "TELEMETRY",
				},
				"24": {
					"key":  "controllerInventoryChanged-Event",
					"type": "TELEMETRY",
				},
				"25": {
					"key":  "flowDiagnosticActive-Event",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      14,
			Sequence:      1,
			OperationCode: 2,
		},
		{
			Address: 33,
			Mapping: map[string]map[string]string{
				"33": {
					"key":  "stationSizeZero-Event",
					"type": "TELEMETRY",
				},
				"34": {
					"key":  "powerOutageDetected-Alarm",
					"type": "TELEMETRY",
				},
				"35": {
					"key":  "MainSafeOrFlowZone-Alarm",
					"type": "TELEMETRY",
				},
				"36": {
					"key":  "ClikSensor-Alarm",
					"type": "TELEMETRY",
				},
				"37": {
					"key":  "decoderModuleOverloaded-Alarm",
					"type": "TELEMETRY",
				},
				"38": {
					"key":  "stationFault-Alarm",
					"type": "TELEMETRY",
				},
				"39": {
					"key":  "P/MVFault-Alarm",
					"type": "TELEMETRY",
				},
				"40": {
					"key":  "weatherSensorCommunicationsFault-Alarm",
					"type": "TELEMETRY",
				},
				"41": {
					"key":  "RTCFault-Alarm",
					"type": "TELEMETRY",
				},
				"42": {
					"key":  "maxTransformerCurrentExceeded-Alarm",
					"type": "TELEMETRY",
				},
				"43": {
					"key":  "CANBusFault-Alarm",
					"type": "TELEMETRY",
				},
				"44": {
					"key":  "lowVoltageFault-Alarm",
					"type": "TELEMETRY",
				},
				"45": {
					"key":  "WeatherSensorFault-Alarm",
					"type": "TELEMETRY",
				},
				"46": {
					"key":  "stationFlowFault-Alarm",
					"type": "TELEMETRY",
				},
				"47": {
					"key":  "ClikSensorRainDelay-Alarm",
					"type": "TELEMETRY",
				},
				"48": {
					"key":  "NWWViolation-Alarm",
					"type": "TELEMETRY",
				},
				"49": {
					"key":  "sensorDecoderFault-Alarm",
					"type": "TELEMETRY",
				},
				"50": {
					"key":  "weatherSensorRainDelay-Alarm",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      18,
			Sequence:      1,
			OperationCode: 2,
		},
		{
			Address: 1196,
			Mapping: map[string]map[string]string{
				"1196": {
					"key":  "FSen1Freq",
					"type": "TELEMETRY",
				},
				"1197": {
					"key":  "FSen1FlowRate",
					"type": "TELEMETRY",
				},
				"1198": {
					"key":  "FSen2Freq",
					"type": "TELEMETRY",
				},
				"1199": {
					"key":  "FSen2FlowRate",
					"type": "TELEMETRY",
				},
				"1200": {
					"key":  "FSen3Freq",
					"type": "TELEMETRY",
				},
				"1201": {
					"key":  "FSen3FlowRate",
					"type": "TELEMETRY",
				},
				"1202": {
					"key":  "FSen4Freq",
					"type": "TELEMETRY",
				},
				"1203": {
					"key":  "FSen4FlowRate",
					"type": "TELEMETRY",
				},
				"1204": {
					"key":  "FSen5Freq",
					"type": "TELEMETRY",
				},
				"1205": {
					"key":  "FSen5FlowRate",
					"type": "TELEMETRY",
				},
				"1206": {
					"key":  "FSen6Freq",
					"type": "TELEMETRY",
				},
				"1207": {
					"key":  "FSen6FlowRate",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      12,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2001,
			Mapping: map[string]map[string]string{
				"2001": {
					"key":  "flowTimeStampDay",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2003,
			Mapping: map[string]map[string]string{
				"2003": {
					"key":  "totalFlowDay",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2005,
			Mapping: map[string]map[string]string{
				"2005": {
					"key":  "flowSensor1Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2007,
			Mapping: map[string]map[string]string{
				"2007": {
					"key":  "flowSensor2Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2009,
			Mapping: map[string]map[string]string{
				"2009": {
					"key":  "flowSensor3Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2011,
			Mapping: map[string]map[string]string{
				"2011": {
					"key":  "flowSensor4Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2013,
			Mapping: map[string]map[string]string{
				"2013": {
					"key":  "flowSensor5Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2015,
			Mapping: map[string]map[string]string{
				"2015": {
					"key":  "flowSensor6Day",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2017,
			Mapping: map[string]map[string]string{
				"2017": {
					"key":  "flowTimeStampWeek",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2019,
			Mapping: map[string]map[string]string{
				"2019": {
					"key":  "totalflowWeek",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2021,
			Mapping: map[string]map[string]string{
				"2021": {
					"key":  "flowSensor1Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2023,
			Mapping: map[string]map[string]string{
				"2023": {
					"key":  "flowSensor2Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2025,
			Mapping: map[string]map[string]string{
				"2025": {
					"key":  "flowSensor3Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2027,
			Mapping: map[string]map[string]string{
				"2027": {
					"key":  "flowSensor4Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2029,
			Mapping: map[string]map[string]string{
				"2029": {
					"key":  "flowSensor5Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2031,
			Mapping: map[string]map[string]string{
				"2031": {
					"key":  "flowSensor6Week",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2033,
			Mapping: map[string]map[string]string{
				"2033": {
					"key":  "flowTimeStampMonth",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2035,
			Mapping: map[string]map[string]string{
				"2035": {
					"key":  "totalflowMonth",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2037,
			Mapping: map[string]map[string]string{
				"2037": {
					"key":  "flowSensor1Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2039,
			Mapping: map[string]map[string]string{
				"2039": {
					"key":  "flowSensor2Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2041,
			Mapping: map[string]map[string]string{
				"2041": {
					"key":  "flowSensor3Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2043,
			Mapping: map[string]map[string]string{
				"2043": {
					"key":  "flowSensor4Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2045,
			Mapping: map[string]map[string]string{
				"2045": {
					"key":  "flowSensor5Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2047,
			Mapping: map[string]map[string]string{
				"2047": {
					"key":  "flowSensor6Month",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2049,
			Mapping: map[string]map[string]string{
				"2049": {
					"key":  "flowTimeStampYear",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2051,
			Mapping: map[string]map[string]string{
				"2051": {
					"key":  "totalflowYear",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2053,
			Mapping: map[string]map[string]string{
				"2053": {
					"key":  "flowSensor1Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2055,
			Mapping: map[string]map[string]string{
				"2055": {
					"key":  "flowSensor2Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2057,
			Mapping: map[string]map[string]string{
				"2057": {
					"key":  "flowSensor3Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2059,
			Mapping: map[string]map[string]string{
				"2059": {
					"key":  "flowSensor4Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2061,
			Mapping: map[string]map[string]string{
				"2061": {
					"key":  "flowSensor5Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 2063,
			Mapping: map[string]map[string]string{
				"2063": {
					"key":  "flowSensor6Year",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      1,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 1285,
			Mapping: map[string]map[string]string{
				"1285": {
					"key":  "controllerFwVers",
					"type": "TELEMETRY",
				},
				"1286": {
					"key":  "controllerEngRev",
					"type": "TELEMETRY",
				},
				"1287": {
					"key":  "controllerModel",
					"type": "TELEMETRY",
				},
				"1288": {
					"key":  "controllerType",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      4,
			Sequence:      1,
			OperationCode: 4,
		},
		{
			Address: 1214,
			Mapping: map[string]map[string]string{
				"1214": {
					"key":  "DecMod1Num",
					"type": "TELEMETRY",
				},
				"1215": {
					"key":  "DecMod1Installed",
					"type": "TELEMETRY",
				},
				"1216": {
					"key":  "DecMod1PathStatus",
					"type": "TELEMETRY",
				},
				"1217": {
					"key":  "DecMod1OutputMode",
					"type": "TELEMETRY",
				},
				"1218": {
					"key":  "DecMod1WateringStatus",
					"type": "TELEMETRY",
				},
				"1219": {
					"key":  "DecMod1OverloadStatus-alarm",
					"type": "TELEMETRY",
				},
				"1220": {
					"key":  "DecMod1A/DFullScale",
					"type": "TELEMETRY",
				},
				"1221": {
					"key":  "DecMod1Ph1/2Mismatch",
					"type": "TELEMETRY",
				},
				"1222": {
					"key":  "DecMod1LockOutTime",
					"type": "TELEMETRY",
				},
				"1223": {
					"key":  "DecMod2Num",
					"type": "TELEMETRY",
				},
				"1224": {
					"key":  "DecMod2Installed",
					"type": "TELEMETRY",
				},
				"1225": {
					"key":  "DecMod2PathStatus",
					"type": "TELEMETRY",
				},
				"1226": {
					"key":  "DecMod2OutputMode",
					"type": "TELEMETRY",
				},
				"1227": {
					"key":  "DecMod2WateringStatus",
					"type": "TELEMETRY",
				},
				"1228": {
					"key":  "DecMod2OverloadStatus-alarm",
					"type": "TELEMETRY",
				},
				"1229": {
					"key":  "DecMod2A/DFullScale",
					"type": "TELEMETRY",
				},
				"1230": {
					"key":  "DecMod2Ph1/2Mismatch",
					"type": "TELEMETRY",
				},
				"1231": {
					"key":  "DecMod2LockOutTime",
					"type": "TELEMETRY",
				},
				"1232": {
					"key":  "DecMod3Num",
					"type": "TELEMETRY",
				},
				"1233": {
					"key":  "DecMod3Installed",
					"type": "TELEMETRY",
				},
				"1234": {
					"key":  "DecMod3PathStatus",
					"type": "TELEMETRY",
				},
				"1235": {
					"key":  "DecMod3OutputMode",
					"type": "TELEMETRY",
				},
				"1236": {
					"key":  "DecMod3WateringStatus",
					"type": "TELEMETRY",
				},
				"1237": {
					"key":  "DecMod3OverloadStatus-alarm",
					"type": "TELEMETRY",
				},
				"1238": {
					"key":  "DecMod3A/DFullScale",
					"type": "TELEMETRY",
				},
				"1239": {
					"key":  "DecMod3Ph1/2Mismatch",
					"type": "TELEMETRY",
				},
				"1240": {
					"key":  "DecMod3LockOutTime",
					"type": "TELEMETRY",
				},
			},
			SlaveID:       1,
			Quantity:      27,
			Sequence:      1,
			OperationCode: 4,
		},
	}
)
