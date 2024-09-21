package model

type Shop struct {
	Packages []Package `json:"packages"`
}

type Package struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Price  string `json:"price"`
	Detail string `json:"detail"`
	Open   string `json:"open"`
}
