package helper

import "fmt"

func GenerateOrderID(table string, someid int) string {
	return fmt.Sprintf("%d_%s", someid, table)
}
