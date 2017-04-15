package stellarApi

import (
	"encoding/json"
	"fmt"
	"sync"

	_b "github.com/stellar/go/build"
	_x "github.com/stellar/go/xdr"
)

// PostTransFrame post transaction帧定义
type PostTransFrame struct {
	MemoInfo
	AccountInfo
	PostTransResp
	Ops []IOperation
}

// NewPostTransFrame 创建新的Transaction实例
func NewPostTransFrame() *PostTransFrame {
	return &PostTransFrame{
		Ops: make([]IOperation, 0),
	}
}

// RegOperation 注册Operation接口实例
func (ths *PostTransFrame) RegOperation(op IOperation) {
	if ths.Ops == nil {
		ths.Ops = make([]IOperation, 0)
	}
	if op != nil {
		ths.Ops = append(ths.Ops, op)
	}
}

// GetSignature 获取签名
func (ths *PostTransFrame) GetSignature(key string) (string, error) {
	if ths.Ops == nil {
		return "", fmt.Errorf("There is not any operations in this transaction")
	}
	if len(key) != 56 {
		return "", fmt.Errorf("The length of signature key is must be 56")
	}

	tx := &_b.TransactionBuilder{}
	ths.AddSource(tx)
	ths.AddSequence(tx)
	for _, iop := range ths.Ops {
		iop.AddOperation(tx)
	}
	ths.addNetwork(tx)
	ths.AddMemo(tx)
	tx.TX.Fee = _x.Uint32(BaseTxFee * len(ths.Ops))
	if tx.Err != nil {
		return "", tx.Err
	}
	ret := tx.Sign(key)
	base64, err := ret.Base64()
	return base64, err
}

func (ths *PostTransFrame) addNetwork(tx *_b.TransactionBuilder) {
	if ths.isTestNet {
		tx.Mutate(_b.TestNetwork)
	} else {
		tx.Mutate(_b.PublicNetwork)
	}
}

// ExecSignature 发送signature到网络
func (ths *PostTransFrame) ExecSignature(wt *sync.WaitGroup, p *RequestParameters) error {
	if wt != nil {
		defer wt.Done()
	}
	ths.isTestNet = p.UseTestNetwork
	if len(p.Address) == 0 {
		p.Address = ths.getAddr(p)
	}
	ths.SetDecodeFunc(ths.decodeFunc)
	_, err := ths.PostResponse(p)
	if err != nil {
		return err
	}
	return nil
}

func (ths *PostTransFrame) getAddr(p *RequestParameters) string {
	if p.UseTestNetwork {
		return fmt.Sprintf("%s/transactions", StellarHorizonTestURL)
	}
	return fmt.Sprintf("%s/transactions", StellarHorizonLiveURL)
}

func (ths *PostTransFrame) decodeFunc(b []byte) (interface{}, error) {
	err := json.Unmarshal(b, &ths.PostTransResp)
	return ths, err
}
