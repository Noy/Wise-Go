package wisego

type Profile struct {
	ID      int64  `json:"id"`
	Type    string `json:"type"`
	Details struct {
		BusinessName          string `json:"name"`
		RegistrationNumber    string `json:"registrationNumber"`
		ACN                   string `json:"acn"`
		ABN                   string `json:"abn"`
		ARBN                  string `json:"arbn"`
		CompanyType           string `json:"companyType"`
		CompanyRole           string `json:"companyRole"`
		DescriptionOfBusiness string `json:"descriptionOfBusiness"`
		Webpage               string `json:"webpage"`
		BusinessCategory      string `json:"businessCategory"`
		BusinessSubCategory   string `json:"businessSubCategory"`
		//PERSONAL
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		DateOfBirth string `json:"dateOfBirth"`
		PhoneNumber string `json:"phoneNumber"`
		Avatar      string `json:"avatar"`
		Occupation  string `json:"occupation"`
		Occupations string `json:"occupations"`
		//BOTH
		PrimaryAddress int64 `json:"primaryAddress"`
	} `json:"details"`
}

type BorderlessAccounts struct {
	ID               int64      `json:"id"`
	ProfileID        int64      `json:"profileId"`
	RecipientID      int64      `json:"recipientId"`
	CreationTime     string     `json:"creationTime"`
	ModificationTime string     `json:"modificationTime"`
	Active           bool       `json:"active"`
	Eligible         bool       `json:"eligible"`
	Balances         []Balances `json:"balances"`
}

type Balances struct {
	ID          int64  `json:"id"`
	BalanceType string `json:"balanceType"`
	Currency    string `json:"currency"`
	Amount      struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"amount"`
	ReservedAmount struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"reservedAmount"`
	BankDetails struct {
		ID                int64  `json:"id"`
		Currency          string `json:"currency"`
		BankCode          string `json:"bankCode"`
		AccountNumber     string `json:"accountNumber"`
		IBAN              string `json:"iban"`
		BankName          string `json:"bankName"`
		AccountHolderName string `json:"accountHolderName"`
		BankAddress       struct {
			AddressFirstLine string `json:"addressFirstLine"`
			PostCode         string `json:"postCode"`
			City             string `json:"city"`
			Country          string `json:"country"`
			StateCode        string `json:"stateCode"`
		} `json:"bankAddress"`
	} `json:"bankDetails"`
}
