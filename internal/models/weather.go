package models

type WeatherInfoRaw struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
}

type WeatherInfo struct {
	Weather     string `json:"weather"`
	Description string `json:"description"`
	Temp        int    `json:"temp"`
	FeelsLike   int    `json:"feels_like"`
}
