package entities

type Retailer struct {
	ID           uint64       `gorm:"primary_key;"`
	RetailerType RetailerType `gorm:"not null"` // retailer's type (廠商使用者的類型)
	UserID       uint64       `gorm:"not null"` // user id (對應的使用者 id)
	CompanyID    uint64       `gorm:"not null"` // company id (對應的廠商 id)
}
