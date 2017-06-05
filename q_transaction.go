package stellarApi

import (
	"encoding/json"
	"fmt"
	"sync"

	_h "github.com/stellar/go/clients/horizon"
)

// TransactionRequest Transacion 查询定义
type TransactionRequest struct {
	AccountTransUnit
	HttpResponse
	_h.Problem
	reqHash string
}

// NewTransactionRequest 创建TransactionRequests实例
func NewTransactionRequest(hash string) *TransactionRequest {
	ret := new(TransactionRequest)
	return ret.Init(hash)
}

// Init set address
func (ths *TransactionRequest) Init(hash string) *TransactionRequest {
	ths.reqHash = hash
	return ths
}

// GetTransInfo 得到最新Account Transaction info
func (ths *TransactionRequest) GetTransInfo(wt *sync.WaitGroup, q *QueryParameters) error {
	if wt != nil {
		defer wt.Done()
	}
	ths.isTestNet = q.UseTestNetwork
	if len(q.Address) == 0 {
		q.Address = ths.getAddr(q)
	}
	ths.SetDecodeFunc(ths.decodeFunc)
	_, err := ths.GetResponse(&q.RequestParameters)
	if err != nil {
		return err
	}
	if ths.Status == 0 {
		return nil
	}
	errBody, _ := json.Marshal(ths.Problem)
	return fmt.Errorf("Transaction '%s' has error : \n%s", ths.reqHash, string(errBody))
}

func (ths *TransactionRequest) getAddr(q *QueryParameters) (ret string) {
	if q.UseTestNetwork {
		return fmt.Sprintf("%s/transactions/%s", HorizonTest, ths.reqHash)
	}
	return fmt.Sprintf("%s/transactions/%s", HorizonLive, ths.reqHash)
}

func (ths *TransactionRequest) decodeFunc(body []byte) (interface{}, error) {
	err := json.Unmarshal(body, ths)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal TransactionRequest has error :\nerror is : %+v\nResponse body is : %s", err, string(body))
	}
	return ths, err
}
