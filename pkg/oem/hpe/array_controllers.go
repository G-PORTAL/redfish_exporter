package hpe

import (
	"encoding/json"
	"fmt"
	"github.com/stmcginnis/gofish/common"
	"io"
)

type ArrayControllers struct {
	common.Entity

	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	MemberType   string `json:"MemberType"`
	Members      []struct {
		OdataID string `json:"@odata.id"`
	} `json:"Members"`
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
	Total             int    `json:"Total"`
	Type              string `json:"Type"`
	Links             struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

func (s *SmartStorage) ArrayControllers() ([]*ArrayController, error) {
	result := make([]*ArrayController, 0)

	controllers, err := s.arrayControllerList()
	if err != nil {
		return nil, fmt.Errorf("error getting ArrayControllers: %w", err)
	}

	for _, member := range controllers.Members {
		data, err := s.GetClient().Get(member.OdataID)
		if err != nil {
			return nil, fmt.Errorf("error getting ArrayController data: %w", err)
		}

		var controller *ArrayController
		jsonPayload, err := io.ReadAll(data.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading SmartStorage data: %w", err)
		}

		err = json.Unmarshal(jsonPayload, &controller)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling SmartStorage data: %w", err)
		}

		controller.SetClient(s.GetClient())

		result = append(result, controller)
	}

	return result, nil
}

func (s *SmartStorage) arrayControllerList() (*ArrayControllers, error) {
	var result *ArrayControllers

	data, err := s.GetClient().Get(s.Links.ArrayControllers.Href)
	if err != nil {
		return nil, fmt.Errorf("error getting ArrayControllers: %w", err)
	}

	jsonPayload, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading SmartStorage data: %w", err)
	}

	err = json.Unmarshal(jsonPayload, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling SmartStorage data: %w", err)
	}

	result.SetClient(s.GetClient())

	return result, nil
}
