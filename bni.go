package bni_ecoll

import (
    "encoding/json"
    "github.com/purwaren/bni-ecoll/config"
    "github.com/purwaren/bni-ecoll/dto"
    "github.com/purwaren/bni-ecoll/util"
    "log"
)

type BNI struct {
    config config.Config
    api *API

}

func NewEcolApi(conf config.Config) *BNI {
    a := newApi(conf)
    bni := BNI{config: conf, api: a}
    return &bni
}

func (b *BNI) CreateBilling (req dto.CreateBillingRequest) (dto.CreateBillingResponse, error) {
    encrypted := util.HashData(req, b.config.ClientId, b.config.SecretKey)
    log.Printf("encrypted request = %s", encrypted)
    encReq := dto.EncryptedRequest{
        ClientId: b.config.ClientId,
        Data:     encrypted,
    }

    encResp, err := b.api.postRequest(encReq)

    if err != nil {
        return dto.CreateBillingResponse{}, err
    }
    log.Printf("encResp: %s", encResp)

    var resp dto.CreateBillingResponse
    decByte, err := util.ParseData(encResp.Data, b.config.ClientId, b.config.SecretKey)
    if err != nil {
        log.Printf("Error when decrypt resp: %s", err)
        return dto.CreateBillingResponse{}, err
    }
    log.Printf("decResp: %s", string(decByte))

    err = json.Unmarshal(decByte, &resp)

    if err != nil {
        return dto.CreateBillingResponse{}, err
    }
    log.Printf("resp: %s", resp)
    resp.Status = encResp.Status
    return resp, nil
}

func (b *BNI) PaymentNotification (req dto.EncryptedRequest) (dto.PaymentNotification, error) {
    return dto.PaymentNotification{}, nil
}

