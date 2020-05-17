package utils

import (
	"efishery.com/micro/dto/request"
	"fmt"
	"strconv"
	"time"
)

func Grouping(lists []request.FetchingMapperDTO) map[string][]request.FetchingMapperDTO {
	grouped := make(map[string][]request.FetchingMapperDTO)

	for _, list := range lists {
		if list.AreaProvinsi == nil || list.Timestamp == nil || list.Price == nil {
			continue
		}

		i, _ := strconv.ParseInt(*list.Timestamp, 10, 64)
		t := time.Unix(i, 0).UTC()
		y, w := t.ISOWeek()
		key := fmt.Sprintf("%s#%d#%d", *list.AreaProvinsi, y, w)
		grouped[key] = append(grouped[key], list)
	}

	return grouped
}

func GetSum(lists []request.FetchingMapperDTO, fn func(i *request.FetchingMapperDTO) float64) float64 {
	var sum float64

	for _, list := range lists {
		sum += fn(&list)
	}

	return sum
}
