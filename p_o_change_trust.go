package stellarApi

import _b "github.com/stellar/go/build"

// ChangeTrustOp 添加或删除信任
type ChangeTrustOp struct {
	AssetCode     string
	AssetIusser   string
	Limit         string
	SourceAccount string
}

// NewChangeTrustOp 创建ChangeTrustOp实例
func NewChangeTrustOp(src, code, iss, limit string) IOperation {
	return &ChangeTrustOp{
		SourceAccount: src,
		AssetCode:     code,
		AssetIusser:   iss,
		Limit:         limit,
	}
}

// AddOperation 添加操作
func (ths *ChangeTrustOp) AddOperation(tx *_b.TransactionBuilder) {
	ct := &_b.ChangeTrustBuilder{}
	ct.Mutate(
		_b.CreditAsset(ths.AssetCode, ths.AssetIusser),
	)
	ct.Mutate(_b.Limit(ths.Limit))
	if len(ths.SourceAccount) > 0 {
		ct.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	tx.Mutate(ct)
}
