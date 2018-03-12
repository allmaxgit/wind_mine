package buyer

import (
	"WindToken/db"
	"WindToken/db/types"
)

// FindByBTCAddress returns buyer by ethAddr.
func FindByBTCAddress(btcAddr string) (*types.Buyer, bool, error) {
	var result types.Buyer
	err := db.Instance.Model(&result).
		Where("btc_addr = ?", btcAddr).
		Select()
	if err != nil {
		if db.IsNotFoundError(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &result, true, nil
}
