package coinpaymentsrequest

import "github.com/yograterol/coinpayments-go/coinpayments"

type TransactionBodyParams struct {
	coinpayments.APIParams
	TransactionParams
}
