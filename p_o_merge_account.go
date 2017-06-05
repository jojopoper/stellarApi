package stellarApi

import (
	_b "github.com/stellar/go/build"
)

// MergeAccountOp 创建账户定义
type MergeAccountOp struct {
	Destination   string
	SourceAccount string
}

// NewMergeAccOp 创建MerageAccOp实例
func NewMergeAccOp(src, dest string) IOperation {
	return &MergeAccountOp{
		SourceAccount: src,
		Destination:   dest,
	}
}

// AddOperation 添加操作
func (ths *MergeAccountOp) AddOperation(tx *_b.TransactionBuilder) {
	amb := &_b.AccountMergeBuilder{}
	amb.Mutate(_b.Destination{AddressOrSeed: ths.Destination})
	if len(ths.SourceAccount) > 0 {
		amb.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	tx.Mutate(amb)
}
