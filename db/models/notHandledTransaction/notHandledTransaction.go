package notHandledTransaction

import (
	"WindToken/db"
	"WindToken/db/types"
)

// All returns all not handled transactions.
func All() (*[]types.NotHandledTransaction, bool, error) {
	var result []types.NotHandledTransaction
	err := db.Instance.Model(&result).Select()
	if err != nil {
		if db.IsNotFoundError(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &result, true, nil
}

// FindByHash returns not handled transaction with particular hash.
func FindByHash(hash string) (*types.NotHandledTransaction, bool, error) {
	var result types.NotHandledTransaction
	err := db.Instance.Model(&result).
		Where("hash = ?", hash).
		Select()
	if err != nil {
		if db.IsNotFoundError(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &result, true, nil
}
