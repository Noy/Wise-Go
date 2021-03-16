package wisego

type Statements struct {
	AccountHolder struct {
		Type    string `json:"type"`
		Address struct {
			AddressFirstLine string `json:"addressFirstLine"`
			City             string `json:"city"`
			PostCode         string `json:"postCode"`
			StateCode        string `json:"stateCode"`
			Country          string `json:"countryName"`
		} `json:"address"`
		BusinessName       string `json:"businessName"`
		RegistrationNumber string `json:"registrationNumber"`
		CompanyType        string `json:"companyType"`
	} `json:"accountHolder"`
	Issuer struct {
		Name      string `json:"name"`
		FirstLine string `json:"firstLine"`
		City      string `json:"city"`
		PostCode  string `json:"postCode"`
		StateCode string `json:"stateCode"`
		Country   string `json:"country"`
	} `json:"issuer"`
	BankDetails  []BankDetails  `json:"bankDetails"`
	Transactions []Transactions `json:"transactions"`
}

type BankDetails struct {
	Address struct {
		FirstLine  string `json:"firstLine"`
		SecondLine string `json:"secondLine"`
		City       string `json:"city"`
		PostCode   string `json:"postCode"`
		StateCode  string `json:"stateCode"`
		Country    string `json:"country"`
	} `json:"address"`
	AccountNumbers []struct {
		AccountType   string `json:"accountType"`
		AccountNumber string `json:"accountNumber"`
	} `json:"accountNumbers"`
	BankCodes []struct {
		Scheme string `json:"scheme"`
		Value  string `json:"value"`
	} `json:"bankCodes"`
}

type Transactions struct {
	Type   string `json:"type"`
	Date   string `json:"date"`
	Amount struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"amount"`
	TotalFees struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"totalFees"`
	Details struct {
		Type         string `json:"type"`
		Description  string `json:"description"`
		SourceAmount struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"sourceAmount"`
		TargetAmount struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"targetAmount"`
		Rate float64 `json:"rate"`
	} `json:"details"`
	ExchangeDetails struct {
		ToAmount struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"toAmount"`
		FromAmount struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"fromAmount"`
		Rate float64 `json:"rate"`
	} `json:"exchangeDetails"`
	RunningBalance struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"runningBalance"`
	ReferenceNumber string `json:"referenceNumber"`
}

type Recipient struct {
	ID                int64  `json:"id"`
	Business          int64  `json:"business"`
	Profile           int64  `json:"profile"`
	AccountHolderName string `json:"accountHolderName"`
	Currency          string `json:"currency"`
	Country           string `json:"country"`
	Type              string `json:"type"`
	Details           struct {
		AccountNumber string `json:"accountNumber"`
		SortCode      string `json:"sortCode"`
		IBAN          string `json:"IBAN"`
		BIC           string `json:"BIC"`
		PhoneNumber   string `json:"phoneNumber"`
		Abartn        string `json:"abartn"`
	} `json:"details"`
}

type ConfirmedTransaction struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
}