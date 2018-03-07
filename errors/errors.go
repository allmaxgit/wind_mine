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

func LogError(err error, description ...interface{}) {
	log.Println("ERROR -", fmt.Sprint(description...), ":", getErrStr(err))
}

func Fatal(err error, description ...interface{}) {
	panic(fmt.Sprint(fmt.Sprint(description...), " : ", getErrStr(err)))
}

func Combine(err error, description ...interface{}) error {
	return errors.New(fmt.Sprintf("%s: %s", fmt.Sprint(description...), getErrStr(err)))
}

func getErrStr(err error) (errStr string) {
	if err != nil { errStr = err.Error() }
	return
}
