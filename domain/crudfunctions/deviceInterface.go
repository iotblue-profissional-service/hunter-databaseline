package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
)

type DeviceInterface interface {
	GetGlobalId() string

	GetName() string

	GetModel() string

	ValidateModel() error

	GetReferenceName() string

	GetClientId() string

	GetIP() string

	GetParentAssetId() string

	GetParentGatewayId() string

	GetDeviceType() string

	GetTags() []string

	GetSearchTag() string

	GetLayerType() string

	SetParentAssetInfo(parentAsset cervello.Asset) error

	SetParentGatewayInfo(parentDevice cervello.Device) error

	GetEssentialKeys() []string

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	Validate() error

	GetNonDuplicatingKeys() []string

	GetParentAssetKey() string

	GetParentGatewayKey() string

	GetMac() string

	GetCommunicationProtocolConfiguration() (string, interface{})
}

func MigrateToCervelloDevice(obj DeviceInterface) (cervello.Device, error) {
	toMap, err := common.StructToMap(obj)
	if err != nil {
		return cervello.Device{}, err
	}

	CommunicationProtocol, ProtocolConfigurations := obj.GetCommunicationProtocolConfiguration()
	if CommunicationProtocol == common.EmptyField {
		CommunicationProtocol = ""
		ProtocolConfigurations = nil
	}
	parentGatewayId := ""

	if obj.GetDeviceType() == cervello.DeviceTypePeriphral {
		parentGatewayId = obj.GetParentGatewayId()
	}
	return cervello.Device{
		ID:                     obj.GetGlobalId(),
		Name:                   obj.GetName(),
		ReferenceName:          obj.GetReferenceName(),
		Tags:                   obj.GetTags(),
		DeviceType:             obj.GetDeviceType(),
		CustomFields:           toMap,
		ClientID:               obj.GetClientId(),
		ParentGatewayID:        parentGatewayId,
		CommunicationProtocol:  CommunicationProtocol,
		ProtocolConfigurations: ProtocolConfigurations,
	}, nil
}

func ValidateDeviceEntity(obj DeviceInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("globalid is missing")
	}

	if !common.IsValidUUID(obj.GetGlobalId()) {
		return errors.New("globalid must be valid uuid")
	}

	if obj.GetReferenceName() == "" {
		return errors.New("IntegrationId is missing")
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
