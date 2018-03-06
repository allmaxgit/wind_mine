package types

type BTCServiceReq struct {
	Type    string
	Address string
}

type BTCServiceResp struct {
	Type  string
	Value float64
	From  string
}
