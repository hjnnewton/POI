package location

// Search poi according to the coordinates
// The searching range is a circle whose radius is
func SearchPOI(){
	lat, lng := GetCoordinate()
	GetByUrl("http://api.map.baidu.com/place/v2/search?" +
		"query=%E9%85%92%E5%BA%97&page_size=10&page_num=0&scope=1" +
		"&location=39.915,116.404&radius=2000&output=json&ak=4zbsyOHMfdK5BDnhnkthr53Z")

}