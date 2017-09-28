package location

import (
	//"github.com/bitly/go-simplejson"
	"go-simplejson"
)

// Get the latitude and lotitude according to the place
func GetCoordinate(location string, ak string)(string, string){
	body := GeoCoderReLoc(location, ak)
	js, err := simplejson.NewJson(body)
	if err != nil{
		panic(err.Error())
	}
	longitude, _ := js.Get("result").Get("location").Get("lng").MarshalJSON()
	latitude, _ := js.Get("result").Get("location").Get("lat").MarshalJSON()
	return string(longitude), string(latitude)
}