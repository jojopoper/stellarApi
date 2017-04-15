package stellarApi

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"

	_h "github.com/stellar/go/clients/horizon"
	_x "github.com/stellar/go/xdr"
)

// AccountTransUnit 账户查询Transaction单元数据定义
type AccountTransUnit struct {
	Links struct {
		Self       Link `json:"self"`
		Account    Link `json:"account"`
		Ledger     Link `json:"ledger"`
		Operations Link `json:"operations"`
		Effects    Link `json:"effects"`
		Precedes   Link `json:"precedes"`
		Succeeds   Link `json:"succeeds"`
	} `json:"_links"`
	_h.Transaction
}

// GetHash 获取Hash的十六进制数组
func (ths *AccountTransUnit) GetHash() ([]byte, error) {
	if len(ths.Hash) == 0 {
		return nil, fmt.Errorf("Response Hash is empty")
	}
	return hex.DecodeString(ths.Hash)
}

// GetEnvelope 获取Transaction的Envelope解码内容
func (ths *AccountTransUnit) GetEnvelope() (*_x.TransactionEnvelope, error) {
	tr := &_x.TransactionEnvelope{}
	err := _x.SafeUnmarshalBase64(ths.EnvelopeXdr, tr)
	return tr, err
}

// GetResult 获取Transaction的Result解码内容
func (ths *AccountTransUnit) GetResult() (*_x.TransactionResult, error) {
	ret := &_x.TransactionResult{}
	err := _x.SafeUnmarshalBase64(ths.ResultXdr, ret)
	return ret, err
}

// GetResultMeta 获取Transaction的Result Meta解码内容
func (ths *AccountTransUnit) GetResultMeta() (*_x.TransactionMeta, error) {
	mt := &_x.TransactionMeta{}
	err := _x.SafeUnmarshalBase64(ths.ResultMetaXdr, mt)
	return mt, err
}

// GetFeeMeta 获取Transaction的Fee Meta解码内容
func (ths *AccountTransUnit) GetFeeMeta() (*_x.TransactionMeta, error) {
	mt := &_x.TransactionMeta{}
	err := _x.SafeUnmarshalBase64(ths.FeeMetaXdr, mt)
	return mt, err
}

// GetSourceAccount 获取Transaction的SourceAccount内容
func (ths *AccountTransUnit) GetSourceAccount() string {
	env, err := ths.GetEnvelope()
	if err == nil {
		return env.Tx.SourceAccount.Address()
	}
	return ""
}

// GetMemo 获取Transaction的Memo解码内容
func (ths *AccountTransUnit) GetMemo() (_x.MemoType, interface{}) {
	switch ths.MemoType {
	case "text":
		return _x.MemoTypeMemoText, ths.Memo
	case "id":
		ret, _ := strconv.ParseUint(ths.Memo, 10, 64)
		return _x.MemoTypeMemoId, ret
	case "hash":
		bs, _ := base64.StdEncoding.DecodeString(ths.Memo)
		return _x.MemoTypeMemoHash, bs
	case "return":
		bs, _ := base64.StdEncoding.DecodeString(ths.Memo)
		return _x.MemoTypeMemoReturn, bs
	}
	return _x.MemoTypeMemoNone, nil
}

// GetMemoString 获取Transaction的Memo解码内容
func (ths *AccountTransUnit) GetMemoString() string {
	t, m := ths.GetMemo()
	switch t {
	case _x.MemoTypeMemoText:
		return m.(string)
	case _x.MemoTypeMemoId:
		return fmt.Sprintf("%d", m.(uint64))
	case _x.MemoTypeMemoReturn, _x.MemoTypeMemoHash:
		return string(m.([]byte))
	}
	return ""
}
