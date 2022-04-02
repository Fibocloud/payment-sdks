package upoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

// SocialPay
var (
	// Лояалти системд бүртгэлтэй хэрэглэгчийг картын дугаар болон утасны дугаараар шалган  үлдэгдлийн мэдээлэл илгээх функц.
	UpointCheckUserInfo = utils.API{
		Url:    "/transaction/thirdparty/check_info/",
		Method: http.MethodPost,
	}

	// Борлуулалтын гүйлгээ бүртгэх API. Хэрэглэгчийн утасны дугаар болон картын дугаараар  хэрэглэгчид оноогоо зарцуулах, зарцуулсанаас үлдсэн дүнд оноо үүсгэн, үүссэн онооны  мэдээллийг буцаана.
	UpointProcessTransaction = utils.API{
		Url:    "/transaction/thirdparty/process_transaction/",
		Method: http.MethodPost,
	}

	// Хэрэглэгч өөрийн худалдан авсан бараагаа буцаах үед Лояалти системд үүсгэгдсэн оноо ба  зарцуулагдсан оноог буцаах шаардлагатай. Борлуулалтын буцаалт хийх үед хэрэглэгчийн Лояалти  системийн онооны үлдэгдэл нь тухайн худалдан авалтаас тухайн хэрэглэгчид үүсгэгдсэн онооноос  бага болсон тохиолдолд боломжит хасагдах үлдэгдэл хүртэл хасах ба зөрүү үүсэх оноог хэрэглэгч  нь зөрүү болгон төлөх шаардлагатай.
	UpointReturnTransaction = utils.API{
		Url:    "/transaction/thirdparty/return_transaction/",
		Method: http.MethodPost,
	}

	// Лояалти системд бүртгэгдсэн гүйлгээний дугаараар шалган өмнө нь хийгдсэн гүйлгээний  мэдээлэл шалгах функц.
	UpointCheckTransaction = utils.API{
		Url:    "/transaction/thirdparty/check_transaction/",
		Method: http.MethodPost,
	}
	UpointCancelTransaction = utils.API{
		Url:    "/transaction/thirdparty/cancel_transaction/",
		Method: http.MethodPost,
	}
	UpointProduct = utils.API{
		Url:    "/product/product/",
		Method: http.MethodGet,
	}
	// QR code үүсгэх API.
	UpointQr = utils.API{
		Url:    "/transaction/thirdparty/get_qr/",
		Method: http.MethodPost,
	}

	// Уг API нь тухайн QR-г уншуулсан эсэхийг шалгана.
	UpointCheckQr = utils.API{
		Url:    "/transaction/thirdparty/check_qr/",
		Method: http.MethodPost,
	}
	// Уг API нь тухайн QR-г уншуулсан хэрэглэгчийн мэдээллийг авах.
	UpointCheckQrInfo = utils.API{
		Url:    "/transaction/thirdparty/check_info_qr/",
		Method: http.MethodPost,
	}
	// Уг API нь тухайн QR-г уншуулсан хэрэглэгчийн мэдээллээр борлуулалтын гүйлгээ хийх.
	UpointTransactionQr = utils.API{
		Url:    "/transaction/thirdparty/process_transaction_qr/",
		Method: http.MethodPost,
	}
)

func (u *upoint) httpRequestUPoint(body interface{}, api utils.API) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, u.endpoint+api.Url, requestBody)

	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Token "+u.token)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(string(response))
		fmt.Printf("err here")
		return
	}
	response, _ = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return
}
