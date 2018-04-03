package types

type (
	BTCServiceReq struct {
		Type          string
		Address       string
		Value         float64
		ReceiveTXHash string
		To            string
	}

	BTCServiceResp struct {
		Type   string
		Value  float64
		From   string
		TXHash string
	}

	RateServiceResp struct {
		Currency     uint8
		Value        float64
		FiatCurrency string
	}
)
