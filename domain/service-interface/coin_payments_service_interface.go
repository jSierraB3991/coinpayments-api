package service_interface

import (
	coinpaymentsmodels "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_models"
	coinpaymentsrequest "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_request"
	coinpaymentsresponse "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_response"
)

type CoinPaymentsServiceInterface interface {
	PayChallenge(paymentData coinpaymentsrequest.PaymentData) (*coinpaymentsresponse.PayResponse, error)
	FindData(buy coinpaymentsmodels.Buy) (*coinpaymentsrequest.TransactionInfo, error)
	ValidatePayments(buysPtr *[]coinpaymentsmodels.Buy) ([]uint, []string)
}
