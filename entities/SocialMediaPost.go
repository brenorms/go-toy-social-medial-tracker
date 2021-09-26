package entities

type SocialMediaPost struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
	Likes int    `json:"likes"`
}
