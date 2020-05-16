package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
)

type RestService interface {
	CombinedHeaders(headers interface{}) interface{}
	GetFullURIs(uri string, queryParams interface{}) string
	Execute(method string, uri string, queryParams interface{}, requestBody interface{}, headers interface{}) (RestServiceResponse, error)
	Post(uri string, queryParams interface{}, requestBody interface{}, headers interface{}) (RestServiceResponse, error)
	Put(uri string, queryParams interface{}, requestBody interface{}, headers interface{}) (RestServiceResponse, error)
	Patch(uri string, queryParams interface{}, requestBody interface{}, headers interface{}) (RestServiceResponse, error)
	Delete(uri string, queryParams interface{}, headers interface{}) (RestServiceResponse, error)
	Get(uri string, queryParams interface{}, headers interface{}) (RestServiceResponse, error)
}

type RestServiceResponse struct {
	ResponseOrigin interface{}
	Body           []byte
	Status         string
	StatusCode     int
	Header         http.Header
	Cookies        []*http.Cookie
}

func (r *RestServiceResponse) ToString() string {
	if r.Body == nil {
		return ""
	}

	return strings.TrimSpace(string(r.Body))
}

type restyRestService struct {
	client *resty.Client
}

func NewRestyRestService() RestService {
	client := resty.New()
	return &restyRestService{
		client: client,
	}
}

func (s *restyRestService) CombinedHeaders(headers interface{}) interface{} {
	defaultHeaders := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	if headers == nil {
		return defaultHeaders
	}

	for key, header := range headers.(map[string]interface{}) {
		defaultHeaders[key] = fmt.Sprint(header)
	}

	return defaultHeaders
}

func (s *restyRestService) GetFullURIs(uri string, queryParams interface{}) string {
	if queryParams == nil {
		return uri
	}

	parameters := url.Values{}

	for key, value := range queryParams.(map[string]interface{}) {
		parameters.Add(key, fmt.Sprint(value))
	}

	return uri + "?" + parameters.Encode()
}

func (s *restyRestService) Execute(method string, uri string, queryParams interface{}, requestBody interface{},
	headers interface{}) (RestServiceResponse, error) {

	errChan := make(chan error)
	responseChan := make(chan *resty.Response)
	var response RestServiceResponse

	go func() {
		var response *resty.Response
		var err error

		request := s.client.R().SetHeaders(s.CombinedHeaders(headers).(map[string]string))
		if requestBody != nil {
			request = request.SetBody(requestBody)
		}

		if method == http.MethodPost {
			response, err = request.Post(s.GetFullURIs(uri, queryParams))
		} else if method == http.MethodPut {
			response, err = request.Put(s.GetFullURIs(uri, queryParams))
		} else if method == http.MethodPatch {
			response, err = request.Patch(s.GetFullURIs(uri, queryParams))
		} else if method == http.MethodDelete {
			response, err = request.Delete(s.GetFullURIs(uri, queryParams))
		} else {
			response, err = request.Get(s.GetFullURIs(uri, queryParams))
		}

		errChan <- err
		responseChan <- response
	}()

	if err := <-errChan; err != nil {
		return RestServiceResponse{}, err
	}

	restResponse := <-responseChan
	response.ResponseOrigin = restResponse
	response.Body = restResponse.Body()
	response.Status = restResponse.Status()
	response.StatusCode = restResponse.StatusCode()
	response.Header = restResponse.Header()
	response.Cookies = restResponse.Cookies()

	return response, nil
}

func (s *restyRestService) Post(uri string, queryParams interface{}, requestBody interface{},
	headers interface{}) (RestServiceResponse, error) {

	return s.Execute(http.MethodPost, uri, queryParams, requestBody, headers)
}

func (s *restyRestService) Put(uri string, queryParams interface{}, requestBody interface{},
	headers interface{}) (RestServiceResponse, error) {

	return s.Execute(http.MethodPut, uri, queryParams, requestBody, headers)
}

func (s *restyRestService) Patch(uri string, queryParams interface{}, requestBody interface{},
	headers interface{}) (RestServiceResponse, error) {

	return s.Execute(http.MethodPatch, uri, queryParams, requestBody, headers)
}

func (s *restyRestService) Delete(uri string, queryParams interface{}, headers interface{}) (RestServiceResponse, error) {
	return s.Execute(http.MethodDelete, uri, queryParams, nil, headers)
}

func (s *restyRestService) Get(uri string, queryParams interface{}, headers interface{}) (RestServiceResponse, error) {
	return s.Execute(http.MethodGet, uri, queryParams, nil, headers)
}
