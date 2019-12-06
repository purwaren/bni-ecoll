package dto

import "database/sql"

type CreateBillingResponse struct {
    Status          string `json:"status"`
    VirtualAccount  string `json:"virtual_account"`
    TrxId           string `json:"trx_id"`
    DateTimeExpired sql.NullTime `json:"datetime_expired"`
}
