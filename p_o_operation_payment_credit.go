package stellarApi

import _b "github.com/stellar/go/build"

// PaymentCreditOp 支付资产签名定义
type PaymentCreditOp struct {
	PaymentOp
	Issuer string
	Code   string
}

// NewPaymentCreditOp 创建PaymentCreditOp实例
func NewPaymentCreditOp(src, dest, amt, iss, code string) IOperation {
	ret := &PaymentCreditOp{
		PaymentOp: PaymentOp{
			SourceAccount: src,
			Destination:   dest,
			Amount:        amt,
		},
		Issuer: iss,
		Code:   code,
	}
	return ret.regFunc(ret.AmountOp)
}

// AmountOp 加入Amount
func (ths *PaymentCreditOp) AmountOp(p *_b.PaymentBuilder) {
	p.Mutate(
		_b.CreditAmount{
			Code:   ths.Code,
			Issuer: ths.Issuer,
			Amount: ths.Amount,
		},
	)
}
