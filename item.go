package worldofwarcraft

type Item struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SellPrice int    `json:"sell_price"`
}
