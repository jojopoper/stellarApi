package stellarApi

import (
	"fmt"

	"github.com/jojopoper/rhttp"
)

// HttpResponse http get response
type HttpResponse struct {
	rhttp.CHttp
	isTestNet bool
}

// GetResponse 获取Http get结果
func (ths *HttpResponse) GetResponse(p *RequestParameters) (interface{}, error) {
	if p == nil {
		return nil, fmt.Errorf("Input parameter is null")
	}
	if p.Timeout < 10 {
		p.Timeout = 60
	}
	switch p.HttpType {
	case OrigHttp:
		return ths.Get(p.Address, rhttp.ReturnCustomType)
	case ClientHttp:
		ths.SetClient(ths.GetClient(p.Timeout))
		return ths.ClientGet(p.Address, rhttp.ReturnCustomType)
	case ProxyHttp:
		proxyClient, err := ths.GetProxyClient(p.Timeout, p.ProxyAddr, p.ProxyPort)
		if err != nil {
			return nil, err
		}
		ths.SetClient(proxyClient)
		return ths.ClientGet(p.Address, rhttp.ReturnCustomType)
	case ConnectionHttp:
		conn, err := ths.GetClientConn(p.Address, p.Timeout, false)
		if err != nil {
			return nil, err
		}
		ths.SetClientConn(conn)
		if p.ConnectionHeader == nil {
			p.OrigConnectHeader()
		}
		err = ths.ClientConnGet(p.Address, p.ConnectionHeader)
		if err != nil {
			return nil, err
		}
		return ths.ClientConnResponse(rhttp.ReturnCustomType)
	}
	return nil, fmt.Errorf("Can not read HttpType from parameter")
}

// PostResponse 获取Http post结果
func (ths *HttpResponse) PostResponse(p *RequestParameters) (interface{}, error) {
	if p == nil {
		return nil, fmt.Errorf("Input parameter is null")
	}
	if p.Timeout < 10 {
		p.Timeout = 60
	}
	switch p.HttpType {
	case OrigHttp:
		if p.PostDataType == rhttp.PostJson {
			return ths.PostJSON(p.Address, rhttp.ReturnCustomType, p.PostDatas)
		}
		return ths.PostForm(p.Address, rhttp.ReturnCustomType, p.PostDatas)
	case ClientHttp:
		ths.SetClient(ths.GetClient(p.Timeout))
		if p.PostDataType == rhttp.PostJson {
			return ths.ClientPostJSON(p.Address, rhttp.ReturnCustomType, p.PostDatas)
		}
		return ths.ClientPostForm(p.Address, rhttp.ReturnCustomType, p.PostDatas)
	case ProxyHttp:
		proxyClient, err := ths.GetProxyClient(p.Timeout, p.ProxyAddr, p.ProxyPort)
		if err != nil {
			return nil, err
		}
		ths.SetClient(proxyClient)
		if p.PostDataType == rhttp.PostJson {
			return ths.ClientPostJSON(p.Address, rhttp.ReturnCustomType, p.PostDatas)
		}
		return ths.ClientPostForm(p.Address, rhttp.ReturnCustomType, p.PostDatas)
	case ConnectionHttp:
		conn, err := ths.GetClientConn(p.Address, p.Timeout, false)
		if err != nil {
			return nil, err
		}
		ths.SetClientConn(conn)
		if p.ConnectionHeader == nil {
			p.OrigConnectHeader()
		}
		if p.PostDataType == rhttp.PostJson {
			err = ths.ClientConnPostJSON(p.Address, p.PostDatas, p.ConnectionHeader)
		} else {
			err = ths.ClientConnPostForm(p.Address, p.PostDatas, p.ConnectionHeader)
		}
		if err != nil {
			return nil, err
		}
		return ths.ClientConnResponse(rhttp.ReturnCustomType)
	}
	return nil, fmt.Errorf("Can not read HttpType from parameter")
}
