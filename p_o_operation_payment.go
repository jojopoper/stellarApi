package stellarApi

import _b "github.com/stellar/go/build"

// getPaymentAmountFunc 获取payment amount 方法定义
type getPaymentAmountFunc func(p *_b.PaymentBuilder)

// PaymentOp 支付签名定义
type PaymentOp struct {
	SourceAccount string
	Destination   string
	Amount        string
	getAmount     getPaymentAmountFunc
}

// NewPaymentOp 创建PaymentOp实例
func NewPaymentOp(src, dest, amt string) IOperation {
	ret := &PaymentOp{
		SourceAccount: src,
		Destination:   dest,
		Amount:        amt,
	}
	return ret.regFunc(ret.AmountOp)
}

// AddOperation 添加操作
func (ths *PaymentOp) AddOperation(tx *_b.TransactionBuilder) {
	pay := &_b.PaymentBuilder{}
	pay.Mutate(_b.Destination{AddressOrSeed: ths.Destination})
	if len(ths.SourceAccount) > 0 {
		pay.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	ths.getAmount(pay)
	tx.Mutate(pay)
}

// RegFunc 注册方法
func (ths *PaymentOp) regFunc(f getPaymentAmountFunc) IOperation {
	ths.getAmount = f
	return ths
}

// AmountOp 加入Amount
func (ths *PaymentOp) AmountOp(p *_b.PaymentBuilder) {
	p.Mutate(
		_b.NativeAmount{Amount: ths.Amount},
	)
}
