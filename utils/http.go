package utils

type API struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

const (
	TimeFormatYYYYMMDDHHMMSS = "20060102150405"
	TimeFormatYYYYMMDD       = "20060102"
	HttpContent              = "application/json"
)

// func HttpRequestMongolchat(body interface{}, api helper.API) (res *http.Response, err error) {
// 	var requestByte []byte
// 	var requestBody *bytes.Reader
// 	if body == nil {
// 		requestBody = bytes.NewReader(nil)
// 	} else {
// 		requestByte, _ = json.Marshal(body)
// 		requestBody = bytes.NewReader(requestByte)
// 	}

// 	req, _ := http.NewRequest(api.Method, viper.GetString("mongolchat.endpoint")+api.Url, requestBody)

// 	req.Header.Add("Content-Type", helper.HttpContent)
// 	req.Header.Add("api-key", viper.GetString("mongolchat.apikey"))
// 	req.Header.Add("Authorization", "WorkerKey "+viper.GetString("mongolchat.workerkey"))

// 	res, err = http.DefaultClient.Do(req)
// 	return
// }

// func HttpRequestSocialpay(body interface{}, api helper.API) (res *http.Response, err error) {
// 	var requestByte []byte
// 	var requestBody *bytes.Reader
// 	if body == nil {
// 		requestBody = bytes.NewReader(nil)
// 	} else {
// 		requestByte, _ = json.Marshal(body)
// 		requestBody = bytes.NewReader(requestByte)
// 	}

// 	req, _ := http.NewRequest(api.Method, viper.GetString("socialpay.endpoint")+api.Url, requestBody)
// 	req.Header.Add("Content-Type", helper.HttpContent)
// 	res, err = http.DefaultClient.Do(req)
// 	return
// }

// func HttpRequestQpay(body interface{}, api helper.API, urlExt string) (res *http.Response, err error) {
// 	authObj, authErr := shared.AuthQPayV2()
// 	if authErr != nil {
// 		fmt.Println(authErr.Error())
// 		err = authErr
// 		return
// 	}
// 	var requestByte []byte
// 	var requestBody *bytes.Reader
// 	if body == nil {
// 		requestBody = bytes.NewReader(nil)
// 	} else {
// 		requestByte, _ = json.Marshal(body)
// 		requestBody = bytes.NewReader(requestByte)
// 	}

// 	req, _ := http.NewRequest(api.Method, viper.GetString("qpay.endpoint")+api.Url+urlExt, requestBody)

// 	req.Header.Add("Content-Type", helper.HttpContent)
// 	req.Header.Add("Authorization", "Bearer "+authObj.AccessToken)

// 	res, err = http.DefaultClient.Do(req)
// 	return
// }

// func HttpRequestUPoint(body interface{}, api helper.API) (res *http.Response, err error) {
// 	var requestByte []byte
// 	var requestBody *bytes.Reader
// 	if body == nil {
// 		requestBody = bytes.NewReader(nil)
// 	} else {
// 		requestByte, _ = json.Marshal(body)
// 		requestBody = bytes.NewReader(requestByte)
// 	}

// 	req, _ := http.NewRequest(api.Method, viper.GetString("upoint.endpoint")+api.Url, requestBody)

// 	req.Header.Add("Content-Type", helper.HttpContent)
// 	req.Header.Add("Authorization", "Token "+viper.GetString("upoint.token"))

// 	res, err = http.DefaultClient.Do(req)
// 	return
// }

// func HttpRequestEbarimt(body interface{}, api helper.API, ext string) (res *http.Response, err error) {
// 	var requestByte []byte
// 	var requestBody *bytes.Reader
// 	if body == nil {
// 		requestBody = bytes.NewReader(nil)
// 	} else {
// 		requestByte, _ = json.Marshal(body)
// 		requestBody = bytes.NewReader(requestByte)
// 	}
// 	req, err := http.NewRequest(api.Method, viper.GetString("ebarimt.endpoint")+api.Url+ext, requestBody)
// 	if err != nil {
// 		err = errors.New("НӨАТ хүсэтийг боловсруулж чадсангүй")
// 		return
// 	}
// 	req.Header.Add("Content-Type", helper.HttpContent)
// 	res, err = http.DefaultClient.Do(req)
// 	return
// }
