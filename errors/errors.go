package errors

import (
	"log"
	"fmt"
)

const (
	ErrorSelectFromDBShort     = "code 1"

	ErrorSelectFromDB          = "failed to select from db"
	ErrorUpdateCache           = "failed to update cache"
	ErrorFindAddress           = "address not found"
)

// TODO: Write logs to file
func LogError(err error, description ...interface{}) {
	log.Println("ERROR -", fmt.Sprint(description...), ":", getErrStr(err))
}

func Fatal(err error, description ...interface{}) {
	panic(fmt.Sprint(fmt.Sprint(description...), " : ", getErrStr(err)))
}

func getErrStr(err error) (errStr string) {
	if err != nil { errStr = err.Error() }
	return
}
