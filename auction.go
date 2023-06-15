package worldofwarcraft

// Auction is an auction of an item from the Auction House
type Auction struct {
	ID   int `json:"id"`
	Item struct {
		ID         int   `json:"id,omitempty"`
		Context    int   `json:"context,omitempty"`
		BonusLists []int `json:"bonus_lists,omitempty"`
		Modifiers  []struct {
			Type  int `json:"type,omitempty"`
			Value int `json:"value,omitempty"`
		} `json:"modifiers,omitempty"`
		PetBreedID   int `json:"pet_breed_id,omitempty"`
		PetLevel     int `json:"pet_level,omitempty"`
		PetQualityID int `json:"pet_quality_id,omitempty"`
		PetSpeciesID int `json:"pet_species_id,omitempty"`
	} `json:"item"`
	Buyout    int    `json:"buyout,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	UnitPrice int    `json:"unit_price,omitempty"`
	TimeLeft  string `json:"time_left,omitempty"`
}
