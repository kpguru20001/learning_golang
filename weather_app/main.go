package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"openWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func apiConfiguration() (apiConfigData, error) {
	content, err := os.ReadFile("config.json")

	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData

	err = json.Unmarshal(content, &config)

	if err != nil {
		return apiConfigData{}, err
	}

	return config, nil
}

func main() {
	http.HandleFunc(("/weather"), func(w http.ResponseWriter, r *http.Request) {
		config, err := apiConfiguration()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		city := r.URL.Query().Get("city")
		url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + config.OpenWeatherMapApiKey

		response, err := http.Get(url)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()

		var data weatherData

		err = json.NewDecoder(response.Body).Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8080", nil)
}
