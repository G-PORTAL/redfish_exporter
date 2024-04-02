package hpe

import "github.com/stmcginnis/gofish/common"

type DiskDrives struct {
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
