package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
)

type LoraDeviceInterface interface {
	GetGlobalId() string

	SetGlobalId(id string)

	GetName() string

	GetModel() string

	ValidateModel() error

	GetIntegrationId() string

	GetClientId() string

	GetFeatureId() string

	GetParentAssetId() string

	GetTags() []string

	GetSearchTag() string

	SetParentAssetInfo(parentAsset cervello.Asset) error

	GetLayerType() string

	GetLoraAppId() string

	GetLoraProfileId() string

	GetMac() string

	Activate(token string) error

	SetActivationInfo(savedAsset cervello.Device)

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	GetEssentialKeys() []string

	Validate() error

	GetNonDuplicatingKeys() []string

	GetParentAssetKey() string
}

func MigrateToLoraDevice(obj LoraDeviceInterface) (cervello.Device, error) {
	toMap, err := common.StructToMap(obj)
	if err != nil {
		return cervello.Device{}, err
	}

	return cervello.Device{
		ID:                    obj.GetGlobalId(),
		Name:                  obj.GetName(),
		DeviceType:            cervello.DeviceTypeStandalone,
		Tags:                  obj.GetTags(),
		CommunicationProtocol: cervello.DeviceProtocolLoraWan,
		ReferenceName:         obj.GetIntegrationId(),
		ClientID:              obj.GetClientId(),
		CustomFields:          toMap,
		ProtocolConfigurations: cervello.LoraWanDeviceProtocolConfiguration{
			LoraProfileID:     obj.GetLoraProfileId(),
			LoraApplicationID: obj.GetLoraAppId(),
		},
	}, nil
}

func ValidateLoraDeviceEntity(obj LoraDeviceInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("GlobalId is missing")
	}

	if !common.IsValidUUID(obj.GetGlobalId()) {
		println(obj.GetGlobalId())
		return errors.New("globalid must be valid uuid")
	}

	if obj.GetIntegrationId() == "" {
		return errors.New("IntegrationId is missing")
	}

	if obj.GetFeatureId() == "" {
		return errors.New("FeatureId is missing")
	}

	//if obj.GetModel() != common.EmptyField {
	//	if obj.GetModel() == "" {
	//		return errors.New("model is missing")
	//	}
	//	if err := obj.ValidateModel(); err != nil {
	//		return err
	//	}
	//}

	if obj.GetParentAssetId() == "" {
		return errors.New("parent asset ID is missing")
	}

	if !common.IsValidUUID(obj.GetParentAssetId()) {
		return errors.New("parent asset ID must be uuid")
	}

	if obj.GetMac() == "" {
		return errors.New("mac address is missing")
	}

	if !common.IsValidMacAddress(obj.GetMac()) {
		return errors.New("invalid mac address")
	}

	if obj.GetLayerType() == "" {
		return errors.New("LayerType is missing")
	}

	return nil
}
