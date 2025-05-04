package entities

// account's status (帳戶狀態)
type AccountStatus int8

const (
	InvalidStatus  AccountStatus = 1 // invalid (帳號未啟用)
	ActiveStatus   AccountStatus = 2 // active (帳號啟用)
	InActiveStatus AccountStatus = 3 // inactive (帳號停權)
	RemoveStatus   AccountStatus = 4 // remove (帳號移除)
)

func (a AccountStatus) IsValid() bool { return InvalidStatus <= a && a <= RemoveStatus }

// user's kind (使用者類型)
type UserKind int8

const (
	AdminUser    UserKind = 1 // admin user (管理者)
	MemberUser   UserKind = 2 // general member (一般用戶)
	RetailerUser UserKind = 3 // retailer user (廠商用戶)
)

func (u UserKind) IsValid() bool { return AdminUser <= u && u <= RetailerUser }

// retailer's type (廠商使用者的類型)
type RetailerType int8

const (
	RetailerAdmin       RetailerType = 1 // retailer's admin user (廠商管理者)
	RetailerEditor      RetailerType = 2 // retailer's general user(廠商小編)
	RetailerSales       RetailerType = 3 // retailer's sales user (廠商業務)
	RetailerAccountants RetailerType = 4 // retailer's accountant (廠商會計)
)

func (r RetailerType) IsValid() bool { return RetailerAdmin <= r && r <= RetailerAccountants }

// user's order status (訂單的付款狀況)
type OrderStatus int8

const (
	Paid     OrderStatus = 1 // user's order is paid (訂單已付款)
	Unpaid   OrderStatus = 2 // user's order is unpaid (訂單未付款)
	Failed   OrderStatus = 3 // user's order is failed (訂單失敗)
	Overtime OrderStatus = 4 // user's order is unpaid overtime (訂單付款超時)
)

func (o OrderStatus) IsValid() bool { return Paid <= o && o <= Overtime }

// product's status (商品上架狀況)
type ProductStatus int8

const (
	Sell   ProductStatus = 1 // product is ready for sale (商品可販售
	UnSell ProductStatus = 2 // product is not for sale (商品不可販售)
	Remove ProductStatus = 3 // product is removed (商品目前下架)
)

func (p ProductStatus) IsValid() bool { return Sell <= p && p <= Remove }

// product's type (商品的類型)
type ProductType int8

const (
	General        ProductType = 1 // product is general type (一般商品)
	LimitedEdition ProductType = 2 // product is limited edition type (限量商品)
)

func (c ProductType) IsValid() bool { return General <= c && c <= LimitedEdition }
