package coinpaymentsrequest

import coinpaymentslibs "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_libs"

type BuyRequest struct {
	FirstNameFacturation string `json:"first_name"`
	LastNameFacturation  string `json:"last_name"`
	CellphoneFacturation uint64 `json:"cellphone"`
	Country              uint   `json:"country"`
	Email                string `json:"email"`
	SaveData             bool   `json:"save_data"`

	AfiliateCode     string `json:"code_afiliate"`
	ChallengePriceId uint   `json:"challenge_price_id"`
	User             string `json:"-"`
	MetaTrader       string `json:"meta_trader"`
	DiscountCopupon  string `json:"discount_coupon"`
}

type PayData struct {
	Pay        coinpaymentslibs.PaySelect `json:"pay"`
	CryptoCode string                     `json:"crypto_code"`
	BuyId      uint                       `json:"buy_id"`
}

type BuyData struct {
	Email string `json:"email"`
}

type PaymentData struct {
	Name        string
	Description string
	Currency    string
	Price       int64
	Quantity    int64
	BuyId       uint
	UserName    string
	UserId      uint
	UserMail    string
	CryptoCode  string
}
