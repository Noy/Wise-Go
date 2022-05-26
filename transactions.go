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
	RecipientContent []RecipientContent `json:"content"`
	Sort             struct {
		Empty    bool `json:"empty"`
		Unsorted bool `json:"unsorted"`
		Sorted   bool `json:"sorted"`
	} `json:"sort"`
	Size int64 `json:"size"`
}

type RecipientContent struct {
	ID        int64 `json:"id"`
	CreatorId int64 `json:"creatorId"`
	ProfileId int64 `json:"profileId"`
	Name      struct {
		FullName                 string `json:"fullName"`
		GivenName                string `json:"givenName"`
		FamilyName               string `json:"familyName"`
		MiddleName               string `json:"middleName"`
		PatronymicName           string `json:"patronymicName"`
		CannotHavePatronymicName bool   `json:"cannotHavePatronymicName"`
	} `json:"name"`
	Currency        string `json:"currency"`
	Country         string `json:"country"`
	Type            string `json:"type"`
	LegalEntityType string `json:"legalEntityType"`
	Active          bool   `json:"active"`
	Details         struct {
		BankName                   string `json:"bankName"`
		Iban                       string `json:"iban"`
		SwiftCode                  string `json:"swiftCode"`
		Bic                        string `json:"bic"`
		Reference                  string `json:"reference"`
		BalanceAccountProfileId    string `json:"balanceAccountProfileId"`
		AccountNumber              string `json:"accountNumber"`
		SortCode                   string `json:"sortCode"`
		HashedByLooseHashAlgorithm string `json:"hashedByLooseHashAlgorithm"`
	} `json:"details"`
	CommonFieldMap struct {
		AccountNumberField string `json:"accountNumberField"`
		BankCodeField      string `json:"bankCodeField"`
	} `json:"commonFieldMap"`
	IsDefaultAccount   bool            `json:"isDefaultAccount"`
	Hash               string          `json:"hash"`
	AccountSummary     string          `json:"accountSummary"`
	LongAccountSummary string          `json:"longAccountSummary"`
	DisplayFields      []DisplayFields `json:"displayFields"`
	OwnedByCustomer    bool            `json:"ownedByCustomer"`
}
type DisplayFields struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ConfirmedTransaction struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
}
