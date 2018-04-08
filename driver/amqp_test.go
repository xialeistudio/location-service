package driver

import (
	"os"
	"testing"
)

func TestAMap_LocationWithIp(t *testing.T) {
	amap := NewAMap(os.Getenv("AMAP_KEY"))
	location, err := amap.LocationWithIp("121.33.115.164")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(location)
}

func TestAMap_LocationWithAddress(t *testing.T) {
	amap := NewAMap(os.Getenv("AMAP_KEY"))
	location, err := amap.LocationWithAddress("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(location)
}

func TestAMap_LocationWithCoordinate(t *testing.T) {
	amap := NewAMap(os.Getenv("AMAP_KEY"))
	location, err := amap.LocationWithCoordinate("119.378827,26.15519")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(location)
}
