package hpe

import (
	"encoding/json"
	"fmt"
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
	"io"
	"strings"
)

type SmartStorage struct {
	common.Entity

	OdataContext string        `json:"@odata.context"`
	OdataID      string        `json:"@odata.id"`
	OdataType    string        `json:"@odata.type"`
	Description  string        `json:"Description"`
	ID           string        `json:"Id"`
	Name         string        `json:"Name"`
	Status       common.Status `json:"Status"`
	Type         string        `json:"Type"`
	Links        struct {
		ArrayControllers struct {
			Href string `json:"href"`
		} `json:"ArrayControllers"`
		HostBusAdapters struct {
			Href string `json:"href"`
		} `json:"HostBusAdapters"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

func GetSmartStorage(system *redfish.ComputerSystem) (*SmartStorage, error) {
	var result *SmartStorage
	url := fmt.Sprintf("%s/SmartStorage", strings.TrimSuffix(system.ODataID, "/"))
	data, err := system.GetClient().Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting SmartStorage data: %w", err)
	}

	jsonPayload, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading SmartStorage data: %w", err)
	}

	err = json.Unmarshal(jsonPayload, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling SmartStorage data: %w", err)
	}

	result.SetClient(system.GetClient())

	return result, nil
}
