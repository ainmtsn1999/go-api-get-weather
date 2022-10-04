package weathercontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getLoc(name string, API_KEY string) map[string]float64 {

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=" + API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//fmt.Println(string(respData))

	var data map[string]interface{}

	if err = json.Unmarshal(respData, &data); err != nil {
		log.Fatal(err)
	}

	datas := data["coord"].(map[string]interface{})

	var loc = map[string]float64{}

	for key, value := range datas {
		//fmt.Println(key, value.(float64))

		loc[key] = value.(float64)
	}

	return loc
}

func getWeather(loc map[string]float64, API_KEY string) []interface{} {

	//fmt.Println(strconv.FormatFloat(loc["lat"], 'f', -1, 64))
	//fmt.Println(strconv.FormatFloat(loc["lon"], 'f', -1, 64))
	lat := strconv.FormatFloat(loc["lat"], 'f', -1, 64)
	lon := strconv.FormatFloat(loc["lon"], 'f', -1, 64)

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//fmt.Println(string(respData))

	var data map[string]interface{}

	if err = json.Unmarshal(respData, &data); err != nil {
		log.Fatal(err)
	}

	datas := data["weather"].([]interface{})

	//fmt.Println(datas)

	return datas
}

func Index(c *gin.Context) {

	city := c.Query("city")
	API_KEY := "f7e80336d482f800cfbdba35d08e9539"

	//get loc (latitude, longitude)
	loc := getLoc(city, API_KEY)

	//get weather
	data := getWeather(loc, API_KEY)

	c.JSON(http.StatusOK, gin.H{"weather": data})
}
