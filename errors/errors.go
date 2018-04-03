package errors

import (
	"errors"
	"fmt"
	"log"
)

// Error codes.
const (
	ErrInsertToDBCode = iota + 1
	ErrSelectFromDBCode
)

// Error messages.
const (
	ErrSelectFromDB       = "failed to select from DB"
	ErrInsertToDB         = "failed to insert to DB"
	ErrFindAddress        = "address not found"
	ErrConnectBTCService  = "btc service shutdown"
	ErrConnectRateService = "rate service shutdown"

	// API
	ErrValidateBTCAddress = "BTC address is not valid"
	ErrValidateETHAddress = "ETH address is not valid"

	ErrTXTimedOut      = "transaction timed out"
	ErrReceiptStatus   = "receipt.Status is 0"
	ErrUpdateGasLimit  = "failed to update gas limit"
	ErrOwner           = "permission denied, please check ownerAddress and ownerPrivateKey"
	ErrICONotStarted   = "ico not started"
	ErrICOFinished     = "ico finished"
	ErrDataRetrieve    = "failed to retrieve data from contract"
	ErrCreateRawTX     = "failed to create raw transaction"
	ErrSendRawTx       = "failed to send raw transaction"
	ErrSignIn          = "failed to sign one or more transaction inputs"
	ErrSignTx          = "failed to sign raw transaction"
	ErrPrivDecode      = "failed to decode private key to bitcoin WIF"
	ErrPayToAddrScript = "failed to create PayToPubKey script"
	ErrNodeInfo        = "failed to get btc node info"

	ErrUnknown = "unknown error"
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

// ErrCode
func CombineErrCode(errCode int) error {
	return errors.New(fmt.Sprintf("code %d", errCode))
}

func getErrStr(err error) (errStr string) {
	if err != nil {
		errStr = err.Error()
	}
	return
}
