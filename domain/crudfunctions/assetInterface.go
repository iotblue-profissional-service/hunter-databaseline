package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"encoding/json"
	"errors"
)

type AssetInterface interface {
	GetGlobalId() string

	GetName() string

	GetModel() string

	ValidateModel() error

	GetReferenceName() string

	GetFeatureId() string

	GetParentAssetId() string

	GetAssetType() string

	SetParentAssetInfo(parentAsset cervello.Asset) error

	GetLayerType() string

	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error

	GetEssentialKeys() []string

	Validate() error

	GetNonDuplicatingKeys() []string

	GetParentAssetKey() string
}

func MigrateToCervelloAsset(obj AssetInterface) (cervello.Asset, error) {
	p, err := json.Marshal(obj)
	if err != nil {
		return cervello.Asset{}, err
	}
	var inInterface map[string]interface{}
	err = json.Unmarshal(p, &inInterface)
	if err != nil {
		return cervello.Asset{}, err
	}
	clientId := ""
	if obj.GetReferenceName() != common.EmptyField {
		clientId = obj.GetReferenceName()
	}

	return cervello.Asset{
		ID:            obj.GetGlobalId(),
		Name:          obj.GetName(),
		ReferenceName: clientId,
		AssetType:     obj.GetAssetType(),
		CustomFields:  inInterface,
	}, nil
}

func ValidateAssetEntity(obj AssetInterface) error {
	if obj.GetGlobalId() == "" {
		return errors.New("globalid is missing")
	}

	if !common.IsValidUUID(obj.GetGlobalId()) {
		return errors.New("globalid must be uuid")
	}

	if obj.GetReferenceName() != common.EmptyField {
		if obj.GetReferenceName() == "" {
			return errors.New("IntegrationId is missing")
		}
	}

	if obj.GetFeatureId() != common.EmptyField {
		if obj.GetFeatureId() == "" {
			return errors.New("FeatureId is missing")
		}
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
		return errors.New("parent asset id must be uuid")
	}

	if obj.GetLayerType() == "" {
		return errors.New("LayerType is missing")
	}

	return nil
}
