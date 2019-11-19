package dto

type CreateBillingResponse struct {
    Status          string `json:"status"`
    VirtualAccount  string `json:"virtual_account"`
    TrxId           string `json:"trx_id"`
}
