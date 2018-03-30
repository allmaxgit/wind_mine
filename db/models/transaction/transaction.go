package transaction

import (
	"WindToken/db"
	"WindToken/db/types"
)

// FindByHash returns transaction with particular hash.
func FindByHash(hash string) (*types.Transaction, bool, error) {
	var result types.Transaction
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
