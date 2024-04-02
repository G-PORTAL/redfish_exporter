package hpe

import "github.com/stmcginnis/gofish/common"

type Drive struct {
	common.Entity

	OdataContext                string   `json:"@odata.context"`
	OdataID                     string   `json:"@odata.id"`
	OdataType                   string   `json:"@odata.type"`
	BlockSizeBytes              int      `json:"BlockSizeBytes"`
	CapacityGB                  int      `json:"CapacityGB"`
	CapacityLogicalBlocks       int64    `json:"CapacityLogicalBlocks"`
	CapacityMiB                 int      `json:"CapacityMiB"`
	CarrierApplicationVersion   string   `json:"CarrierApplicationVersion"`
	CarrierAuthenticationStatus string   `json:"CarrierAuthenticationStatus"`
	CurrentTemperatureCelsius   int      `json:"CurrentTemperatureCelsius"`
	Description                 string   `json:"Description"`
	DiskDriveStatusReasons      []string `json:"DiskDriveStatusReasons"`
	EncryptedDrive              bool     `json:"EncryptedDrive"`
	FirmwareVersion             struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"FirmwareVersion"`
	ID                                string        `json:"Id"`
	InterfaceSpeedMbps                int           `json:"InterfaceSpeedMbps"`
	InterfaceType                     string        `json:"InterfaceType"`
	Location                          string        `json:"Location"`
	LocationFormat                    string        `json:"LocationFormat"`
	MaximumTemperatureCelsius         int           `json:"MaximumTemperatureCelsius"`
	MediaType                         string        `json:"MediaType"`
	Model                             string        `json:"Model"`
	Name                              string        `json:"Name"`
	PowerOnHours                      interface{}   `json:"PowerOnHours"`
	RotationalSpeedRpm                int           `json:"RotationalSpeedRpm"`
	SSDEnduranceUtilizationPercentage interface{}   `json:"SSDEnduranceUtilizationPercentage"`
	SerialNumber                      string        `json:"SerialNumber"`
	Status                            common.Status `json:"Status"`
	Type                              string        `json:"Type"`
	Links                             struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}
