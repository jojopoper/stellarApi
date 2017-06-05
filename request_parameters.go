package stellarApi

import "github.com/jojopoper/rhttp"

// RequestParameters 请求参数定义
type RequestParameters struct {
	rhttp.CRequestParam
	HttpType       int
	UseTestNetwork bool
}

// NewOrigReqParam 创建一个原生态的Http请求参数
func NewOrigReqParam(isTestNet bool) *RequestParameters {
	ret := new(RequestParameters)
	ret.HttpType = OrigHttp
	ret.UseTestNetwork = isTestNet
	return ret
}

// NewClientReqParam 创建一个Http client请求参数
func NewClientReqParam(isTestNet bool) *RequestParameters {
	ret := new(RequestParameters)
	ret.HttpType = ClientHttp
	ret.Timeout = 60
	ret.UseTestNetwork = isTestNet
	return ret
}

// NewProxyReqParam 创建一个Http 代理请求参数
func NewProxyReqParam(isTestNet bool, ip, port string) *RequestParameters {
	ret := new(RequestParameters)
	ret.HttpType = ProxyHttp
	ret.Timeout = 60
	ret.UseTestNetwork = isTestNet
	ret.ProxyAddr = ip
	ret.ProxyPort = port
	return ret
}

// ResetFormPostData 重置Post类型(Form)、参数和地址
func (ths *RequestParameters) ResetFormPostData(d string) {
	ths.PostDataType = rhttp.PostForm
	ths.PostDatas = d
	ths.Address = ""
}

// ResetJsonPostData 重置Post类型(json)、参数和地址
func (ths *RequestParameters) ResetJsonPostData(d string) {
	ths.PostDataType = rhttp.PostJson
	ths.PostDatas = d
	ths.Address = ""
}
