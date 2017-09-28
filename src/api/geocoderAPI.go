package location

import (
	"net/http"
	"io/ioutil"
)

// Get the service content using url
func GetByUrl(url string)[]byte{
	response,_:=http.Get(url)
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	return body
}
// Call the geo-coder api of BaiDu
// Get location by latitude and longitude
// ak is your authorised key {key}
func GeoCoder(latitude string, longitude string, ak string)[]byte{
	senUrl := "http://api.map.baidu.com/geocoder/v2/?location=" +
		latitude + "," + longitude +"&output=json" +
		"&pois=1&radius=1000&ak=" + ak
	return GetByUrl(senUrl)
}

// Call the geo-coder api of BaiDu
// Get the place's latitude and longitude
// ak is your authorised key {key}
func GeoCoderReLoc(location string, ak string)[]byte{
	senUrl := "http://api.map.baidu.com/geocoder/v2/?address=" +
		location + "&output=json&ak=" + ak
	return GetByUrl(senUrl)
}

