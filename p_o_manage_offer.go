package stellarApi

import (
	_b "github.com/stellar/go/build"
)

// ManageOfferBase Manage offer base 定义
type ManageOfferBase struct {
	_b.Rate
	SourceAccount string
	Amount        string
	OfferID       uint64
}

// NewAsset 创建Asset实例
func NewAsset(code, iss string) *_b.Asset {
	return &_b.Asset{
		Code:   code,
		Issuer: iss,
		Native: len(iss) == 0,
	}
}

// NewManageOfferBase 创建ManageOffer实例
func NewManageOfferBase(src string, buy, sell *_b.Asset, price, amt string, oid uint64) ManageOfferBase {
	return ManageOfferBase{
		Rate: _b.Rate{
			Selling: *sell,
			Buying:  *buy,
			Price:   _b.Price(price),
		},
		SourceAccount: src,
		Amount:        amt,
		OfferID:       oid,
	}
}

// AddOperation 添加操作
func (ths *ManageOfferBase) AddOperation(tx *_b.TransactionBuilder) {
	mo := &_b.ManageOfferBuilder{}
	if len(ths.SourceAccount) > 0 {
		mo.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	mo.Mutate(ths.Rate)
	mo.Mutate(_b.Amount(ths.Amount))
	mo.Mutate(_b.OfferID(ths.OfferID))
	tx.Mutate(mo)
}

// OfferNewOp 创建一个新的Offer
type OfferNewOp struct {
	ManageOfferBase
}

// NewOfferNewOp 创建新订单实例
func NewOfferNewOp(src string, buy, sell *_b.Asset, price, amt string) IOperation {
	return &OfferNewOp{
		ManageOfferBase: NewManageOfferBase(src, buy, sell, price, amt, 0),
	}
}

// OfferUpdateOp 更新一个Offer
type OfferUpdateOp struct {
	ManageOfferBase
}

// NewOfferUpdateOp 创建新订单实例
func NewOfferUpdateOp(src string, buy, sell *_b.Asset, price, amt string, oid uint64) IOperation {
	return &OfferUpdateOp{
		ManageOfferBase: NewManageOfferBase(src, buy, sell, price, amt, oid),
	}
}

// OfferDeleteOp 删除一个Offer
type OfferDeleteOp struct {
	ManageOfferBase
}

// NewOfferDeleteOp 创建新订单实例
func NewOfferDeleteOp(src string, buy, sell *_b.Asset, price string, oid uint64) IOperation {
	return &OfferDeleteOp{
		ManageOfferBase: NewManageOfferBase(src, buy, sell, price, "0", oid),
	}
}
