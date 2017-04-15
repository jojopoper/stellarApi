package stellarApi

import (
	"encoding/hex"
	"strings"

	_b "github.com/stellar/go/build"
	_x "github.com/stellar/go/xdr"
)

// MemoInfo memo定义
type MemoInfo struct {
	mType   _x.MemoType
	memoStr string
	memoId  uint64
}

// SetTextMemo 设置Text memo
func (ths *MemoInfo) SetTextMemo(m string) {
	ths.mType = _x.MemoTypeMemoText
	ths.memoStr = m
}

// SetIdMemo 设置Text id
func (ths *MemoInfo) SetIdMemo(m uint64) {
	ths.mType = _x.MemoTypeMemoId
	ths.memoId = m
}

// SetHashMemo 设置Text hash
func (ths *MemoInfo) SetHashMemo(m string) {
	ths.mType = _x.MemoTypeMemoHash
	val := strings.TrimPrefix(m, "0x")
	ths.memoStr = val
}

// SetReturnMemo 设置Text return hash
func (ths *MemoInfo) SetReturnMemo(m string) {
	ths.mType = _x.MemoTypeMemoReturn
	val := strings.TrimPrefix(m, "0x")
	ths.memoStr = val
}

// GetMemo 获取Memo结构
func (ths *MemoInfo) GetMemo() _b.TransactionMutator {
	switch ths.mType {
	case _x.MemoTypeMemoText:
		return _b.MemoText{Value: ths.memoStr}
	case _x.MemoTypeMemoId:
		return _b.MemoID{Value: ths.memoId}
	case _x.MemoTypeMemoHash, _x.MemoTypeMemoReturn:
		b, err := hex.DecodeString(ths.memoStr)
		if err == nil {
			hash := new(_x.Hash)
			copy(hash[:], b[0:32])
			return _b.MemoHash{Value: *hash}
		}
	}
	return nil
}

// AddMemo 添加memo到tx
func (ths *MemoInfo) AddMemo(tx *_b.TransactionBuilder) {
	m := ths.GetMemo()
	if m != nil {
		tx.Mutate(m)
	}
}
