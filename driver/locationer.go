package driver

import "errors"

var (
	ErrInvalidResponse  = errors.New("-1: 解析响应失败")
	ErrAddressNotExists = errors.New("-2: 地址不存在")
)

type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Code     int    `json:"code"`
}

type Locationer interface {
	// ip定位
	LocationWithIp(ip string) (*Location, error)
	// 地址解析
	LocationWithAddress(address string) (*Location, error)
	// 逆地址解析
	LocationWithCoordinate(point string) (*Location, error)
}
