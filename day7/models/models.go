package models

type Deal struct {
	Hand string `json:"hand"`
	Bid  int    `json:"bid"`
	Type int
	Rank int
}
