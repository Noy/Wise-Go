package wisego

type Quote struct {
	ID               int64   `json:"id"`
	Source           string  `json:"source"`
	Target           string  `json:"target"`
	SourceAmount     float64 `json:"sourceAmount"`
	TargetAmount     float64 `json:"targetAmount"`
	Type             string  `json:"type"`
	Rate             float64 `json:"rate"`
	CreatedTime      string  `json:"createdTime"`
	CreatedByUserId  int64   `json:"createdByUserId"`
	Profile          int64   `json:"profile"`
	Business         int64   `json:"business"`
	RateType         string  `json:"rateType"`
	DeliveryEstimate string  `json:"deliveryEstimate"`
	Fee              float64 `json:"fee"`
	FeeDetails       struct {
		TransferWise float64 `json:"transferwise"`
		PayIn        float64 `json:"payIn"`
		Discount     float64 `json:"discount"`
		PriceSetId   float64 `json:"priceSetId"`
		Partner      float64 `json:"partner"`
	} `json:"feeDetails"`
	AllowedProfileTypes    []string `json:"allowedProfileTypes"`
	GuaranteedTargetAmount bool     `json:"guaranteedTargetAmount"`
	OfSourceAmount         bool     `json:"ofSourceAmount"`
}

type Transfer struct {
	ID                    int64   `json:"id"`
	User                  int64   `json:"user"`
	TargetAccount         int64   `json:"targetAccount"`
	QuoteID               int64   `json:"quote"`
	Status                string  `json:"status"`
	Reference             string  `json:"reference"`
	Rate                  float64 `json:"rate"`
	Created               string  `json:"created"`
	Business              int64   `json:"business"`
	HasActiveIssues       bool    `json:"hasActiveIssues"`
	SourceCurrency        string  `json:"sourceCurrency"`
	SourceValue           float64 `json:"sourceValue"`
	TargetCurrency        string  `json:"targetCurrency"`
	TargetValue           float64 `json:"targetValue"`
	CustomerTransactionID string  `json:"customerTransactionId"`
}

type TransferInitiationDetails struct {
	ID              int64
	Created         string
	Fee             string
	Rate            string
	TargetAmount    string
	HasActiveIssues bool
	SourceCurrency  string
	SourceValue     string
	TargetCurrency  string
	TargetValue     string
}
