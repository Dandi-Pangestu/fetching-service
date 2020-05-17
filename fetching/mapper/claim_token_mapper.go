package mapper

import "efishery.com/micro/dto/response"

func ToClaimTokenDTO(name string, phone string, role string, timeStamp string) response.ClaimTokenDTO {
	return response.ClaimTokenDTO{
		Name:      name,
		Phone:     phone,
		Role:      role,
		TimeStamp: timeStamp,
	}
}
