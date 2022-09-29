package delivery

import "capstone/happyApp/features/user"

type Response struct {
	Photo     string        `json:"photo"`
	Name      string        `json:"name"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Gender    string        `json:"gender"`
	Community []myCommunity `json:"community"`
}

type myCommunity struct {
	ID    uint   `json:"id"`
	Logo  string `json:"logo"`
	Title string `json:"title"`
	Role  string `json:"role"`
}

func toRespon(data user.CoreUser, comu []user.CommunityProfile) Response {

	var myComu []myCommunity
	for _, v := range comu {
		myComu = append(myComu, myCommunity{
			ID:    v.ID,
			Logo:  v.Logo,
			Title: v.Title,
			Role:  v.Role,
		})
	}

	return Response{
		Photo:     data.Photo,
		Name:      data.Name,
		Username:  data.Username,
		Gender:    data.Gender,
		Email:     data.Email,
		Community: myComu,
	}

}
