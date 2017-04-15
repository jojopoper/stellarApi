package stellarApi

import (
	"encoding/hex"
	"fmt"

	_h "github.com/stellar/go/clients/horizon"
	_x "github.com/stellar/go/xdr"
)

// PostTransResp transaction post response结果定义
type PostTransResp struct {
	_h.TransactionSuccess
	_h.Problem
}

// GetHash 获取Hash的十六进制数组
func (ths *PostTransResp) GetHash() ([]byte, error) {
	if len(ths.Hash) == 0 {
		return nil, fmt.Errorf("Response Hash is empty")
	}
	return hex.DecodeString(ths.Hash)
}

// GetEnvelope 获取Transaction的Envelope解码内容
func (ths *PostTransResp) GetEnvelope() (*_x.TransactionEnvelope, error) {
	tr := &_x.TransactionEnvelope{}
	err := _x.SafeUnmarshalBase64(ths.Env, tr)
	return tr, err
}

// GetResult 获取Transaction的Result解码内容
func (ths *PostTransResp) GetResult() (*_x.TransactionResult, error) {
	ret := &_x.TransactionResult{}
	err := _x.SafeUnmarshalBase64(ths.Result, ret)
	return ret, err
}

// GetMeta 获取Transaction的Meta解码内容
func (ths *PostTransResp) GetMeta() (*_x.TransactionMeta, error) {
	mt := &_x.TransactionMeta{}
	err := _x.SafeUnmarshalBase64(ths.Meta, mt)
	return mt, err
}
