# 定位服务
面向接口设计的定位服务
![Build Status](https://www.travis-ci.org/xialeistudio/location-service.svg?branch=master)


## 使用方法

1. `go get github.com/xialeistudio/location-service`

2. 
```go
amap := driver.NewAMap("您申请的高德地图Key")
location, err := amap.LocationWithAddress("广东省广州市天河区")
log.Print(location, err)
```

## 功能列表

+ LocationWithIp IP定位
+ LocationWithAddress 地址解析
+ LocationWithCoordinate 逆地址解析

## 支持

+ 高德地图

## TODO

+ [ ] 百度地图