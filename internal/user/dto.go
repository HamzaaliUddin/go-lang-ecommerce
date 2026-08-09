package user

type UserResponse struct {
	ID        uint     `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	IsActive  bool     `json:"isActive"`
	Roles     []string `json:"roles"`
}

func toUserResponse(account User) UserResponse {
	roles := make([]string, 0, len(account.Roles))

	for _, item := range account.Roles {
		roles = append(roles, item.Slug)
	}

	return UserResponse{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
		IsActive:  account.IsActive,
		Roles:     roles,
	}
}