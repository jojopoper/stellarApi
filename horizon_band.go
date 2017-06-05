package stellarApi

// HorizonTest Horizon 测试服务器地址
var HorizonTest = StellarHorizonTestURL

// HorizonLive Horizon 生产服务器地址
var HorizonLive = StellarHorizonFlyURL

// SetHorizonBand 设置Horizon服务器绑定
// 当前只影响Live，对Test只使用官方服务器
func SetHorizonBand(b HorizonServBand) {
	switch b {
	case OfficalHorizon:
		HorizonLive = StellarHorizonLiveURL
	case FlyHorizon:
		HorizonLive = StellarHorizonFlyURL
	}
}

// GetHorizonBand 获取当前Horizon服务器绑定状态
func GetHorizonBand() HorizonServBand {
	switch HorizonLive {
	case StellarHorizonFlyURL:
		return FlyHorizon
	case StellarHorizonLiveURL:
		return OfficalHorizon
	}
	return UnknownHorizon
}
