package delivery

import "capstone/happyApp/features/cart"

type ResposeCommunity struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
}

type ResponseCart struct {
	ID           uint   `json:"cartid"`
	ProductID    uint   `json:"productid"`
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	Photo        string `json:"photo"`
	Price        int    `json:"price"`
}

func CoreToResCommunity(data cart.CoreCommunity) ResposeCommunity {
	return ResposeCommunity{
		ID:           data.ID,
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
		Members:      data.Members,
	}
}

func CoreToResponseCart(data cart.CoreCart) ResponseCart {
	return ResponseCart{
		ID:           data.ID,
		ProductID:    data.ProductID,
		Name:         data.Name,
		Descriptions: data.Descriptions,
		Photo:        data.Photo,
		Price:        data.Price,
	}
}

func CoreToResponseCartList(data []cart.CoreCart) []ResponseCart {
	var list []ResponseCart
	for _, v := range data {
		list = append(list, CoreToResponseCart(v))
	}

	return list
}
