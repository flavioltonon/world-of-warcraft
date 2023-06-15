package worldofwarcraft

type Realm struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Region RealmRegion `json:"region"`
	Slug   string      `json:"slug"`
}

type RealmRegion struct {
	Name string `json:"name"`
}
