package util

import (
	"encoding/json"
	"io/ioutil"
)

var Json struct {
		Client ClientType `json:"web"`
	}

type ClientType struct {
	AuthUri string `json:"auth_uri"`
	ClientSecret string `json:"client_secret"`
	TokenUri string `json:"token_uri"`
	ClientEmail string `json:"client_email"`
	ClientX509CertUrl string `json:"client_x509_cert_url"`
	ClientId string `json:"client_id"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	Javascript_origins   []string `json:"Javascript_origins"`
}

func Parse (filePath string ) (e error) {
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	json.Unmarshal(file, &Json)
	return
}
