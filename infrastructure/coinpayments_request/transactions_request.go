package coinpaymentsrequest

import "github.com/yograterol/coinpayments-go/coinpayments"

type TransactionInfo struct {
	TimeCreated    int    `json:"time_created"`
	TimeExpires    int    `json:"time_expires"`
	Status         int    `json:"status"`
	StatusText     string `json:"status_text"`
	Type           string `json:"type"`
	Coin           string `json:"coin"`
	Amount         int    `json:"amount"`
	Amountf        string `json:"amountf"`
	Received       int    `json:"received"`
	Receivedf      string `json:"receivedf"`
	RecvConfirms   int    `json:"recv_confirms"`
	Fee            int    `json:"fee"`
	Feef           string `json:"feef"`
	PaymentAddress string `json:"payment_address"`
	SenderIP       string `json:"sender_ip"`
}

type TransactionInfoResponse struct {
	Error  string           `json:"error"`
	Result *TransactionInfo `json:"result"`
}

type TransactionInfoParams struct {
	PaymentId string `url:"txid"`
	Full      uint8  `json:"full"`
}

type TransactionInfoBodyParams struct {
	coinpayments.APIParams
	TransactionInfoParams
}

type Transaction struct {
	Amount         string `json:"amount"`
	Address        string `url:"address"`
	TXNId          string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        uint32 `json:"timeout"`
	StatusUrl      string `json:"status_url"`
	QRCodeUrl      string `json:"qrcode_url"`
	CheckoutUrl    string `json:"checkout_url"`
}

type TransactionResponse struct {
	Error  string       `json:"error"`
	Result *Transaction `json:"result"`
}

type TransactionParams struct {
	Amount     float64 `url:"amount"`
	Currency1  string  `url:"currency1"`
	Currency2  string  `url:"currency2"`
	Address    string  `url:"address"`
	BuyerEmail string  `url:"buyer_email"`
	BuyerName  string  `url:"buyer_name"`
	ItemName   string  `url:"item_name"`
	ItemNumber string  `url:"item_number"`
	Invoice    string  `url:"invoice"`
	Custom     string  `url:"custom"`
	Successurl string  `url:"success_url"`
}
