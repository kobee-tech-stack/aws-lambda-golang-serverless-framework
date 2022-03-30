package services

type Product struct {
	Id    string  `dynamodbav:"id" json:"id"`
	Name  string  `dynamodbav:"name" json:"name"`
	Price float64 `dynamodbav:"price" json:"price"`
}

type ProductRange struct {
	Products []Product `json:"products"`
	Next     *string   `json:"next,omitempty"`
}

type ProductModel struct {
	Pk    string  `dynamodbav:"pk" json:"pk"`
	Sk    string  `dynamodbav:"sk" json:"sk"`
	Id    string  `dynamodbav:"id" json:"id"`
	Name  string  `dynamodbav:"name" json:"name"`
	Price float64 `dynamodbav:"price" json:"price"`
}
