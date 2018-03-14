package errors

import (
	"log"
	"fmt"
	"errors"
)

// TODO: Replace Error -> Err
const (
	ErrorSelectFromDBShort  = "code 1"

	ErrorSelectFromDB       = "failed to select from db"
	ErrorUpdateCache        = "failed to update cache"
	ErrorFindAddress        = "address not found"
	ErrorConnectBTCService  = "btc service shutdown"
	ErrorConnectRateService = "rate service shutdown"

	ErrorTXTimedOut         = "transaction timed out"
	ErrorReceiptStatus      = "receipt.Status is 0"
	ErrorUpdateGasLimit     = "failed to update gas limit"
	ErrorOwner              = "permission denied, please check ownerAddress and ownerPrivateKey"
	ErrorICONotStarted      = "ico not started"
	ErrorICOFinished        = "ico finished"

	UnknownError            = "unknown error"
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
