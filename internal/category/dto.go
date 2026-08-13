package category

type CategoryResponse struct{
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	IsActive string `json:"isActive"`
	
}


func toCategoryResponse(category CategoryResponse) CategoryResponse {
	return CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		Slug:     category.Slug,
		IsActive: category.IsActive,
	}
}