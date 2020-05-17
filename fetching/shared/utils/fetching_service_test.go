package utils

import (
	"efishery.com/micro/dto/request"
	"github.com/magiconair/properties/assert"
	"strconv"
	"testing"
)

func TestGetSum(t *testing.T) {
	price1 := "20000"
	price2 := "25000"
	lists := []request.FetchingMapperDTO{
		{
			Price: &price1,
		},
		{
			Price: &price2,
		},
	}
	expected := float64(45000)

	result := GetSum(lists, func(i *request.FetchingMapperDTO) float64 {
		price, _ := strconv.ParseFloat(*i.Price, 64)
		return price
	})
	assert.Equal(t, result, expected)
}

func TestGrouping(t *testing.T) {
	uuid1 := "8eba7e44-49da-4393-a625-c77a3441dd3a"
	komoditas1 := "Ikan Tuna Update Data"
	provinsi1 := "BANTEN"
	kota1 := "PANDEGLANG"
	size1 := "80"
	price1 := "88000"
	tglParsed1 := "2019-11-11T17:00:00.000Z"
	timestamp1 := "1573491600"

	uuid2 := "0ddd2c7f-d740-4c92-b3a8-7b66da6a0dc1"
	komoditas2 := "Penaeus Vannamei"
	provinsi2 := "SUMATERA BARAT"
	kota2 := "PADANG PARIAMAN"
	size2 := "160"
	price2 := "30000"
	tglParsed2 := "2019-11-11T17:00:00.000Z"
	timestamp2 := "1573491600"

	lists := []request.FetchingMapperDTO{
		{
			Uuid:         &uuid1,
			Komoditas:    &komoditas1,
			AreaProvinsi: &provinsi1,
			AreaKota:     &kota1,
			Size:         &size1,
			Price:        &price1,
			TglParsed:    &tglParsed1,
			Timestamp:    &timestamp1,
		},
		{
			Uuid:         &uuid2,
			Komoditas:    &komoditas2,
			AreaProvinsi: &provinsi2,
			AreaKota:     &kota2,
			Size:         &size2,
			Price:        &price2,
			TglParsed:    &tglParsed2,
			Timestamp:    &timestamp2,
		},
	}

	expected := make(map[string][]request.FetchingMapperDTO)
	expected["BANTEN#2019#46"] = []request.FetchingMapperDTO{
		lists[0],
	}
	expected["SUMATERA BARAT#2019#46"] = []request.FetchingMapperDTO{
		lists[1],
	}

	result := Grouping(lists)
	assert.Equal(t, result, expected)
}
