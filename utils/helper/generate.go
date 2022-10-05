package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateOrderID(table string, idInTable, userId int) string {
	id := uuid.New()
	return fmt.Sprintf("%s-%d-%d-%s", table, idInTable, userId, id.String())
}

func GenerateTransactionID(table string, transid int) string {
	return fmt.Sprintf("%s-%d", table, transid)
}
