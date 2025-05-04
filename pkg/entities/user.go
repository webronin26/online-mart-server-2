package entities

type User struct {
	ID            uint64        `gorm:"primary_key;"`
	Account       string        `gorm:"unique"`   // user's account (帳戶帳號)
	Password      string        `gorm:"not null"` // user's password (帳戶密碼)
	UserName      string        // user's name (帳戶名稱)
	AccountStatus AccountStatus `gorm:"not null"` // account's status (帳戶狀態)
	UserKind      UserKind      `gorm:"not null"` // user's kind (使用者類型)
	Token         string        `gorm:"unique"`   // user's token (帳戶的 token)
}
