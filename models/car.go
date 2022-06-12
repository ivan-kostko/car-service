package models

type Car struct {
	Model           string `json:"model"`
	Engine          string `json:"engine"`
	Infotainment    string `json:"infotainment"`
	Interrior       string `json:"interrior"`
	CurrentLocation string `json:"location"`
}

type CarEntity struct {
	Entity
	Car
}
