package model

type Ip struct {
	Id       int32 `gorm:"primaryKey"`
	Ip       string
	Country  string
	City     string
	District string
	Code     int32
}

func (Ip) TableName() string {
	return "location_ip"
}
