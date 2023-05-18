package infrastructuredomain

import (
	"databaselineservice/sdk/cervello"
)

func UpdateCity(cityAsset cervello.Asset) (string, error) {

	err := cervello.UpdateAsset(cityAsset.ID, cityAsset, "")
	if err != nil {
		return "", err
	}

	return "Updated successfully", nil
}
