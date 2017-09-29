package location

import (
	"encoding/json"
	"go-simplejson"
	"fmt"
)

// Search poi according to the coordinates
// The searching range is a circle whose radius is
func SearchPOI(location string, ak string, interest string) string{
	lng, lat := GetCoordinate(location, ak)

	body := GetByUrl("http://api.map.baidu.com/place/v2/search?" +
		"query=" + interest + "&page_size=20&page_num=0&scope=1" +
		"&location=" + lat + "," + lng +"&radius=2000&output=json&ak=" + ak)
	js, _ := simplejson.NewJson(body)
	pois, _ := json.Marshal(js.Get("results"))
	status,_ := json.Marshal(js.Get("status"))
	if string(status) == "302"{
		fmt.Println("超出日调用限额")
		return "超出限额，查询失败"
	}

	fmt.Printf("%s周边的%s有：\n %s", location, interest, string(pois))
	return string(pois)
}
