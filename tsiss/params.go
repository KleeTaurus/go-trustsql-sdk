package tsiss

// IssQuery 共享信息查询参数
type IssQuery struct {
	/// 协议参数
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名

	/// 业务参数
	NodeID      string                 `json:"node_id,omitempty"`             // 节点ID,可以通过baas平台->相应链->基本信息页面中获取
	ChainID     string                 `json:"chain_id" validate:"required"`  // 链ID,可以通过baas平台->相应链->基本信息页面中获取
	LedgerID    string                 `json:"ledger_id" validate:"required"` // 账本ID,可以通过baas平台->相应链->相应账本页面中获取
	InfoKey     string                 `json:"info_key,omitempty"`            // 信息标识,查询符合此信息标识号的记录
	InfoVersion string                 `json:"info_version,omitempty"`        // XXX 信息版本号, 查询符合此信息版本号的记录
	State       string                 `json:"state,omitempty"`               // XXX 记录状态, 由业务自行定义。
	Content     map[string]interface{} `json:"content,omitempty"`             // 记录内容,可以选填部分json字段，查询符合这些字段的记录
	Notes       map[string]interface{} `json:"notes,omitempty"`               // 记录注释
	Range       map[string]interface{} `json:"range,omitempty"`               // 范围查询条件,作为查询条件，范围查询条件
	Account     string                 `json:"account,omitempty"`             // 记录方地址,查询符合此记录方地址的记录
	THash       string                 `json:"t_hash,omitempty"`              // 记录哈希,查询符合此记录哈希的记录
	PageNo      string                 `json:"page_no,omitempty"`             // XXX 页码, 第几页，默认1
	PageLimit   string                 `json:"page_limit,omitempty"`          // XXX 每页数量,分页显示时每页显示多少条，默认10
	Timestamp   string                 `json:"timestamp" validate:"required"` // XXX 请求时间戳,当前unix时间戳(秒)

	/// Range范围查询条件(string)
	BHeight    map[string]interface{} `json:"b_height,omitempty"`    // 区块高度,条件范围，区块高度范围
	CommitTime map[string]interface{} `json:"commit_time,omitempty"` // 记录时间,条件范围，记录时间范围

	/// 条件范围(JsonObject)
	From string `json:"from,omitempty"` // 开始,作为查询条件，查询这个条件之后的信息记录（>=）
	To   string `json:"to,omitempty"`   // 结束,作为查询条件，查询这个条件之前的信息记录（<=）
}

// Info 记录列表
type Info struct {
	InfoKey     string                 `json:"info_key"`     // 信息标识,同信息录入接口
	InfoVersion string                 `json:"info_version"` // 信息版本号
	State       string                 `json:"state"`        // 记录状态
	Content     map[string]interface{} `json:"content"`      // 记录内容
	Notes       map[string]interface{} `json:"notes"`        // 记录注释
	CommitTime  string                 `json:"commit_time"`  // 记录时间
	Account     string                 `json:"account"`      // 记录方地址
	PublicKey   string                 `json:"public_key"`   // 记录方公钥
	Sign        string                 `json:"sign"`         // 记录方签名
	THash       string                 `json:"t_hash"`       // 记录哈希
	BHeight     string                 `json:"b_height"`     // 区块高度
	BPrevHash   string                 `json:"b_prev_hash"`  // 前块哈希
	BHash       string                 `json:"b_hash"`       // 本块哈希
	BTime       string                 `json:"b_time"`       // 区块时间
}

// IssResponse 返回参数
type IssResponse struct {
	// 协议参数
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名
	// 业务参数
	Retcode    string `json:"retcode"`     // 返回状态码，0表示成功，其它为失败
	Retmsg     string `json:"retmsg"`      // 返回信息，如非空，为错误原因。
	TotalInfos string `json:"total_infos"` // 记录总数,符合条件的记录总条数
	Infos      []Info `json:"infos"`       // 记录列表,本次查询出的记录列表json数组
}

// IssAppend 共享信息新增/追加,请求参数
type IssAppend struct {
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名

	/// 业务参数
	NodeID      string                 `json:"node_id,omitempty"`                          // 节点ID,可以通过baas平台->相应链->基本信息页面中获取
	ChainID     string                 `json:"chain_id" validate:"required"`               // 链ID,可以通过baas平台->相应链->基本信息页面中获取
	LedgerID    string                 `json:"ledger_id" validate:"required"`              // 账本ID,可以通过baas平台->相应链->相应账本页面中获取
	InfoKey     string                 `json:"info_key,omitempty" validate:"required"`     // 信息标识,查询符合此信息标识号的记录
	InfoVersion string                 `json:"info_version,omitempty" validate:"required"` // XXX 信息版本号, 查询符合此信息版本号的记录
	State       string                 `json:"state,omitempty" validate:"required"`        // XXX 记录状态, 由业务自行定义。
	Content     map[string]interface{} `json:"content,omitempty" validate:"required"`      // 记录内容,可以选填部分json字段，查询符合这些字段的记录
	Notes       map[string]interface{} `json:"notes,omitempty" validate:"required"`        // 记录注释

	CommitTime string `json:"commit_time,omitempty" validate:"required"` // 由业务自行定义。 格式：YYYY-MM-DD HH:mm:SS
	Account    string `json:"account,omitempty" validate:"required"`     // 记录方地址,查询符合此记录方地址的记录
	PublicKey  string `json:"public_key" validate:"required"`            // 记录方公钥
	Sign       string `json:"sign,omitempty"`                            // 记录方签名,使用SDK函数IssSign进行签名
	TimeStamp  int64  `json:"timestamp" validate:"required"`             // 当前unix时间戳(秒)，与服务器时间相差过远会返回失败
}

// IssAppendResponse 共享信息新增/追加, 返回参数
type IssAppendResponse struct {
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名

	/// 业务参数
	Retcode     string                 `json:"retcode"`                // XXX 返回状态码，0表示成功，其它为失败
	Retmsg      string                 `json:"retmsg"`                 // 返回信息，如非空，为错误原因。
	InfoKey     string                 `json:"info_key,omitempty"`     // 信息标识,查询符合此信息标识号的记录
	InfoVersion string                 `json:"info_version,omitempty"` // XXX 信息版本号, 查询符合此信息版本号的记录
	State       string                 `json:"state,omitempty"`        // XXX 记录状态, 由业务自行定义。
	Content     map[string]interface{} `json:"content,omitempty"`      // 记录内容,可以选填部分json字段，查询符合这些字段的记录
	Notes       map[string]interface{} `json:"notes,omitempty"`        // 记录注释
	CommitTime  string                 `json:"commit_time,omitempty"`  // 由业务自行定义。 格式：YYYY-MM-DD HH:mm:SS
	Account     string                 `json:"account,omitempty"`      // 记录方地址,查询符合此记录方地址的记录
	PublicKey   string                 `json:"public_key"`             // 记录方公钥
	Sign        string                 `json:"sign"`                   // 记录方签名,使用SDK函数IssSign进行签名
	THash       string                 `json:"t_hash,omitempty"`       // 记录哈希
	BHeight     string                 `json:"b_height,omitempty"`     // XXX 区块高度
	BPrevHash   string                 `json:"b_prev_hash"`            // 前块哈希
	BHash       string                 `json:"b_hash"`                 // 本块哈希
	BTime       string                 `json:"b_time"`                 // 区块时间
}
