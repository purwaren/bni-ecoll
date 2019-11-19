package bni_ecoll

import (
    "bytes"
    "encoding/json"
    "github.com/hashicorp/go-retryablehttp"
    "github.com/purwaren/bni-ecoll/config"
    "github.com/purwaren/bni-ecoll/dto"
    "io/ioutil"
    "log"
    "net/http"
)

type API struct {
    conf config.Config
    client *retryablehttp.Client
}

func newApi(config config.Config) *API {
    httpClient := retryablehttp.NewClient()
    return &API{client: httpClient, conf:config}
}

func (a *API) postRequest(req dto.EncryptedRequest) (dto.EncryptedResponse, error) {
    bodyReq, _ := json.Marshal(req)
    request, err := retryablehttp.NewRequest(http.MethodPost, a.conf.URL, bytes.NewBuffer(bodyReq))

    if err != nil {
        log.Printf("Error NewRequest")
        return dto.EncryptedResponse{}, err
    }

    request.Header.Set("Content-Type", "application/json")

    resp, err := a.client.Do(request)
    if err != nil {
        return dto.EncryptedResponse{}, err
    }
    defer resp.Body.Close()

    bodyRespBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return dto.EncryptedResponse{}, err
    }

    resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRespBytes))

    var jsonResp dto.EncryptedResponse
    err = json.NewDecoder(resp.Body).Decode(&jsonResp)

    if err != nil {
        return dto.EncryptedResponse{}, err
    }

    return jsonResp, nil
}