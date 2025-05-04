package entities

type Company struct {
	ID                             uint64  `gorm:"primary_key;"`
	CompanyName                    string  `gorm:"not null"` // company's name (公司名稱)
	ResponsiblePerson              string  // company's responsible person name (公司負責人)
	GovernmentUniformInvoiceNumber string  // company's invoice number (統一編號)
	RemittanceAccount              string  // company's remittance account (分局號 + 號碼)
	OfficePhoneNumber              string  // company's office phone (公司電話)
	PersonalPhoneNumber            string  // company's office personal phone (聯絡人手機)
	OfficeAddress                  string  // company's office address(公司地址)
	CorrespondenceAddress          string  // company's correspondence address(通訊地址)
	DeliveryFee                    float64 // company's delivery fee(運送費用)
}
