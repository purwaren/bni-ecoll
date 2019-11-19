package dto

type EncryptedRequest struct {
    ClientId    string `json:"client_id"`
    Data        string `json:"data"`
}

type EncryptedResponse struct {
    Status  string `json:"status"`
    Data    string `json:"data"`
}

type CreateBillingRequest struct {
    Type            string `json:"type"`
    ClientId        string `json:"client_id"`
    TrxId           string `json:"trx_id"`
    TrxAmount       string `json:"trx_amount"`
    BillingType     string `json:"billing_type"`
    CustomerName    string `json:"customer_name"`
    CustomerEmail   string `json:"customer_email"`
    CustomerPhone   string `json:"customer_phone"`
    VirtualAccount  string `json:"virtual_account"`
    DateTimeExpired string //in minutes
    Description     string  `json:"description"`
}

type PaymentNotification struct {
    VirtualAccount  string `json:"virtual_account"`
    CustomerName    string `json:"customer_name"`
    TrxId           string `json:"trx_id"`
    TrxAmount       string `json:"trx_amount"`
    PaymentAmount   string `json:"payment_amount"`
    CumulativePay   string `json:"cumulative_payment_amount"`
    PaymentNtb      string `json:"payment_ntb"`
    DateTimePay     string `json:"date_time_payment"`
    DateTimeIso     string `json:"date_time_payment_iso8601"`
}


