package repositoryinterface

type CoinPamentsRepository interface {
	GetConfigurationByCode(key string) (string, error)
	UpdateBuyByCoinPayment(idCoinpayment string, buyId uint) error
	PayByCoinPayment(idCoinPayment string, buyId uint) (string, bool)
}
