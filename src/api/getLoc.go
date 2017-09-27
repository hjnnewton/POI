package location

import (
	"github.com/bitly/go-simplejson"
)

// Get location info by longitude and latitude
// The return values are the formatted address and address components

func GetLocation(latitude string, longitude string, ak string)(string, string){
	body:=GeoCoder(latitude, longitude, ak)
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}
	address, _ := js.Get("result").Get("formatted_address").String()
	addressInJson, _ := js.Get("result").Get("addressComponent").MarshalJSON()
	return address, string(addressInJson)
}