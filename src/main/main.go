package main

import "fmt"
import "api"

func main(){
	address, _ := location.GetLocation("39.983424",
		"116.322987", "4zbsyOHMfdK5BDnhnkthr53Z")
	surrounding := location.GetSurrounding("39.983424",
		"116.322987", "4zbsyOHMfdK5BDnhnkthr53Z")
	// 经纬度返回
	fmt.Println(address)
	// 返回经纬度周边标志建筑
	fmt.Println(location.GetNames(surrounding))
	// 返回某一条地点具体信息
	fmt.Println(location.GetInfo(surrounding, 2))
}