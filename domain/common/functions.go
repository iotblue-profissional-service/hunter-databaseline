package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func IsValidIp(ip string) bool {
	result := net.ParseIP(ip)
	return result != nil
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	p, err := json.Marshal(obj)
	if err != nil {
		return map[string]interface{}{}, err
	}
	var inInterface map[string]interface{}
	err = json.Unmarshal(p, &inInterface)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return inInterface, nil
}

func IsValidMacAddress(mac string) bool {
	if _, err := net.ParseMAC(mac); err != nil {
		return false
	}
	return true
}

func ValidateAllColumnsExist(keysMap map[string]int, columnNames []string) error {
	missingColumns := ""
	newMap := map[string]int{}

	if len(columnNames) == 0 {
		missingColumns += "cant accept empty table\n"
	}

	for key, value := range keysMap {
		// + 1 is add because golang consider the element have zero value doesnt exist
		newMap[key] = value + 1
	}

	for _, colName := range columnNames {
		if _, ok := newMap[colName]; !ok {
			missingColumns += fmt.Sprintf("%s column is missing in the table\n", colName)
		}
	}

	if missingColumns != "" {
		return errors.New(missingColumns)
	}
	return nil
}

func SleepExecution() {
	time.Sleep(1 * time.Millisecond)
}

func SetupAdditionalInfo(keysMap map[string]int, essentialKeys []string, csvLine []string) map[string]interface{} {
	result := map[string]interface{}{}
	// mapping essential keys
	essentialKeysMap := map[string]bool{}
	essentialKeys = append(essentialKeys, "initialValue")

	for _, key := range essentialKeys {
		essentialKeysMap[key] = true
	}

	// construct result
	for key, value := range keysMap {
		if !essentialKeysMap[key] {
			result[key] = csvLine[value]
		}
	}
	return result
}
