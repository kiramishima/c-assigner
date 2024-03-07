package models

// Credit struct
type Credit struct {
	Invest    int32 `json:"-"`
	Credit300 int32 `json:"credit_type_300"`
	Credit500 int32 `json:"credit_type_500"`
	Credit700 int32 `json:"credit_type_700"`
	Status    int   `json:"-"`
}
