package models

type Pokemon []struct {
	ID   int      `json:"id"`
	Name Name     `json:"name"`
	Type []string `json:"type"`
	Base Base     `json:"base"`
}
type Name struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}
type Base struct {
	Hp        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}
