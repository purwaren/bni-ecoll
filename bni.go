package bni_ecoll

import (
    "encoding/json"
    "errors"
    "github.com/purwaren/bni-ecoll/config"
    "github.com/purwaren/bni-ecoll/dto"
    "github.com/purwaren/bni-ecoll/util"
    "github.com/purwaren/bni-ecoll/util/crypto"
    "github.com/purwaren/bni-ecoll/util/curl"
)

type BNI struct {
    config config.Config
    api API

}

func New(conf config.Config) *BNI {
    a := newApi(conf)
    bni := BNI{config: conf, api: a}
    return &bni
}

func (b *BNI) CreateBilling (req dto.CreateBillingRequest) (*dto.CreateBillingResponse, error) {
    encrypted := util.HashData(req, b.config.ClientId, b.config.SecretKey)

    encReq := dto.EncryptedRequest{
        ClientId: b.config.ClientId,
        Data:     encrypted,
    }

    encResp, err := b.api.postRequest(encReq)
    if err != nil {
        return &dto.CreateBillingResponse{}, err
    }

    var resp dto.CreateBillingResponse
    decByte, err := util.ParseData(encResp.Data, b.config.ClientId, b.config.SecretKey)
    if err != nil {
        return &dto.CreateBillingResponse{}, err
    }
    err = json.Unmarshal(decByte, &resp)
    if err != nil {
        return &dto.CreateBillingResponse{}, err
    }
    return &resp, nil
}

func (b *BNI) PaymentNotification (req dto.EncryptedRequest) (*dto.PaymentNotification, error) {
    return nil, nil
}

