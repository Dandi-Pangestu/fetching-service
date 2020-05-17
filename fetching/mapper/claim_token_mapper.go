package mapper

import "efishery.com/micro/dto/response"

func ToClaimTokenDTO(name string, phone string, role string) response.ClaimTokenDTO {
	return response.ClaimTokenDTO{
		Name:  name,
		Phone: phone,
		Role:  role,
	}
}
