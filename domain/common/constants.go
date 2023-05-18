package common

const (
	EmptyField           = "EMPTY"
	BmsSystem            = "bms"
	IptvSystem           = "iptv"
	DigitalSignageSystem = "digitalSignage"
	FireSystem           = "fire"
	AccessControlSystem  = "accessControl"
	PublicAddressSystem  = "publicAddress"
	UnlockedTag          = "unlocked"
	MergeAction          = "merge"
	ValidateAction       = "validate"
	CompareAction        = "compare"
	CreateOnlyAction     = "createOnly"
	UpdateOnlyAction     = "updateOnly"
	DeleteAction         = "delete"
	DeleteCsvAction      = "deleteCSV"
	DeleteOthersAction   = "deleteOthers"
	GisDevice            = "GIS"
	MockDevice           = "Mock"
	IsPhysicalDevice     = true
)

var (
	Systems = map[string]bool{
		BmsSystem:            true,
		IptvSystem:           true,
		DigitalSignageSystem: true,
		FireSystem:           true,
		AccessControlSystem:  true,
		PublicAddressSystem:  true,
	}
	Actions = map[string]bool{
		MergeAction:        true,
		ValidateAction:     true,
		CompareAction:      true,
		CreateOnlyAction:   true,
		UpdateOnlyAction:   true,
		DeleteAction:       true,
		DeleteCsvAction:    true,
		DeleteOthersAction: true,
	}
	DeleteActions = map[string]bool{
		DeleteAction:       true,
		DeleteCsvAction:    true,
		DeleteOthersAction: true,
	}
)
