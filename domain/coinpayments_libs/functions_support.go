package coinpaymentslibs

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	goquery "github.com/google/go-querystring/query"
)

func GetHMAC(payload, apiPrivateKey string) string {
	mac := hmac.New(sha512.New, []byte(apiPrivateKey))
	mac.Write([]byte(payload))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func GetPayload(payload interface{}) string {
	bodyForm, err := goquery.Values(payload)
	if err != nil {
		log.Fatal(err)
	}
	return bodyForm.Encode()
}
