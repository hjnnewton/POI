package location

import (
	"github.com/bitly/go-simplejson"
	"fmt"
	"encoding/json"

)

// According to the latitude and longitude, return the surrounding places information
func GetSurrounding(latitude string, longitude string, ak string)[]interface {}{
	body:=GeoCoder(latitude, longitude, ak)
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}
	surroundings, _ := json.Marshal(js.Get("result"))
	var mapTmp []interface{}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(surroundings)), &dat); err == nil {
		// Transfer json str into map type
		mapTmp = dat["pois"].([]interface {})
		fmt.Println(GetNames(mapTmp))
		fmt.Println(GetInfo(mapTmp, 2))
		return mapTmp
	}
	// Return value is a list of maps of information of surrounding places
	return mapTmp
}

// Return the names of the surrounding places
func GetNames(result []interface{})[]string{
	var names []string
	for _, poi := range result {
		//fmt.Println(poi.(map[string]interface{})["name"])
		names = append(names, poi.(map[string]interface{})["name"].(string))
	}
	return names
}

// Return the information of specific place, index is the ith place returned
func GetInfo(result []interface{}, index int)map[string]interface{}{
	var info map[string]interface{}
	if index < len(result){
		info = result[index].(map[string]interface{})
		// point := info["point"]
		return info
	}
	return info

}