package wisego

type Profile struct {
	Type    string `json:"type"`
	ID      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Address struct {
		AddressFirstLine string `json:"addressFirstLine"`
		City             string `json:"city"`
		CountryIso2Code  string `json:"countryIso2Code"`
		CountryIso3Code  string `json:"countryIso3Code"`
		PostCode         string `json:"postCode"`
		StateCode        string `json:"stateCode"`
	}
	Email                 string   `json:"email"`
	CreatedAt             string   `json:"createdAt"`
	UpdatedAt             string   `json:"updatedAt"`
	Obfuscated            bool     `json:"obfuscated"`
	Avatar                string   `json:"avatar"`
	BusinessName          string   `json:"businessName"`
	RegistrationNumber    string   `json:"registrationNumber"`
	DescriptionOfBusiness string   `json:"descriptionOfBusiness"`
	CompanyType           string   `json:"companyType"`
	CompanyRole           string   `json:"companyRole"`
	FirstLevelCategory    string   `json:"firstLevelCategory"`
	SecondLevelCategory   string   `json:"secondLevelCategory"`
	OperationalAddresses  []string `json:"operationalAddresses"`
	FullName              string   `json:"fullName"`
	//Personal
	CurrentState string `json:"currentState"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PlaceOfBirth struct {
		RawValue        string `json:"rawValue"`
		City            string `json:"city"`
		CountryIso2Code string `json:"countryIso2Code"`
		CountryIso3Code string `json:"countryIso3Code"`
		PostCode        string `json:"postCode"`
		StateCode       string `json:"stateCode"`
	} `json:"placeOfBirth"`
	PhoneNumber string `json:"phoneNumber"`
	Occupation  string `json:"occupation"`
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
	ID       int64  `json:"id"`
	Currency string `json:"currency"`
	Amount   struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"amount"`
	ReservedAmount struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"reservedAmount"`
	CreationTime     string               `json:"creationTime"`
	ModificationTime string               `json:"modificationTime"`
	Visible          bool                 `json:"visible"`
	BankDetails      []BalanceBankDetails `json:"bankDetails"`
}

type BalanceBankDetails struct {
	ID                int64  `json:"id"`
	ProfileId         int64  `json:"profileId"`
	Currency          string `json:"currency"`
	AccountHolderName string `json:"accountHolderName"`
	BankCode          struct {
		SwiftCode        string `json:"swiftCode"`
		SortCode         string `json:"sortCode"`
		AchRoutingNumber string `json:"achRoutingNumber"`
	} `json:"bankCode"`
	BankFeatures []struct {
		Description string `json:"description"`
		Key         string `json:"key"`
		Id          int64  `json:"id"`
		Supported   bool   `json:"supported"`
	} `json:"bankFeatures"`
	AccountNumber struct {
		Iban      string `json:"iban"`
		LocalIban string `json:"localIban"`
		Default   string `json:"default"`
	} `json:"accountNumber"`
	Address struct {
		Default struct {
			FirstLine  string `json:"firstLine"`
			SecondLine string `json:"secondLine"`
			PostCode   string `json:"postCode"`
			City       string `json:"city"`
			Country    string `json:"country"`
			StateCode  string `json:"stateCode"`
		}
	} `json:"address"`
	Limits struct {
		Daily struct {
			Currency string `json:"currency"`
			Value    int64  `json:"value"`
		} `json:"daily"`
		Yearly struct {
			Currency string `json:"currency"`
			Value    int64  `json:"value"`
		} `json:"yearly"`
	} `json:"limits"`
	Translations struct {
		AccountHolderName   string `json:"accountHolderName"`
		HolderName          string `json:"holderName"`
		AccountHolderIban   string `json:"accountNumber.iban"`
		BankCodeSwift       string `json:"bankCode.swiftCode"`
		AddressDefault      string `json:"address.default"`
		Title               string `json:"title"`
		AccountLimitsYearly string `json:"accountLimits.yearly"`
		AccountLimitsDaily  string `json:"accountLimits.daily"`
	} `json:"translations"`
	Status             string `json:"status"`
	VerificationStatus string `json:"verificationStatus"`
	Deprecated         bool   `json:"deprecated"`
}
