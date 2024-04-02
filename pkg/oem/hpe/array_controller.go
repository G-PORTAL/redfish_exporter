package hpe

import (
	"encoding/json"
	"fmt"
	"github.com/stmcginnis/gofish/common"
	"io"
)

type ArrayController struct {
	common.Entity

	OdataContext                                  string `json:"@odata.context"`
	OdataID                                       string `json:"@odata.id"`
	OdataType                                     string `json:"@odata.type"`
	AdapterType                                   string `json:"AdapterType"`
	BackupPowerSourceStatus                       string `json:"BackupPowerSourceStatus"`
	CurrentOperatingMode                          string `json:"CurrentOperatingMode"`
	Description                                   string `json:"Description"`
	EncryptionCryptoOfficerPasswordSet            bool   `json:"EncryptionCryptoOfficerPasswordSet"`
	EncryptionEnabled                             bool   `json:"EncryptionEnabled"`
	EncryptionFwLocked                            bool   `json:"EncryptionFwLocked"`
	EncryptionHasLockedVolumesMissingBootPassword bool   `json:"EncryptionHasLockedVolumesMissingBootPassword"`
	EncryptionMixedVolumesEnabled                 bool   `json:"EncryptionMixedVolumesEnabled"`
	EncryptionStandaloneModeEnabled               bool   `json:"EncryptionStandaloneModeEnabled"`
	ExternalPortCount                             int    `json:"ExternalPortCount"`
	FirmwareVersion                               struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"FirmwareVersion"`
	HardwareRevision  string        `json:"HardwareRevision"`
	ID                string        `json:"Id"`
	InternalPortCount int           `json:"InternalPortCount"`
	Location          string        `json:"Location"`
	LocationFormat    string        `json:"LocationFormat"`
	Model             string        `json:"Model"`
	Name              string        `json:"Name"`
	SerialNumber      string        `json:"SerialNumber"`
	Status            common.Status `json:"Status"`
	Type              string        `json:"Type"`
	Links             struct {
		LogicalDrives struct {
			Href string `json:"href"`
		} `json:"LogicalDrives"`
		PhysicalDrives struct {
			Href string `json:"href"`
		} `json:"PhysicalDrives"`
		StorageEnclosures struct {
			Href string `json:"href"`
		} `json:"StorageEnclosures"`
		UnconfiguredDrives struct {
			Href string `json:"href"`
		} `json:"UnconfiguredDrives"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

func (s *ArrayController) PhysicalDrives() ([]*Drive, error) {
	result := make([]*Drive, 0)

	drives, err := s.diskDrivesList()
	if err != nil {
		return nil, fmt.Errorf("error getting ArrayControllers: %w", err)
	}

	for _, member := range drives.Members {
		data, err := s.GetClient().Get(member.OdataID)
		if err != nil {
			return nil, fmt.Errorf("error getting ArrayController data: %w", err)
		}

		var drive *Drive
		jsonPayload, err := io.ReadAll(data.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading SmartStorage data: %w", err)
		}

		err = json.Unmarshal(jsonPayload, &drive)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling SmartStorage data: %w", err)
		}

		drive.SetClient(s.GetClient())

		result = append(result, drive)
	}

	return result, nil

}

func (s *ArrayController) diskDrivesList() (*DiskDrives, error) {
	var result *DiskDrives

	data, err := s.GetClient().Get(s.Links.PhysicalDrives.Href)
	if err != nil {
		return nil, fmt.Errorf("error getting physical drives: %w", err)
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
