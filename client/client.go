package client

import (
	"dhis2cli/config"
	"errors"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

func GetDHIS2BaseURL(url string) (string, error) {
	if strings.Contains(url, "/api/") {
		pos := strings.Index(url, "/api/")
		return url[:pos], nil
	}
	return url, errors.New("URL doesn't contain /api/ part")
}

type Client struct {
	RestClient *resty.Client
	BaseURL    string
}

type Server struct {
	BaseUrl    string `json:"base_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthToken  string `json:"auth_token"`
	AuthMethod string `json:"auth_method"`
}

var Dhis2Client *Client
var Dhis2Server *Server

func InitServer() {
	Dhis2Server = &Server{
		BaseUrl:    config.Cfg.Server.BaseUrl,
		Username:   config.Cfg.Server.Username,
		Password:   config.Cfg.Server.Password,
		AuthToken:  config.Cfg.Server.AuthToken,
		AuthMethod: config.Cfg.Server.AuthMethod,
	}
	// Dhis2Client, _= Dhis2Server.NewClient()
}

func (s *Server) NewClient() (*Client, error) {
	client := resty.New()
	baseUrl, err := GetDHIS2BaseURL(s.BaseUrl)
	if err != nil {
		log.WithFields(log.Fields{
			"URL": s.BaseUrl, "Error": err}).Error("Failed to get base URL from URL")
		return nil, err
	}
	client.SetBaseURL(baseUrl + "/api")
	client.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   "Dispatcher2-Go",
	})
	client.SetDisableWarn(true)
	switch s.AuthMethod {
	case "Basic":
		client.SetBasicAuth(s.Username, s.Password)
	case "Token":
		client.SetAuthScheme("Token")
		client.SetAuthToken(s.AuthToken)
	}
	return &Client{
		RestClient: client,
		BaseURL:    baseUrl + "/api",
	}, nil
}

func (c *Client) GetResource(resourcePath string, params map[string]string) (*resty.Response, error) {
	request := c.RestClient.R()

	if params != nil {
		request.SetQueryParams(params)
	}

	resp, err := request.Get(resourcePath)
	if err != nil {
		log.Fatalf("Error when calling `GetResource`: %v", err)
	}
	return resp, err
}

func (c *Client) PostResource(resourcePath string, data interface{}) (*resty.Response, error) {
	resp, err := c.RestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(resourcePath)
	if err != nil {
		log.Fatalf("Error when calling `PostResource`: %v", err)
	}
	return resp, err
}

func (c *Client) PutResource(resourcePath string, data interface{}) (*resty.Response, error) {
	resp, err := c.RestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Put(resourcePath)
	if err != nil {
		log.Fatalf("Error when calling `PutResource`: %v", err)
	}
	return resp, err
}

func (c *Client) DeleteResource(resourcePath string) (*resty.Response, error) {
	resp, err := c.RestClient.R().
		Delete(resourcePath)
	if err != nil {
		log.Fatalf("Error when calling `DeleteResource`: %v", err)
	}
	return resp, err
}

func (c *Client) PatchResource(resourcePath string, data interface{}) (*resty.Response, error) {
	resp, err := c.RestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Patch(resourcePath)
	if err != nil {
		log.Fatalf("Error when calling `PatchResource`: %v", err)
	}
	return resp, err
}
