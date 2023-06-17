package worldofwarcraft

const (
	US_Region Region = "US"
	EU_Region Region = "EU"
	KR_Region Region = "KR"
	TW_Region Region = "TW"
	CN_Region Region = "CN"
)

type Region string

func (r Region) String() string {
	return string(r)
}
