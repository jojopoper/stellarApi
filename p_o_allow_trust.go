package stellarApi

import _b "github.com/stellar/go/build"

// AllowTrustOp 添加或删除信任
type AllowTrustOp struct {
	AssetCode     string
	Trustor       string
	SourceAccount string
	Authorize     bool
}

// NewAllowTrustOp 创建AllowTrustOp实例
func NewAllowTrustOp(src, code, trustor string, a bool) IOperation {
	return &AllowTrustOp{
		SourceAccount: src,
		AssetCode:     code,
		Trustor:       trustor,
		Authorize:     a,
	}
}

// AddOperation 添加操作
func (ths *AllowTrustOp) AddOperation(tx *_b.TransactionBuilder) {
	at := &_b.AllowTrustBuilder{}
	at.Mutate(
		_b.Trustor{Address: ths.Trustor},
		_b.AllowTrustAsset{Code: ths.AssetCode},
		_b.Authorize{Value: ths.Authorize},
	)
	if len(ths.SourceAccount) > 0 {
		at.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	tx.Mutate(at)
}
