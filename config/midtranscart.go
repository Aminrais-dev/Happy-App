package config

import (
	"os"
)

func MidtransServerKey() string {
	// err := godotenv.Load(".env")
	// // if err != nil {
	// // 	fmt.Println("error loading .env file")
	// // }
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_PRODUCT_SERVER_KEY")
	return MIDTRANS_SERVER_KEY
}

func MidtransClientKey() string {
	// err := godotenv.Load(".env")
	// // if err != nil {
	// // 	fmt.Println("error loading .env file")
	// // }
	MIDTRANS_CLIENT_KEY := os.Getenv("MIDTRANS_PRODUCT_CLIENT_KEY")
	return MIDTRANS_CLIENT_KEY
}
