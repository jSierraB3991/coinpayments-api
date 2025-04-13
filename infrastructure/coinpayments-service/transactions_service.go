package coinpaymentsservice

import (
	"net/http"

	"fmt"

	"github.com/dghubble/sling"
	coinpaymentslibs "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_libs"
	coinpaymentsrequest "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_request"
)

type TransactionService struct {
	sling         *sling.Sling
	ApiPublicKey  string
	ApiPrivateKey string
	Params        coinpaymentsrequest.TransactionBodyParams
}

func NewTransactionService(sling *sling.Sling, apiPublicKey string, privateKey string) *TransactionService {
	transactionService := &TransactionService{
		sling:         sling.Path("api.php"),
		ApiPublicKey:  apiPublicKey,
		ApiPrivateKey: privateKey,
	}
	transactionService.getParams()
	return transactionService
}

func (s *TransactionService) getHMAC() string {
	return coinpaymentslibs.GetHMAC(coinpaymentslibs.GetPayload(s.Params), s.ApiPrivateKey)
}

func (s *TransactionService) NewTransaction(transactionParams *coinpaymentsrequest.TransactionParams) (coinpaymentsrequest.TransactionResponse, *http.Response, error) {
	transactionResponse := new(coinpaymentsrequest.TransactionResponse)
	s.Params.TransactionParams = *transactionParams
	fmt.Println(coinpaymentslibs.GetPayload(s.Params))
	fmt.Println(s.getHMAC())
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(transactionResponse)
	return *transactionResponse, resp, err
}

func (s *TransactionService) getParams() {
	s.Params.Command = "create_transaction"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
