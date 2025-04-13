package repositoryinterface

type CoinPamentsRepository interface {
	GetConfigurationByCode(key string) (string, error)
	UpdateBuyByCoinPayment(idCoinpayment string, buyId uint)
	PayByCoinPayment(idCoinPayment string, buyId uint) (string, bool)
}
