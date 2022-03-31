package models

type Tecnologias struct {
	Id        int64  `json:"id"`
	Nome      string `json:"nome"`
	Criador   string `json:"criador"`
	Ano       string `json:"ano"`
	Objetivos string `json:"objetivos"`
}

var Tecnologia []Tecnologias
