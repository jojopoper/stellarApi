package stellarApi

import (
	"encoding/hex"

	_b "github.com/stellar/go/build"
	_sk "github.com/stellar/go/strkey"
)

// ISubOperations 子操作
type ISubOperations interface {
	AddSubOption(so *_b.SetOptionsBuilder)
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

// HomeDomainOpSub 主域名
type HomeDomainOpSub struct {
	HomeDomain string
}

// AddSubOption 添加到Operation中
func (ths *HomeDomainOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	if len(ths.HomeDomain) > 0 {
		so.Mutate(_b.HomeDomain(ths.HomeDomain))
	}
}

// WeightOpSub 权重
type WeightOpSub struct {
	Weight int
}

// AddSubOption 添加到Operation中
func (ths *WeightOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	if ths.Weight >= 0 && ths.Weight <= 255 {
		so.Mutate(_b.MasterWeight(uint32(ths.Weight)))
	}
}

// ThresholdOpSub 门限
type ThresholdOpSub struct {
	HighThreshold int
	MedThreshold  int
	LowThreshold  int
}

// AddSubOption 添加到Operation中
func (ths *ThresholdOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	threshld := &_b.Thresholds{}
	if ths.HighThreshold >= 0 && ths.HighThreshold <= 255 {
		htmp := uint32(ths.HighThreshold)
		threshld.High = &htmp
	}
	if ths.MedThreshold >= 0 && ths.MedThreshold <= 255 {
		mtmp := uint32(ths.MedThreshold)
		threshld.Medium = &mtmp
	}
	if ths.LowThreshold >= 0 && ths.LowThreshold <= 255 {
		ltmp := uint32(ths.LowThreshold)
		threshld.Low = &ltmp
	}

	if threshld.High != nil || threshld.Medium != nil || threshld.Low != nil {
		so.Mutate(threshld)
	}
}

// SignerOpSub 添加或删除签名
type SignerOpSub struct {
	Key      string
	Weight   int
	IsHas256 bool
}

// AddSubOption 添加到Operation中
func (ths *SignerOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	if ths.Weight >= 0 && ths.Weight <= 255 {
		if len(ths.Key) == 56 {
			so.Mutate(_b.AddSigner(ths.Key, uint32(ths.Weight)))
		} else if len(ths.Key) == 64 {
			byteHex, err := hex.DecodeString(ths.Key)
			if err == nil {
				addr := ""
				if ths.IsHas256 {
					addr, err = _sk.Encode(_sk.VersionByteHashX, byteHex)
					if err != nil {
						return
					}
				} else {
					addr, err = _sk.Encode(_sk.VersionByteHashTx, byteHex)
					if err != nil {
						return
					}
				}
				so.Mutate(_b.AddSigner(addr, uint32(ths.Weight)))
			}
		}
	}
}

// FlagOpSub 设置/清除标志位
type FlagOpSub struct {
	IsSet         bool
	AuthRequired  bool
	AuthRevocable bool
	AuthImmutable bool
}

// AddSubOption 添加到Operation中
func (ths *FlagOpSub) AddSubOption(so *_b.SetOptionsBuilder) {
	// var flag _b.SetFlag
	if ths.IsSet {
		if ths.AuthRequired {
			so.Mutate(_b.SetAuthRequired())
		}
		if ths.AuthRevocable {
			so.Mutate(_b.SetAuthRevocable())
		}
		if ths.AuthImmutable {
			so.Mutate(_b.SetAuthImmutable())
		}
	} else {
		if ths.AuthRequired {
			so.Mutate(_b.ClearAuthRequired())
		}
		if ths.AuthRevocable {
			so.Mutate(_b.ClearAuthRevocable())
		}
		if ths.AuthImmutable {
			so.Mutate(_b.ClearAuthImmutable())
		}
	}
}
