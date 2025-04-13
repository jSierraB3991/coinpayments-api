package coinpaymentsservice

import (
	"net/http"

	"fmt"

	"github.com/dghubble/sling"
	coinpaymentslibs "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_libs"
	coinpaymentsrequest "github.com/jSierraB3991/coinpayments-api/infrastructure/coinpayments_request"
)

type TransactionInfoService struct {
	sling         *sling.Sling
	ApiPublicKey  string
	ApiPrivateKey string
	Params        coinpaymentsrequest.TransactionInfoBodyParams
}

func NewTransactionInfoService(sling *sling.Sling, apiPublicKey, privateKey string) *TransactionInfoService {
	transactionService := &TransactionInfoService{
		sling:         sling.Path("api.php"),
		ApiPublicKey:  apiPublicKey,
		ApiPrivateKey: privateKey,
	}
	transactionService.setParams()
	return transactionService
}

func (s *TransactionInfoService) getHMAC() string {
	return coinpaymentslibs.GetHMAC(coinpaymentslibs.GetPayload(s.Params), s.ApiPrivateKey)
}

func (s *TransactionInfoService) NewTransactionInfo(transactionParams *coinpaymentsrequest.TransactionInfoParams) (coinpaymentsrequest.TransactionInfoResponse, *http.Response, error) {
	transactionResponse := new(coinpaymentsrequest.TransactionInfoResponse)
	s.Params.TransactionInfoParams = *transactionParams
	fmt.Println(coinpaymentslibs.GetPayload(s.Params))
	fmt.Println(s.getHMAC())
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(transactionResponse)
	return *transactionResponse, resp, err
}

func (s *TransactionInfoService) setParams() {
	s.Params.Command = "get_tx_info"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
