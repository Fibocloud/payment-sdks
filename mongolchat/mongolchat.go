package mongolchat

import (
	"encoding/json"
)

type mongolchat struct {
	endpoint  string
	apikey    string
	workerkey string
	appsecret string
	branchno  string
}
type MongolChat interface {
	GenerateQR(input MchatOnlineQrGenerateRequest) (MchatOnlineQrGenerateResponse, error)
	CheckQR(qr string) (MchatOnlineQrCheckResponse, error)
}

func New(endpoint, apikey, workerkey, appsecret, branchno string) MongolChat {
	return mongolchat{
		endpoint:  endpoint,
		apikey:    apikey,
		workerkey: workerkey,
		appsecret: appsecret,
		branchno:  branchno,
	}
}

func (s mongolchat) GenerateQR(input MchatOnlineQrGenerateRequest) (response MchatOnlineQrGenerateResponse, err error) {
	res, err := s.httpRequestMongolChat(input, MchatOnlineQrGenerate)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s mongolchat) CheckQR(qr string) (response MchatOnlineQrCheckResponse, err error) {
	request := make(map[string]interface{})
	request["qr"] = qr
	res, err := s.httpRequestMongolChat(request, MchatOnlineQrcheck)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}
