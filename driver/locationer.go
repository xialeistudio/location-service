package driver

type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Code     int32  `json:"code"`
}

type Locationer interface {
	// ip定位
	LocationWithIp(ip string) (*Location, error)
	// 地址解析
	LocationWithAddress(address string) (*Location, error)
	// 逆地址解析
	LocationWithCoordinate(points ...string) ([]Location, error)
}
