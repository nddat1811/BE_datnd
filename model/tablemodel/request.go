package tablemodel

import (
	"time"
)

const (
	WaitAdminConfirm = iota
	WaitItConfirm
	ConfirmBorrow
	Return
	Reject
	Cancel
)

type Request struct {
	Id                      int        `gorm:"column:id;primary_key"`
	Title                   string     `gorm:"column:title"`
	RequestStatus           int        `gorm:"column:request_status"`
	UserId                  int        `gorm:"column:user_id"`
	BorrowReason            string     `gorm:"column:borrow_reason"`
	DeviceId                int        `gorm:"column:device_id"`
	DeviceReceiveDesireDate time.Time  `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate time.Time  `gorm:"column:device_receive_return_date"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at"`
	DeletedAt               *time.Time `gorm:"column:deleted_at"`
}

type RequestList struct {
	Id                      int       `gorm:"column:id;primary_key"`
	Title                   string    `gorm:"column:title"`
	RequestStatus           int       `gorm:"column:request_status"`
	UserName                string    `gorm:"column:user_name"`
	BorrowReason            string    `gorm:"column:borrow_reason"`
	DeviceId                int       `gorm:"column:device_id"`
	DeviceName              string    `gorm:"column:device_name"`
	DeviceReceiveDesireDate time.Time `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate time.Time `gorm:"column:device_receive_return_date"`
	CreatedAt               time.Time `gorm:"column:created_at"`
}

type Requests struct {
	Id                      int        `gorm:"column:id;primary_key"`
	Title                   string     `gorm:"column:title"`
	RequestStatus           int        `gorm:"column:request_status"`
	UserId                  int        `gorm:"column:user_id"`
	BorrowReason            string     `gorm:"column:borrow_reason"`
	DeviceId                int        `gorm:"column:device_id;"`
	DeviceReceiveDesireDate time.Time  `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate time.Time  `gorm:"column:device_receive_return_date"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at"`
	DeletedAt               *time.Time `gorm:"column:deleted_at"`
}

func (b *Request) TableName() string {
	return "requests"
}

type RequestDetailUser struct {
	ID                      int        `gorm:"column:id;primary_key"`
	Title                   string     `gorm:"column:title"`
	RequestStatus           int        `gorm:"column:request_status"`
	BorrowReason            string     `gorm:"column:borrow_reason"`
	DeviceId                int        `gorm:"column:device_id"`
	DeviceName              string     `gorm:"column:device_name"`
	DeviceSerial            string     `gorm:"column:device_serial"`
	CategoryId              int        `gorm:"column:category_id"`
	CategoryName            string     `gorm:"column:category_name"`
	DeviceReceiveDesireDate time.Time  `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate time.Time  `gorm:"column:device_receive_return_date"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at"`
	DeletedAt               *time.Time `gorm:"column:deleted_at"`
}

type RequestDetailIT struct {
	ID                      int        `gorm:"column:id;primary_key"`
	Title                   string     `gorm:"column:title"`
	RequestStatus           int        `gorm:"column:request_status"`
	BorrowReason            string     `gorm:"column:borrow_reason"`
	DeviceId                int        `gorm:"column:device_id"`
	DeviceName              string     `gorm:"column:device_name"`
	DeviceSerial            string     `gorm:"column:device_serial"`
	CategoryId              int        `gorm:"column:category_id"`
	CategoryName            string     `gorm:"column:category_name"`
	UserId                  int        `gorm:"column:user_id"`
	UserName                string     `gorm:"column:user_name"`
	DeviceReceiveDesireDate time.Time  `gorm:"column:device_receive_desire_date"`
	DeviceReceiveReturnDate time.Time  `gorm:"column:device_receive_return_date"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at"`
	DeletedAt               *time.Time `gorm:"column:deleted_at"`
}
