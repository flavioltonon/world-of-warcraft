package worldofwarcraft

const (
	en_US_Locale Locale = "en_US"
	pt_BR_Locale Locale = "pt_BR"
)

type Locale string

func (r Locale) String() string {
	return string(r)
}
