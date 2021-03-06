package errors

import (
	"log"
	"fmt"
	"errors"
)

const (
	ErrorSelectFromDBShort     = "code 1"

	ErrorSelectFromDB          = "failed to select from db"
	ErrorUpdateCache           = "failed to update cache"
	ErrorFindAddress           = "address not found"
	ErrorConnectBTCService     = "btc service shutdown"
)

// LogError records error in log.
func LogError(err error, description ...interface{}) {
	log.Println("ERROR -", fmt.Sprint(description...), ":", getErrStr(err))
}

// Fatal records fatal error in log and stops process.
func Fatal(err error, description ...interface{}) {
	panic(fmt.Sprint(fmt.Sprint(description...), " : ", getErrStr(err)))
}

// Combine creates and returns error with description.
func Combine(err error, description ...interface{}) error {
	return errors.New(fmt.Sprintf("%s: %s", fmt.Sprint(description...), getErrStr(err)))
}

func getErrStr(err error) (errStr string) {
	if err != nil { errStr = err.Error() }
	return
}
