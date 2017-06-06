package stellarApi

import _b "github.com/stellar/go/build"

// SetOptionsOp 设置
type SetOptionsOp struct {
	SourceAccount string
	SubOperations []ISubOperations
}

// NewSetOptionsOp 创建SetOptionsOp实例
func NewSetOptionsOp(src string, v ...ISubOperations) IOperation {
	return &SetOptionsOp{
		SourceAccount: src,
		SubOperations: v,
	}
}

// AddSubOp 添加子命令操作
func (ths *SetOptionsOp) AddSubOp(src string, subOp ISubOperations) IOperation {
	if len(src) > 0 {
		ths.SourceAccount = src
	}

	if ths.SubOperations == nil {
		ths.SubOperations = make([]ISubOperations, 0)
	}
	ths.SubOperations = append(ths.SubOperations, subOp)
	return ths
}

// AddOperation 添加操作
func (ths *SetOptionsOp) AddOperation(tx *_b.TransactionBuilder) {
	if ths.SubOperations == nil {
		return
	}
	so := &_b.SetOptionsBuilder{}
	if len(ths.SourceAccount) > 0 {
		so.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	for _, subItm := range ths.SubOperations {
		subItm.AddSubOption(so)
	}
	tx.Mutate(so)
}
