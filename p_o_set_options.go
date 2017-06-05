package stellarApi

import (
	"reflect"

	_b "github.com/stellar/go/build"
)

// ISubOperations 子操作
type ISubOperations interface {
	AddSubOption(so *_b.SetOptionsBuilder)
}

// SetOptionsOp 设置
type SetOptionsOp struct {
	SourceAccount string
	SubOperations []ISubOperations
}

// InflationDestOpSub 通胀
type InflationDestOpSub struct {
	InflationDestination string
}

// AddSubOption 添加到Operation中
func (ths *InflationDestOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	if len(ths.InflationDestination) > 0 {
		so.Mutate(_b.InflationDest(ths.InflationDestination))
	}
}

// HomeDomainOpSub 通胀
type HomeDomainOpSub struct {
	HomeDomain string
}

// AddSubOption 添加到Operation中
func (ths *HomeDomainOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	if len(ths.HomeDomain) > 0 {
		so.Mutate(_b.HomeDomain(ths.HomeDomain))
	}
}

// NewSetOptionsOp 创建SetOptionsOp实例
// v 的顺序 Inflation Destination(string);
//			Home Domain(string);
//			Master Weight(int);
//			Low Threshold(int);
//			Medium Threshold(int);
//			High Threshold(int);
// 			Set Flags(SetFlags);
// 			Clear Flags(ClearFlags);
//			Sign Type(SignerType);
func NewSetOptionsOp(src string, v ...interface{}) IOperation {
	ret := &SetOptionsOp{
		SourceAccount: src,
		SubOperations: make([]ISubOperations, 0),
	}
	if len(v) >= 1 && reflect.TypeOf(v[0]).Kind() == reflect.String {
		ret.SubOperations = append(ret.SubOperations,
			&InflationDestOpSub{
				InflationDestination: v[0].(string),
			})
	}
	if len(v) >= 2 && reflect.TypeOf(v[1]).Kind() == reflect.String {
		ret.SubOperations = append(ret.SubOperations,
			&HomeDomainOpSub{
				HomeDomain: v[1].(string),
			})
	}
	return ret
}

// AddOperation 添加操作
func (ths *SetOptionsOp) AddOperation(tx *_b.TransactionBuilder) {
	so := &_b.SetOptionsBuilder{}
	if len(ths.SourceAccount) > 0 {
		so.Mutate(_b.SourceAccount{AddressOrSeed: ths.SourceAccount})
	}
	tx.Mutate(so)
}
