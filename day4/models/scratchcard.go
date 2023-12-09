package models

type ScratchCard struct {
	Id   int   `json:"id"`
	Win  []int `json:"win"`
	Nums []int `json:"nums"`
}

func (c *ScratchCard) isWinningNumber(n int) bool {
	for _, w := range c.Win {
		if n == w {
			return true
		}
	}

	return false
}

func (c *ScratchCard) SumWinningNumbers() int {
	winningNumbers := 0
	for _, n := range c.Nums {
		if c.isWinningNumber(n) {
			winningNumbers++
		}
	}
	return winningNumbers
}
