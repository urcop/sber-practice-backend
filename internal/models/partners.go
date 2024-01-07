package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Partners struct {
	PartnerId      string      `gorm:"primaryKey" json:"partnerId,omitempty"`
	Title          string      `json:"title"`
	Conditions     StringSlice `gorm:"type:json" json:"conditions"`
	AdditionalInfo string      `json:"additionalInfo"`
	Places         []Places    `gorm:"foreignKey:PartnerID" json:"places"`
}

type Places struct {
	PlaceID     uint   `gorm:"primaryKey" json:"place_id,omitempty"`
	PartnerID   string `json:"partner_id,omitempty"`
	Address     string `json:"address"`
	Coordinates JSONB  `gorm:"type:jsonb" json:"coordinates"`
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(src interface{}) error {
	switch data := src.(type) {
	case []byte:
		return json.Unmarshal(data, &j)
	}
	return nil
}

type StringSlice []string

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}

func (ss *StringSlice) Scan(src interface{}) error {
	switch data := src.(type) {
	case []byte:
		return json.Unmarshal(data, ss)
	case string:
		return json.Unmarshal([]byte(data), ss)
	}
	return nil
}
