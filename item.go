package worldofwarcraft

type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SellPrice int    `json:"sell_price"`
}
