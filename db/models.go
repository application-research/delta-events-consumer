package db

import "time"

type DeltaNodeGeoLocation struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	Ip        string    `gorm:"column:ip" json:"ip"`
	Country   string    `gorm:"column:country" json:"country"`
	City      string    `gorm:"column:city" json:"city"`
	Region    string    `gorm:"column:region" json:"region"`
	Zip       string    `gorm:"column:zip" json:"zip"`
	Lat       float64   `gorm:"column:lat" json:"lat"`
	Lon       float64   `gorm:"column:lon" json:"lon"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
