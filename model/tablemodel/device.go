package tablemodel

import (
	"time"
)

const (
	InStock   = 0
	Borrowed  = 1
	Repairing = 2
	Deteleted = 3
)

type Device struct {
	Id           int        `gorm:"column:id;primary_key" json:",omitempty"`
	Name         string     `gorm:"column:name"`
	Serial       string     `gorm:"column:serial"`
	Description  string     `gorm:"column:description"`
	CategoryId   int        `gorm:"column:category_id"`
	DeviceStatus int        `gorm:"column:device_status"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

type DeviceDetailUser struct {
	Id                      int        `gorm:"column:id;primary_key"`
	Name                    string     `gorm:"column:name"`
	Serial                  string     `gorm:"column:serial"`
	Description             string     `gorm:"column:description"`
	CategoryId              int        `gorm:"column:category_id"`
	CategoryName            string     `gorm:"column:category_name"`
	DeviceStatus            int        `gorm:"column:device_status"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at"`
	DeletedAt               *time.Time `gorm:"column:deleted_at"`
	DeviceReceiveDesireDate *time.Time `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate *time.Time `gorm:"column:device_receive_returns_date"`
}

type DeviceListUser struct {
	Id                      int        `gorm:"column:id;primary_key"`
	Name                    string     `gorm:"column:name"`
	Serial                  string     `gorm:"column:serial"`
	CategoryId              int        `gorm:"column:category_id"`
	CategoryName            string     `gorm:"column:category_name"`
	DeviceStatus            int        `gorm:"column:device_status"`
	Description             string     `gorm:"column:description"`
	DeviceReceiveDesireDate *time.Time `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate *time.Time `gorm:"column:device_receive_return_date"`
}

type Devices struct {
	Device
}
type DeviceIT struct {
	Id                      int        `gorm:"column:id;primary_key"`
	Name                    string     `gorm:"column:name"`
	Serial                  string     `gorm:"column:serial"`
	CategoryId              int        `gorm:"column:category_id"`
	CategoryName            string     `gorm:"column:category_name"`
	DeviceStatus            int        `gorm:"column:device_status"`
	DeviceReceiveDesireDate *time.Time `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate *time.Time `gorm:"column:device_receive_return_date"`
	UserId                  *int       `gorm:"column:user_id"`
	UserName                *string    `gorm:"column:user_name"`
}

type DeviceListIT struct {
	Total   int
	Devices []DeviceIT
}

func (Device) TableName() string {
	return "devices"
}

func (Devices) TableName() string {
	return "devices"
}
