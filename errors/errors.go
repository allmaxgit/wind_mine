package errors

import (
	"log"
	"fmt"
	"errors"
)

// Error periods.
const (
	ErrStart   = iota
	ErrRunning
)

// TODO: Replace Error -> Err
const (
	ErrorFindAddress        = "address not found"
	ErrorConnectBTCService  = "btc service shutdown"
	ErrorConnectRateService = "rate service shutdown"

	// API
	ErrorValidateBTCAddress = "BTC address is not valid"
	ErrorValidateETHAddress = "ETH address is not valid"


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

//// PrintError prints error.
//func PrintFatal(err error, description ...interface{}) {
//	fmt.Println("ERROR -", fmt.Sprint(description...), ":", getErrStr(err))
//	os.Exit(1)
//}

// Combine creates and returns error with description.
func Combine(err error, description ...interface{}) error {
	return errors.New(fmt.Sprintf("%s: %s", fmt.Sprint(description...), getErrStr(err)))
}

func getErrStr(err error) (errStr string) {
	if err != nil { errStr = err.Error() }
	return
}
