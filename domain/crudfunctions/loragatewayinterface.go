package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
)

type LoraGatewayInterface interface {
	GetGlobalId() string

	GetName() string

	GetModel() string

	ValidateModel() error

	GetDescription() (string, error)

	GetParentAssetId() string

	SetParentAssetInfo(parentAsset cervello.Asset) error

	GetIntegrationId() string

	GetMAC() string

	GetFeatureId() string

	GetLayerType() string

	GetIP() string

	Validate() error

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	GetEssentialKeys() []string

	GetParentAssetKey() string
}

func MigrateToCervelloLoraGateway(obj LoraGatewayInterface) (cervello.LoraGateway, error) {
	description, err := obj.GetDescription()
	if err != nil {
		return cervello.LoraGateway{}, err
	}
	return cervello.LoraGateway{
		ID:          obj.GetGlobalId(),
		Name:        obj.GetName(),
		Description: description,
	}, nil
}

func ValidateLoraGatewayEntity(obj LoraGatewayInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("GlobalId is missing")
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
		return errors.New("parent asset id is missing")
	}

	if !common.IsValidUUID(obj.GetParentAssetId()) {
		return errors.New("parent asset id  must be uuid")
	}

	if obj.GetMAC() == "" {
		return errors.New("mac address is missing")
	}

	if !common.IsValidMacAddress(obj.GetMAC()) {
		return errors.New("invalid mac address")
	}

	if obj.GetLayerType() == "" {
		return errors.New("LayerType is missing")
	}

	if obj.GetIP() == "" {
		return errors.New("ip is missing")
	}

	if !common.IsValidIp(obj.GetIP()) {
		return errors.New("ip field is not valid ip address")
	}

	return nil
}
