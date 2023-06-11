package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
)

type ModbusInterface interface {
	GetGlobalId() string

	GetName() string

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

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	GetEssentialKeys() []string

	Validate() error

	GetModbusConfig() *cervello.ModbusDeviceConfig

	GetHost() string

	GetPort() int64

	GetNonDuplicatingKeys() []string

	GetParentAssetKey() string

	GetParentGatewayKey() string

	GetMac() string
}

func MigrateToCervelloModBusDevice(obj ModbusInterface) (cervello.Device, error) {
	toMap, err := common.StructToMap(obj)
	if err != nil {
		return cervello.Device{}, err
	}

	parentGatewayId := ""
	if obj.GetDeviceType() == cervello.DeviceTypePeriphral {
		parentGatewayId = obj.GetParentGatewayId()
	}
	return cervello.Device{
		ID:              obj.GetGlobalId(),
		Name:            obj.GetName(),
		ReferenceName:   obj.GetReferenceName(),
		Tags:            obj.GetTags(),
		DeviceType:      obj.GetDeviceType(),
		CustomFields:    toMap,
		ClientID:        obj.GetClientId(),
		ParentGatewayID: parentGatewayId,
		ProtocolConfigurations: cervello.ModbusDeviceProtocolConfiguration{
			Host: obj.GetHost(),
			Port: obj.GetPort(),
		},
	}, nil
}

func ValidateModbusDeviceEntity(obj ModbusInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("GlobalId is missing")
	}

	if !common.IsValidUUID(obj.GetGlobalId()) {
		println(obj.GetGlobalId())
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
	if obj.GetHost() == "" {
		return errors.New("host is missing")
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
