package stellarApi

import (
	_b "github.com/stellar/go/build"
)

// CreateAccountOp 创建账户定义
type CreateAccountOp struct {
	Destination     string
	StartingBalance string
	SourceAccount   string
}

// NewCreateAccount 创建PaymentOp实例
func NewCreateAccountOp(src, dest, amt string) IOperation {
	return &CreateAccountOp{
		SourceAccount:   src,
		Destination:     dest,
		StartingBalance: amt,
	}
}

// AddOperation 添加操作
func (ths *CreateAccountOp) AddOperation(tx *_b.TransactionBuilder) {
	ca := &_b.CreateAccountBuilder{}
	ca.Mutate(_b.Destination{AddressOrSeed: ths.Destination})
	if len(ths.SourceAccount) > 0 {
		ca.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	ca.Mutate(_b.NativeAmount{Amount: ths.StartingBalance})
	tx.Mutate(ca)
}
