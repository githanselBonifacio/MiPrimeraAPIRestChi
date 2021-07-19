package smartphones

//smartphone model:  structure for smartphone
type SmartPhone struct {
	Id           int64 //similar to a long in java
	Name         string
	Price        int64
	CoutryOrigin string
	OS           string
}

// Create smartPhone
type CreateSmartPhone struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	CoutryOrigin string `json:"countryOrigin"`
	OS           string `json:"os"`
}
