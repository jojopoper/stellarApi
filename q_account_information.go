package stellarApi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	_b "github.com/stellar/go/build"
	_h "github.com/stellar/go/clients/horizon"
)

// // AssetInfo asset info
// type AssetInfo struct {
// 	Type    string `json:"asset_type"`
// 	Balance string `json:"balance"`
// 	Code    string `json:"asset_code"`
// 	Issuer  string `json:"asset_issuer"`
// }

// AccountInfo account info
type AccountInfo struct {
	HttpResponse
	_h.Problem
	_h.Account
	Balance  float64
	sequence uint64
}

// Init set address
func (ths *AccountInfo) Init(id string) {
	ths.ID = id
}

// GetInfo get base information
func (ths *AccountInfo) GetInfo(wt *sync.WaitGroup, p *RequestParameters) error {
	if wt != nil {
		defer wt.Done()
	}
	ths.isTestNet = p.UseTestNetwork
	if len(p.Address) == 0 {
		p.Address = ths.getAddr(p)
	}

	ths.SetDecodeFunc(ths.decodeFunc)
	_, err := ths.GetResponse(p)
	if err != nil {
		return err
	}
	if ths.Status == 0 {
		ths.Balance, _ = strconv.ParseFloat(ths.GetNativeBalance(), 64)
		ths.sequence, err = strconv.ParseUint(ths.Sequence, 10, 64)
		return err
	}
	errBody, _ := json.Marshal(ths.Problem)
	return fmt.Errorf("Account '%s' is not exist\r\n[%s]", ths.ID, string(errBody))
}

func (ths *AccountInfo) getAddr(p *RequestParameters) string {
	if p.UseTestNetwork {
		return fmt.Sprintf("%s/accounts/%s", HorizonTest, ths.ID)
	}
	return fmt.Sprintf("%s/accounts/%s", HorizonLive, ths.ID)
}

func (ths *AccountInfo) decodeFunc(body []byte) (interface{}, error) {
	err := json.Unmarshal(body, ths)
	if err != nil {
		err = fmt.Errorf("Decode http response body has error :\n%+v\nResponse body : [\n%+s\n]", err, string(body))
	}
	return ths, err
}

// GetNextSequence get next sequence
func (ths *AccountInfo) GetNextSequence() uint64 {
	ths.sequence++
	return ths.sequence
}

// GetCurrentSequence get current sequence
func (ths *AccountInfo) GetCurrentSequence() uint64 {
	return ths.sequence
}

// ResetSequence reset current sequence
func (ths *AccountInfo) ResetSequence() {
	ths.sequence--
}

// AddSequence add sequence
func (ths *AccountInfo) AddSequence(tx *_b.TransactionBuilder) {
	tx.Mutate(_b.Sequence{Sequence: ths.GetNextSequence()})
}

// AddSource add source account id
func (ths *AccountInfo) AddSource(tx *_b.TransactionBuilder) {
	tx.Mutate(_b.SourceAccount{AddressOrSeed: ths.ID})
}
