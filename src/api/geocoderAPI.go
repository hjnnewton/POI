package location

import (
	"net/http"
	"io/ioutil"
)

// Call the geo-coder api of BaiDu
func GeoCoder(latitude string, longitude string, ak string)[]byte{
	senUrl := "http://api.map.baidu.com/geocoder/v2/?location=" +
		latitude + "," + longitude +"&output=json" +
		"&pois=1&radius=1000&ak=" + ak
	response,_:=http.Get(senUrl)
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	return body
}