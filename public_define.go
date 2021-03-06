package stellarApi

import _b "github.com/stellar/go/build"

// OrderTypeDef 查询时使用正序或倒序的定义
type OrderTypeDef string

// HorizonServBand horizon服务器选择定义
type HorizonServBand int

const (
	BaseTxFee = 100

	StellarHorizonTestURL = "https://horizon-testnet.stellar.org:443"
	StellarHorizonLiveURL = "https://horizon.stellar.org:443"
	StellarHorizonFlyURL  = "https://api.chinastellar.com:443"

	OrigHttp       = 1
	ClientHttp     = 2
	ProxyHttp      = 3
	ConnectionHttp = 4

	AscOrderType  OrderTypeDef = "asc"
	DescOrderType OrderTypeDef = "desc"

	UnknownHorizon HorizonServBand = -1
	OfficalHorizon HorizonServBand = 0
	FlyHorizon     HorizonServBand = 1
)

// QueryParameters 查询参数定义
type QueryParameters struct {
	RequestParameters
	Size      int
	Cursor    string
	OrderType OrderTypeDef
}

// IOperation 操作接口定义
type IOperation interface {
	AddOperation(tx *_b.TransactionBuilder)
}

// Link link base defined
type Link struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
}
