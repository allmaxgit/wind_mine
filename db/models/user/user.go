package user

import (
	"WindToken/db"
	"WindToken/db/types"
)

// FindUserByETH returns user by ethAddr
func FindUserByETH(ethAddr string) (*types.User, bool, error) {
	var result types.User
	err := db.Instance.Model(&result).
		Where("eth_addr = ?", ethAddr).
		Select()
	if err != nil {
		if db.IsNotFoundError(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &result, true, nil
}
