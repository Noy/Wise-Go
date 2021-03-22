package wisego

type Address struct {
	Id      int64 `json:"id"`
	Profile int64 `json:"profile"`
	Details struct {
		Country     string        `json:"country"`
		FirstLine   string        `json:"firstLine"`
		PostCode    string        `json:"postCode"`
		City        string        `json:"city"`
		State       string        `json:"state"`
		Occupation  string        `json:"occupation"`
		Occupations []Occupations `json:"occupations"`
	} `json:"details"`
}

type Occupations struct {
	Code   string `json:"code"`
	Format string `json:"format"`
}
