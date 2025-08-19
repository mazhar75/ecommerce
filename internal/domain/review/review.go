package review

type Reviews struct {
	ReviewId  int    `json:"review_id"`
	ProductId int    `json:"product_id"`
	UserId    int    `json:"user_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
