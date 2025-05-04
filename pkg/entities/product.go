package entities

type Product struct {
	ID              uint64        `gorm:"primary_key;"`
	ProductName     string        `gorm:"not null"`           // product's name (產品名稱)
	ProductImageURL string        `gorm:"not null"`           // product's image url (產品圖片)
	Summary         string        `gorm:"not null;type:text"` // product's summary (產品簡介)
	Information     string        `gorm:"not null;type:text"` // product's information (產品介紹)
	ProductPrice    float64       `gorm:"not null"`           // product's price (產品價錢)
	InventoryNumber int64         `gorm:"not null"`           // product's inventory (產品庫存數量)
	MaxBuy          int64         `gorm:"not null"`           // product's max buy (產品最大購買量)
	CompanyID       uint64        `gorm:"not null"`           // company id (廠商的 ID)
	ProductStatus   ProductStatus // product's status (產品的狀態)
	ProductType     ProductType   // product's type (產品的類型)
}
