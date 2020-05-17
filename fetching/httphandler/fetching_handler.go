package httphandler

import (
	"efishery.com/micro/dto/request"
	"efishery.com/micro/dto/response"
	"efishery.com/micro/mapper"
	"efishery.com/micro/shared/domains"
	"efishery.com/micro/shared/utils"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

type FetchingHandler struct {
	RedisClient *redis.Client
	RestService utils.RestService
}

func (h *FetchingHandler) Fetch(c echo.Context) error {
	listUrl := viper.GetString("application.resources.urls.list")
	convertUrl := viper.GetString("application.resources.urls.convertion")

	restRes, err := h.RestService.Get(listUrl, nil, nil)
	if err != nil || restRes.StatusCode != 200 {
		return c.JSON(http.StatusInternalServerError, domains.ResponseError{Message: domains.ErrInternalServerError.Error()})
	}

	var fetchingMappers []request.FetchingMapperDTO
	err = json.Unmarshal(restRes.Body, &fetchingMappers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ResponseError{Message: domains.ErrInternalServerError.Error()})
	}

	var usdToIdr float64
	result, err := h.RedisClient.Get("usd_to_idr").Result()
	if err != nil || result == "" {
		restRes, _ := h.RestService.Get(convertUrl, nil, nil)
		var convertMapper request.ConvertionMapperDTO
		err = json.Unmarshal(restRes.Body, &convertMapper)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, domains.ResponseError{Message: domains.ErrInternalServerError.Error()})
		}

		h.RedisClient.Set("usd_to_idr", convertMapper.UsdToIdr, 0)
		usdToIdr = convertMapper.UsdToIdr
	} else {
		usdToIdr, _ = strconv.ParseFloat(result, 64)
	}

	res := mapper.ToListToFetchingDTOs(&fetchingMappers, func(m *request.FetchingMapperDTO, dto *response.FetchingDTO) {
		if dto.Price == nil {
			return
		}

		priceIDR, _ := strconv.ParseFloat(*dto.Price, 64)
		priceUSD := utils.ConvertIDRTOUSD(priceIDR, usdToIdr)
		priceUSDStr := fmt.Sprintf("%f", priceUSD)
		dto.PriceInUSD = &priceUSDStr
	})

	return c.JSON(http.StatusOK, res)
}

func (h *FetchingHandler) Aggregate(c echo.Context) error {
	listUrl := viper.GetString("application.resources.urls.list")

	restRes, err := h.RestService.Get(listUrl, nil, nil)
	if err != nil || restRes.StatusCode != 200 {
		return c.JSON(http.StatusInternalServerError, domains.ResponseError{Message: domains.ErrInternalServerError.Error()})
	}

	var fetchingMappers []request.FetchingMapperDTO
	err = json.Unmarshal(restRes.Body, &fetchingMappers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ResponseError{Message: domains.ErrInternalServerError.Error()})
	}

	res := mapper.ToListToFetchingDTOs(&fetchingMappers, func(m *request.FetchingMapperDTO, dto *response.FetchingDTO) {})

	return c.JSON(http.StatusOK, res)
}

func (h *FetchingHandler) ClaimToken(c echo.Context) error {
	name := c.Get("name").(string)
	phone := c.Get("phone").(string)
	role := c.Get("role").(string)

	return c.JSON(http.StatusOK, mapper.ToClaimTokenDTO(name, phone, role))
}
