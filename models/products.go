package models

type Product struct {
	Id                    int     `json:"id"`
	Productkey            int     `json:"ProductKey"`
	Productsubcategorykey int     `json:"ProductSubcategoryKey"`
	Productsku            string  `json:"ProductSKU"`
	Productname           string  `json:"ProductName"`
	Modelname             string  `json:"ModelName"`
	Productdescription    string  `json:"ProductDescription"`
	Productcolor          string  `json:"ProductColor"`
	Productsize           string  `json:"ProductSize"`
	Productstyle          string  `json:"ProductStyle"`
	Productcost           float64 `json:"ProductCost"`
	Productprice          float64 `json:"ProductPrice"`
}

var Products []Product
