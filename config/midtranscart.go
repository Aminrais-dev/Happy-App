package config

import "os"

func MidtransServerKey() string {
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_PRODUCT_SERVER_KEY")
	return MIDTRANS_SERVER_KEY
}

func MidtransClientKey() string {
	MIDTRANS_CLIENT_KEY := os.Getenv("MIDTRANS_PRODUCT_CLIENT_KEY")
	return MIDTRANS_CLIENT_KEY
}
