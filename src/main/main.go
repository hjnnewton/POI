package main

import "fmt"
import "api"

func main(){
	// 纬度
	lat := "39.983424"
	// 经度
	lng := "116.322987"
	// 密钥
	ak := "4zbsyOHMfdK5BDnhnkthr53Z"
	// 根据经纬度返回地址
	address, _ := location.GetLocation(lat,
		lng, ak)
	fmt.Println(address)

	// 返回经纬度周边标志建筑
	surrounding := location.GetSurrounding(lat,
		lng, ak)
	fmt.Println(location.GetNames(surrounding))
	// 返回周边某一条地点具体信息
	fmt.Println("第二个地点的信息：")
	fmt.Println(location.GetInfo(surrounding, 2))

	// 根据输入地点返回经纬度
	loc := "北京望京SOHO"
	lng2, lat2 := location.GetCoordinate(loc, ak)
	fmt.Println(loc + "的经纬度是" + "longitude:" +lng2 + "latitude" + lat2)

	// 返回地点周边的酒店
	a := location.SearchPOI(loc, ak, "酒店")
	fmt.Println(a)
}