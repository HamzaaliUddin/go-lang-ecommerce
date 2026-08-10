package banner

type BannerResponse struct {
	ID       uint   `json:"id"`
	ImageURL string `json:"imageUrl"`
	LinkURL  string `json:"linkUrl"`
}

func toBannerResponse(banner Banner) BannerResponse {
	return BannerResponse{
		ID:       banner.ID,
		ImageURL: banner.ImageURL,
		LinkURL:  banner.LinkURL,
	}
}