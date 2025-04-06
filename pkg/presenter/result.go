package presenter

type Result struct {
	Data  interface{} `json:"data"`
	Code  StatusCode  `json:"code"`
	Count int         `json:"count"`
}

type StatusCode uint16

// return http code
// 回應的 http 狀態碼
const (
	StatusSuccess = 2001 // 成功

	StatusParamError                         = 4001  // param error / 參數有誤
	StatusParamValidateFailed                = 4002  // param validate error / 參數數值有誤
	StatusBindFailed                         = 4003  // param binding error 參數綁定失敗
	StatusRecordNotFound                     = 4004  // no record / 找不到紀錄
	StatusCreatePostFailed                   = 4005  // create post error / Post 建立失敗
	StatusCreatePostFailedInventoryNumber    = 40051 // create post error (inventor number) / Post 建立失敗，庫存數量問題
	StatusCreatePostFailedCreateRecord       = 40052 // create post error (create record) / Post 建立失敗，建立新紀錄問題
	StatusCreatePostFailedProductCompany     = 40053 // create post error (product company) / Post 建立失敗，商品廠商異常
	StatusCreatePostFailedConvert            = 40054 // create post error (convert) / Post 建立失敗，資料轉型失敗
	StatusCreatePostFailedCreateRecordName   = 40055 // create post error (create record name) / Post 建立失敗，建立新紀錄名稱問題
	StatusUpdatePostFailed                   = 4006  // update post error / Post 更新失敗
	StatusUpdatePostFailedInventoryNotEnough = 40061 // update post error (inventor not enough) / Post 更新失敗，庫存不足

	StatusTokenError   = 4011 // Token error / Token 有誤
	StatusAccountError = 4041 // account access error / 帳密認證有誤

	StatusServerError        = 5001  // server error / 伺服器錯誤
	StatusSQLError           = 5002  // sql error / 資料庫錯誤
	StatusSQLErrorScanFailed = 50021 // sql scan error / 資料庫錯誤，轉型錯誤
)
