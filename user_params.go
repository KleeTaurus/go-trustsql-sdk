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

// UserRegisterResponse 注册用户返回信息
type UserRegisterResponse struct {
	UserID       string `json:"user_id"`       // 用户ID
	PublicKey    string `json:"public_key"`    // 用户公钥
	UserAddress  string `json:"user_address"`  // 用户密钥地址
	UserFullName string `json:"user_fullName"` // 用户名
	Created      string `json:"created"`       // 创建时间
	State        string `json:"state"`         // 状态
}

// UserInfo 获取用户信息参数
type UserInfo struct {
	UserID string `json:"user_id" validate:"required"`
}

// UserInfoResponse 获取用户信息返回值
type UserInfoResponse struct {
	UserID      string `json:"user_id"`      // 用户ID
	PublicKey   string `json:"public_key"`   // 用户公钥
	UserAddress string `json:"user_address"` // 用户密钥地址
	Created     string `json:"created"`      // 创建时间
	State       string `json:"state"`        // 状态

}

// Account 创建用户账户参数
type Account struct {
	UserID    string `json:"user_id"     validate:"required"`
	PublicKey string `json:"public_key"  validate:"required"`
}

// AccountResponse 创建用户账户返回值
type AccountResponse struct {
	UserID         string `json:"user_id"`         // 用户ID
	AccountAddress string `json:"account_address"` // 创建的账户地址
	PublicKey      string `json:"public_key"`      // 用户公钥
	Created        string `json:"created"`         // 创建时间
	State          string `json:"state"`           // 状态
}

// Accounts 获取用户的账户地址列表参数
type Accounts struct {
	UserID    string `json:"user_id"     validate:"required"`
	State     string `json:"state"       validate:"omitempty"`
	BeginTime string `json:"begin_time"  validate:"omitempty"`
	EndTime   string `json:"end_time"    validate:"omitempty"`
	Page      int    `json:"page"        validate:"omitempty"`
	Limit     int    `json:"limit"       validate:"omitempty"`
}

// AccountsResponse 获取用户的账户地址列表返回值
type AccountsResponse struct {
	UserID         string `json:"user_id"`         // 用户ID
	AccountAddress string `json:"account_address"` // 创建的账户地址
	PublicKey      string `json:"public_key"`      // 用户公钥
	Created        string `json:"created"`         // 创建时间
	State          string `json:"state"`           // 状态
}

// PubkeyOfAccount 获取用户的账户公钥参数
type PubkeyOfAccount struct {
	UserID         string `json:"user_id"         validate:"required"`
	AccountAddress string `json:"account_address" validate:"required"`
}

// PubkeyOfAccountResponse 获取用户的账户公钥返回值
type PubkeyOfAccountResponse struct {
	UserID         string `json:"user_id"`         // 用户ID
	AccountAddress string `json:"account_address"` // 创建的账户地址
	PublicKey      string `json:"public_key"`      // 用户公钥
	Created        string `json:"created"`         // 创建时间
	State          string `json:"state"`           // 状态
}
