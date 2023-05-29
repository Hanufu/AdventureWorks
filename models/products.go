package models

type Product struct {
	ProductKey            int     `json:"productKey"`
	ProductSubcategoryKey int     `json:"productSubcategoryKey"`
	ProductName           string  `json:"productName"`
	ModelName             string  `json:"modelName"`
	ProductDescription    string  `json:"productDescription"`
	ProductColor          string  `json:"productColor"`
	ProductSize           string  `json:"productSize"`
	ProductStyle          string  `json:"productStyle"`
	ProductCost           float64 `json:"productCost"`
	ProductPrice          float64 `json:"productPrice"`
}

var Products []Product
