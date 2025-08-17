package product

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json: "description"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}
