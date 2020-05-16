package mapper

import (
	"efishery.com/micro/dto/request"
	"efishery.com/micro/dto/response"
)

func ToFetchingDTO(m *request.FetchingMapperDTO,
	fn func(m *request.FetchingMapperDTO, dto *response.FetchingDTO)) response.FetchingDTO {

	dto := response.FetchingDTO{
		Uuid:         m.Uuid,
		AreaKota:     m.AreaKota,
		AreaProvinsi: m.AreaProvinsi,
		Komoditas:    m.Komoditas,
		Price:        m.Price,
		PriceInUSD:   nil,
		Size:         m.Size,
		TglParsed:    m.TglParsed,
		Timestamp:    m.Timestamp,
	}

	fn(m, &dto)

	return dto
}

func ToListToFetchingDTOs(m *[]request.FetchingMapperDTO,
	fn func(m *request.FetchingMapperDTO, dto *response.FetchingDTO)) []response.FetchingDTO {

	dtos := make([]response.FetchingDTO, len(*m))
	for i, item := range *m {
		dto := ToFetchingDTO(&item, fn)
		dtos[i] = dto
	}

	return dtos
}
