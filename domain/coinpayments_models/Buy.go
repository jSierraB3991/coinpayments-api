package coinpaymentsmodels

import (
	"time"

	coinpaymentslibs "github.com/jSierraB3991/coinpayments-api/domain/coinpayments_libs"
	"gorm.io/gorm"
)

type Buy struct {
	gorm.Model
	BuyID uint `gorm:"column:id;not null"`

	FirstNameFacturation string                     `gorm:"column:first_name"`
	LastNameFacturation  string                     `gorm:"column:last_name"`
	CellphoneFacturation string                     `gorm:"column:cellphone"`
	CountryId            uint                       `gorm:"column:country_id;not null"`
	Email                string                     `gorm:"column:email"`
	AfiliateCode         string                     `gorm:"column:code_afiliate"`
	ChallengePriceId     uint                       `gorm:"column:challenge_price_id;not null"`
	UserId               uint                       `gorm:"column:user_id;not null"`
	AccountId            uint                       `gorm:"column:account_id;not null"`
	BuyStatus            coinpaymentslibs.PayStatus `gorm:"column:buy_status"`
	CoinBaseID           string                     `gorm:"column:coin_base_id"`
	CoinPaymentID        string                     `gorm:"column:coin_payment_id"`
	StripeId             string                     `gorm:"column:stripe_id"`
	BuyDay               *time.Time                 `gorm:"column:buy_day;"`
	ValueChallengePrice  float64                    `gorm:"column:value_challenge_price;"`
	ValueToPaid          float64                    `gorm:"column:value_to_paid;"`
	SuccessCredentials   bool                       `gorm:"column:success_credentials_data;"`
	DiscountCupon        string                     `gorm:"column:discount_coupon"`
	Step                 coinpaymentslibs.StepEnum  `gorm:"column:step"`
	IsCompetition        bool                       `gorm:"column:is_competion;not null;default:false"`
	IsWithDrawal         bool                       `gorm:"column:is_with_drawal;not null;default:false"`
}
