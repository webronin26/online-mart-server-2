package entities

type Member struct {
	ID                 uint64 `gorm:"primary_key;"`
	MemberPhoneNumber  string // member's phone (會員手機號碼)
	MemberAddress      string // member's address (會員地址)
	MemberAvatarURL    string // member's avatar url (大頭像 url)
	MemberEmail        string // member's email (帳戶 email)
	MemberPhoneCertify bool   // member's phone is certify or not (手機是否驗證過了)
	UserID             uint64 `gorm:"not null"` // user id (對應的使用者 id)
}
