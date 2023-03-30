package test

import (
	"testing"

	"github.com/Fibocloud/payment-sdks/pass"
)

var passClient = pass.New("https://ecomstg.pass.mn/openapi/v1/ecom", "c655c1885f024247a78560eaef10f326", "https://webhook.site/05ec7be4-4484-49c1-b1ee-b9669684e335")

func TestCreateOrderAndCancel(t *testing.T) {
	createRes, err := passClient.CreateOrder(100, map[string]string{
		"order": "mongol",
	})
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("createRes: %+v\n", createRes.Ret)

	checkRes, err := passClient.InqueryOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", checkRes.Ret)

	notifyRes, err := passClient.NotifyOrder(createRes.Ret.OrderID, "55555576")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("notifyRes: %+v\n", notifyRes.Ret)

	cancelRes, err := passClient.CancelOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", cancelRes.Ret)

	checkRes, err = passClient.InqueryOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", checkRes.Ret)
}

func TestCreateOrderAndVoid(t *testing.T) {
	createRes, err := passClient.CreateOrder(100, map[string]string{
		"order": "mongol",
	})
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("createRes: %+v\n", createRes.Ret)

	checkRes, err := passClient.InqueryOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", checkRes.Ret)

	notifyRes, err := passClient.NotifyOrder(createRes.Ret.OrderID, "55555576")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("notifyRes: %+v\n", notifyRes.Ret)

	voidRes, err := passClient.VoidOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", voidRes.Ret)

	checkRes, err = passClient.InqueryOrder(createRes.Ret.OrderID)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("checkRes: %+v\n", checkRes.Ret)
}
