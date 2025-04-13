package coinpaymentsservice

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
	coinpaymentslibs "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_libs"
	coinpaymentsmodels "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_models"
	repositoryinterface "github.com/jSierraB3991/coinpayments-api/domain/repository-interface"
	coinpaymentsrequest "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_request"
	coinpaymentsresponse "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_response"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

type CoinPaymentService struct {
	publicKey  string
	privateKey string
	repo       repositoryinterface.CoinPamentsRepository
	APIBase    string
}

func NewCoinPayment(publicKey, privateKey string, repo repositoryinterface.CoinPamentsRepository) *CoinPaymentService {
	return &CoinPaymentService{
		publicKey:  publicKey,
		privateKey: privateKey,
		repo:       repo,
		APIBase:    "https://www.coinpayments.net/",
	}
}

func (coinPaymentService CoinPaymentService) PayChallenge(paymentData coinpaymentsrequest.PaymentData) (*coinpaymentsresponse.PayResponse, error) {

	successPage, err := coinPaymentService.repo.GetConfigurationByCode(coinpaymentslibs.VARIABLE_SUCCESS_PAGE)
	if err != nil {
		return nil, err
	}

	transactionParams := coinpaymentsrequest.TransactionParams{
		Amount:     float64(paymentData.Price),
		Currency1:  strings.ToUpper(paymentData.Currency),
		Currency2:  paymentData.CryptoCode,
		BuyerEmail: paymentData.UserMail,
		BuyerName:  paymentData.UserName,
		ItemName:   paymentData.Name,
		ItemNumber: jsierralibs.GetStringToInt64(paymentData.Quantity),
		Successurl: successPage,
	}

	clientHttp := &http.Client{}
	baseClient := sling.New().Client(clientHttp).Base(coinPaymentService.APIBase)
	transactionService := NewTransactionService(baseClient.New(), coinPaymentService.publicKey, coinPaymentService.privateKey)
	result, httpResponse, err := transactionService.NewTransaction(&transactionParams)
	if err != nil || httpResponse.StatusCode != 200 {
		return nil, errors.New(result.Error)
	}
	coinPaymentService.repo.UpdateBuyByCoinPayment(result.Result.TXNId, paymentData.BuyId)
	return &coinpaymentsresponse.PayResponse{
		Url: result.Result.CheckoutUrl,
	}, nil
}

func (service *CoinPaymentService) FindData(buy coinpaymentsmodels.Buy) (*coinpaymentsrequest.TransactionInfo, error) {
	clientHttp := &http.Client{}
	baseClient := sling.New().Client(clientHttp).Base(service.APIBase)
	transactionService := NewTransactionInfoService(baseClient.New(), service.publicKey, service.privateKey)
	result, httpResult, err := transactionService.NewTransactionInfo(&coinpaymentsrequest.TransactionInfoParams{
		PaymentId: buy.CoinPaymentID,
		Full:      1,
	})
	if err != nil || httpResult.StatusCode != 200 {
		return nil, errors.New(result.Error)
	}

	return result.Result, nil
}

func (service *CoinPaymentService) ValidatePayments(buysPtr *[]coinpaymentsmodels.Buy) ([]uint, []string) {
	var slice []uint
	var accounts []string
	buys := *buysPtr
	for i, buy := range *buysPtr {
		if buy.BuyStatus == coinpaymentslibs.PAY_PENDING && buy.CoinPaymentID != "" {
			result, err := service.FindData(buy)
			if err == nil {
				if result.StatusText == "Complete" {
					account, sendMail := service.repo.PayByCoinPayment(buy.CoinPaymentID, buy.BuyID)
					if sendMail {
						slice = append(slice, buy.BuyID)
					}
					accounts = append(accounts, account)
					buys[i].BuyStatus = coinpaymentslibs.PAY_CONFIRM
				}
			}
		}
	}
	return slice, accounts
}
