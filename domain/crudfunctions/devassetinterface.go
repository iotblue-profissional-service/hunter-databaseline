package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
)

type DeviceAssetInterface interface {
	GetGlobalId() string

	GetName() string

	GetModel() string

	ValidateModel() error

	GetReferenceName() string

	GetClientId() string

	GetIP() string

	GetFeatureId() string

	GetParentAssetId() string

	GetParentGatewayId() string

	GetDeviceType() string

	GetTags() []string

	GetSearchTag() string

	SetParentAssetInfo(parentAsset cervello.Asset) error

	SetParentGatewayInfo(parentDevice cervello.Device) error

	GetLayerType() string

	GetAssetType() string

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	GetEssentialKeys() []string

	Validate() error

	GetNonDuplicatingKeys() []string

	GetParentAssetKey() string

	GetParentGatewayKey() string

	GetMac() string

	GetCommunicationProtocolConfiguration() (string, interface{})
}

func ValidateDeviceAssetEntity(obj DeviceAssetInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("globalid is missing")
	}

	if !common.IsValidUUID(obj.GetGlobalId()) {
		println(obj.GetGlobalId())
		return errors.New("globalid must be valid uuid")
	}

	if obj.GetReferenceName() == "" {
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

	if obj.GetLayerType() == "" {
		return errors.New("LayerType is missing")
	}

	if obj.GetParentAssetId() == "" {
		return errors.New("parent asset id is missing")
	}

	if !common.IsValidUUID(obj.GetParentAssetId()) {
		return errors.New("parent asset id is not valid UUID ")
	}

	//

	if obj.GetIP() != common.EmptyField {
		if obj.GetIP() == "" {
			return errors.New("ip is missing")
		}

		if !common.IsValidIp(obj.GetIP()) {
			return errors.New("ip field is not valid ip address")
		}

		//if len(obj.GetClientId()) != 16 {
		//	return errors.New("client Id must be 16 char")
		//}
	}

	if obj.GetDeviceType() == cervello.DeviceTypePeriphral {
		if obj.GetParentGatewayId() == "" {
			return errors.New("parent gateway id is missing")
		}

		if !common.IsValidUUID(obj.GetParentGatewayId()) {
			return errors.New("parent gateway id is not valid UUID ")
		}
	}

	return nil
}
