package location

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

// Get location info by longitude and latitude
func GetLocation(latitude string, longitude string, ak string){

	senUrl := "http://api.map.baidu.com/geocoder/v2/?location=" +
		latitude + "," + longitude +"&output=json" +
		"&pois=1&radius=1000&ak=" + ak
	response,_:=http.Get(senUrl)
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}