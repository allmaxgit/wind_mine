package types

type (
	Buyer struct {
		Id           int
		EthAddr      string `sql:"type:varchar(100),unique"`
		BtcAddr      string `sql:"type:varchar(100),unique"`
		Transactions []*Transaction
	}

	Transaction struct {
		Id      int
		BuyerId int
		Buyer   *Buyer
		From    string `sql:"type:varchar(100),notnull"`
		Amount  float64
		Hash    string `sql:"type:varchar(100),unique"`
	}

	NotHandledTransaction struct {
		Id     int
		From   string `sql:"type:varchar(100),notnull"`
		Amount float64
		Hash   string `sql:"type:varchar(100),unique"`
	}

	NotHandledBTCReturn struct {
		Id     int
		From   string `sql:"type:varchar(100),notnull"`
		Amount float64
		Hash   string `sql:"type:varchar(100),unique"`
		// IMPORTANT this flag means that backend has failed to check whether user had passed KYC or not.
		// If KycFailed is set to true, then backend should try to check KYC again and then act correspondingly.
		// If KycFailed is set to false, then user hasn't passed KYC and backend should just return BTC.
		KycFailed bool
	}
)
