package stellarApi

import _b "github.com/stellar/go/build"

// InflationOp 通胀
type InflationOp struct {
	SourceAccount string
}

// NewInflationOp 创建InflationOp实例
func NewInflationOp(src string) IOperation {
	return &InflationOp{
		SourceAccount: src,
	}
}

// AddOperation 添加操作
func (ths *InflationOp) AddOperation(tx *_b.TransactionBuilder) {
	inf := &_b.InflationBuilder{}
	if len(ths.SourceAccount) > 0 {
		inf.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	tx.Mutate(inf)
}
