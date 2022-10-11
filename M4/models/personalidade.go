package models

type Personalidade struct {
	Id       int    `json:"id"` //denotação para conseguir fazer o decode/encode
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
