package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type JSONB map[string]any

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type Click struct {
	gorm.Model
	Link   Link `gorm:"not null"`
	LinkID int  `gorm:"not null"`

	UserIP       string
	UserAgentRaw string
	Referer      string
	UtmContent   JSONB `gorm:"type:jsonb;"`
	Target       string

	// geospatial data
	Country string
	City    string
	Region  string

	// client data
	DeviceType string
	OS         string
	Browser    string
	IsBot      bool

	QueryParams    JSONB `gorm:"type:jsonb;"`
	RedirectStatus string
}
