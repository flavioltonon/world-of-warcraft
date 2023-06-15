package worldofwarcraft

import "net/http"

type service struct {
	httpClient *http.Client
	options    ClientOptions
	token      *Token
}
