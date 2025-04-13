package coinpaymentslibs

type PaySelect string

const (
	STRIPE       PaySelect = "STRIPE"
	COIN_BASE    PaySelect = "COIN_BASE"
	COIN_PAYMENT PaySelect = "COIN_PAYMENT"
)

type PayStatus string

const (
	PAY_CREATED PayStatus = "PAY_CREATED"
	PAY_PENDING PayStatus = "PAY_PENDING"
	PAY_CONFIRM PayStatus = "PAY_CONFIRM"
	PAY_FAILED  PayStatus = "PAY_FAILED"
)
