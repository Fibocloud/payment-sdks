package storepay

import (
	"encoding/json"
	"errors"
	"fmt"
)

type storepay struct {
	appUsername string
	appPassword string
	username    string
	password    string
	endpoint    string
	storeId     string
	callbackUrl string
	ExpireIn    *int64
	loginObject *storepayLoginResponse
}

type Storepay interface {
	LoanCheck(id string) (bool, error)
	Loan(input StorepayLoanInput) (string, error)
	UserPossibleAmount(mobileNumber string) (float64, error)
}

func New(appUsername, appPassword, username, password, endpoint, storeId, callbackUrl string) Storepay {
	return &storepay{
		appUsername: appUsername,
		appPassword: appPassword,
		username:    username,
		password:    password,
		endpoint:    endpoint,
		storeId:     storeId,
		callbackUrl: callbackUrl,
		ExpireIn:    nil,
		loginObject: nil,
	}
}

func (s *storepay) Loan(input StorepayLoanInput) (string, error) {

	request := StorepayLoanRequest{
		StoreId:      s.storeId,
		MobileNumber: input.MobileNumber,
		Description:  input.Description,
		Amount:       fmt.Sprintf("%f", input.Amount),
		CallbackUrl:  s.callbackUrl,
	}

	res, err := s.httpRequest(request, StorepayLoan, "")
	if err != nil {
		return "", err
	}

	var response StorepayLoanResponse
	json.Unmarshal(res, &response)
	if response.Status != "Success" {
		return "", errors.New(response.Status + " - " + response.MsgList[0].Code)
	}
	return response.Value, nil
}

func (s *storepay) LoanCheck(id string) (bool, error) {
	res, err := s.httpRequest(nil, StorepayLoanCheck, id)
	if err != nil {
		return false, err
	}
	var response StorepayCheckResponse
	json.Unmarshal(res, &response)
	if response.Status != "Success" {
		return false, errors.New(response.Status + " - " + response.MsgList[0].Code)
	}
	return response.Value, nil

}

func (s *storepay) UserPossibleAmount(mobileNumber string) (float64, error) {
	request := StorepayUserCheckRequest{
		MobileNumber: mobileNumber,
	}

	res, err := s.httpRequest(request, StorepayUserPossibleAmount, "")
	if err != nil {
		return 0, err
	}

	var response StorepayUserCheckResponse
	json.Unmarshal(res, &response)
	if response.Status != "Success" {
		return 0, errors.New(response.Status + " - " + response.MsgList[0].Code)
	}
	return response.Value, nil
}
