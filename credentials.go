package worldofwarcraft

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

// Credentials are the credentials required to access the Client functionalities. More info at https://develop.battle.net/access/clients
type Credentials struct {
	ClientID string
	Secret   string
}

func (vo Credentials) validate() error {
	return ozzo.ValidateStruct(&vo,
		ozzo.Field(&vo.ClientID, ozzo.Required),
		ozzo.Field(&vo.Secret, ozzo.Required),
	)
}
