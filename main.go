package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	data := map[string]interface{}{
		"title":  "bandung",
		"body":   "bandung adalah ibu kota jawa barat",
		"userId": 1,
	}

	i := 0
	for tick := range ticker.C {
		var _ = tick
		// fmt.Println(tick)
		i++

		min := 1
		max := 100

		water := rand.Intn(max-min+1) + min
		wind := rand.Intn(max-min+1) + min

		var waterStatus string
		var windStatus string

		if water < 5 {
			waterStatus = "aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else if water > 8 {
			waterStatus = "bahaya"
		}

		if wind < 6 {
			windStatus = "aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else if wind > 15 {
			windStatus = "bahaya"
		}

		requestJson, err := json.Marshal(data)

		client := &http.Client{}

		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		waterAndWindData := map[string]interface{}{
			"water": fmt.Sprintf("%v meter", water),
			"wind":  fmt.Sprintf("%v m/s", wind),
		}

		statusJson, _ := json.Marshal(waterAndWindData)
		fmt.Println(bytes.NewBuffer(statusJson))
		fmt.Printf("Status water : %s\n", waterStatus)
		fmt.Printf("Status wind : %s\n", windStatus)

		// fmt.Println("response body :", string(body))

		// if i == 10 {
		// 	fmt.Println("selesai")
		// 	break
		// }
	}

}
