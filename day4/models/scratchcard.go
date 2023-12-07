package models

type ScratchCard struct {
	Id   int   `json:"id"`
	Win  []int `json:"win"`
	Nums []int `json:"nums"`
}

func (c *ScratchCard) IsWinningNumber(n int) bool {
	for _, w := range c.Win {
		if n == w {
			return true
		}
	}

	return false
}
