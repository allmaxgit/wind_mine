package types

type Buyer struct {
	Id      int64
	EthAddr string   `sql:"type:varchar(255),unique"`
	BtcAddr string   `sql:"type:varchar(255),unique"`
}

type Transaction struct {
	Id     int64
	UserId int64
	From   string
	Amount int64
	Hash   string
}
