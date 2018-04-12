package trustsql

// Common 公共信息
type Common struct {
	MchID       string `json:"mch_id"       validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
	SeqNo       string `json:"seq_no"       validate:"required"`
	Sign        string `json:"sign"         validate:"required"`
	Type        string `json:"type"         validate:"required"`
	TimeStamp   int64  `json:"time_stamp"   validate:"required"`
	ReqData     string `json:"req_data"     validate:"required"`
}

// UserRegister 注册用户需要的信息
type UserRegister struct {
	PublicKey    string `json:"public_key"     validate:"required"`
	UserID       string `json:"user_id"        validate:"required"`
	UserFullName string `json:"user_fullName"  validate:"required"`
}

// UserInfo 获取用户信息参数
type UserInfo struct {
	UserID string `json:"user_id" validate:"required"`
}

// Account 创建用户账户参数
type Account struct {
	UserID    string `json:"user_id"        validate:"required"`
	PublicKey string `json:"public_key"     validate:"required"`
}

// Accounts 获取用户的账户地址列表参数
type Accounts struct {
	UserID    string `json:"user_id"        validate:"required"`
	State     string `json:"state"          validate:"omitempty"`
	BeginTime string `json:"begin_time"     validate:"omitempty"`
	EndTime   string `json:"end_time"       validate:"omitempty"`
	Page      int    `json:"page"           validate:"omitempty"`
	Limit     int    `json:"limit"          validate:"omitempty"`
}

// PubkeyOfAccount 获取用户的账户公钥参数
type PubkeyOfAccount struct {
	UserID         string `json:"user_id"         validate:"required"`
	AccountAddress string `json:"account_address" validate:"required"`
}
