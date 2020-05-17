package mapper

import (
	"efishery.com/micro/dto/request"
	"efishery.com/micro/dto/response"
	"efishery.com/micro/shared/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ToAggregateDTO(province string, year string, week string, min string, max string, med string,
	avg string) response.AggregateDTO {

	return response.AggregateDTO{
		AreaProvinsi: province,
		Tahun:        year,
		MingguKe:     week,
		Min:          min,
		Max:          max,
		Median:       med,
		Avg:          avg,
	}
}

func ToListAggregateDTOs(grouped map[string][]request.FetchingMapperDTO) []response.AggregateDTO {
	var dtos []response.AggregateDTO

	for k, lists := range grouped {
		keys := strings.Split(k, "#")
		sort.SliceStable(lists, func(i, j int) bool {
			priceI, _ := strconv.ParseFloat(*lists[i].Price, 64)
			priceJ, _ := strconv.ParseFloat(*lists[j].Price, 64)

			return priceI < priceJ
		})
		listLen := len(lists)
		min := *lists[0].Price
		max := *lists[listLen-1].Price
		med := utils.CalcMedian(len(lists))
		sumOfPrice := utils.GetSum(lists, func(i *request.FetchingMapperDTO) float64 {
			price, _ := strconv.ParseFloat(*i.Price, 64)
			return price
		})
		avg := sumOfPrice / float64(listLen)

		dtos = append(dtos, ToAggregateDTO(keys[0], keys[1], keys[2], min, max, fmt.Sprintf("%f", med),
			fmt.Sprintf("%f", avg)))
	}

	return dtos
}
