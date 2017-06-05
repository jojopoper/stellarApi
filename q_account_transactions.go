package stellarApi

import (
	"encoding/json"
	"fmt"
	"sync"

	_h "github.com/stellar/go/clients/horizon"
)

// AccountTransactions 账户Transactions查询定义
type AccountTransactions struct {
	Links struct {
		Self Link `json:"self"`
		Next Link `json:"next"`
		Prev Link `json:"prev"`
	} `json:"_links"`
	Embedded struct {
		Records []*AccountTransUnit `json:"records"`
	} `json:"_embedded"`
	HttpResponse
	_h.Problem
	accid           string
	TransactionSize int
}

// NewAccountTransactions 创建AccountTransactions实例
func NewAccountTransactions(accid string) *AccountTransactions {
	ret := new(AccountTransactions)
	return ret.Init(accid)
}

// Init set address
func (ths *AccountTransactions) Init(id string) *AccountTransactions {
	ths.accid = id
	return ths
}

// GetTransInfo 得到最新Account Transaction info
func (ths *AccountTransactions) GetTransInfo(wt *sync.WaitGroup, q *QueryParameters) error {
	if wt != nil {
		defer wt.Done()
	}
	ths.isTestNet = q.UseTestNetwork
	if len(q.Address) == 0 {
		q.Address = ths.getAddr(q)
	}
	ths.TransactionSize = 0
	ths.SetDecodeFunc(ths.decodeFunc)
	_, err := ths.GetResponse(&q.RequestParameters)
	if err != nil {
		return err
	}
	if ths.Status == 0 {
		return ths.formatResult(q)
	}
	errBody, _ := json.Marshal(ths.Problem)
	return fmt.Errorf("Account transaction '%s' has error : \n%s", ths.accid, string(errBody))
}

// GetAccTransUnit 得到查询后的某一个Transaction单元内容
func (ths *AccountTransactions) GetAccTransUnit(index int) *AccountTransUnit {
	if ths.TransactionSize > index {
		return ths.Embedded.Records[index]
	}
	return nil
}

func (ths *AccountTransactions) getAddr(q *QueryParameters) (ret string) {
	if q.UseTestNetwork {
		ret = fmt.Sprintf("%s/accounts/%s/transactions", HorizonTest, ths.accid)
	} else {
		ret = fmt.Sprintf("%s/accounts/%s/transactions", HorizonLive, ths.accid)
	}
	return fmt.Sprintf("%s?limit=%d&order=%s&cursor=%s", ret, q.Size, q.OrderType, q.Cursor)
}

func (ths *AccountTransactions) decodeFunc(body []byte) (interface{}, error) {
	err := json.Unmarshal(body, ths)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal AccountTransactions has error :\nerror is : %+v\nResponse body is : %s", err, string(body))
	}
	return ths, err
}

func (ths *AccountTransactions) formatResult(q *QueryParameters) error {
	ths.TransactionSize = len(ths.Embedded.Records)
	if q.OrderType == AscOrderType {
		if ths.TransactionSize > 1 {
			orgTrasInfos := ths.Embedded.Records
			ths.Embedded.Records = make([]*AccountTransUnit, 0)
			for idx := ths.TransactionSize - 1; idx >= 0; idx-- {
				ths.Embedded.Records = append(ths.Embedded.Records, orgTrasInfos[idx])
			}
		}
	}
	return nil
}
