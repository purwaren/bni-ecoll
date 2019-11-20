package bni_ecoll

import (
    "github.com/purwaren/bni-ecoll/config"
    "github.com/purwaren/bni-ecoll/dto"
    "log"
    "testing"
)

func TestBNI_CreateBilling(t *testing.T) {
    givenConfig := config.Config{
        URL: "https://apibeta.bni-ecollection.com:10443",
        ClientId: "00208",
        PrefixId: "988",
        SecretKey: "a5536b63a9f44eda2ce96f87e859d150",
    }

    bni := New(givenConfig)

    req := dto.NewBillingRequest()
    req.TrxId = "20191119204304304008"
    req.TrxAmount = "100000"
    req.CustomerName = "Purwa Ren"
    req.CustomerPhone = "039239993435"
    req.CustomerEmail = "and.thau@gmail.com"
    req.DateTimeExpired = "2019-11-19T21:00:00+07:00"

    log.Printf("req: %s", req)
    res, err := bni.CreateBilling(req)

    if err != nil {
        log.Println(err)
    } else {
        log.Println(res)
    }

}
