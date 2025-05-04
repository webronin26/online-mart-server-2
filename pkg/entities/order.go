package entities

import "time"

type Order struct {
	ID           uint64      `gorm:"primary_key;"`
	OrderNumber  string      `gorm:"not null"`            // order's number (訂單編號)
	OrderAddress string      `gorm:"not null"`            // order's address (訂單收件人地址)
	ProductList  string      `gorm:"not null;type:JSON;"` // order's product list (訂貨商品+數量 的列表)
	TotalPrice   float64     `gorm:"not null"`            // order's total price (訂單總額)
	OrderStatus  OrderStatus `gorm:"not null"`            // order's status (訂單付款狀況)
	PaidTime     time.Time   `gorm:"type:datetime"`       // order's paid time (訂單付款時間)
	MemberID     uint64      `gorm:"not null"`            // member id (用戶 ID)
	CompanyID    uint64      `gorm:"not null"`            // company id (廠商 ID)
}
