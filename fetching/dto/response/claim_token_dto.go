package response

type ClaimTokenDTO struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	TimeStamp string `json:"timestamp"`
}
