package product

type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	ImgUrl string `json:"img_url"`
}
