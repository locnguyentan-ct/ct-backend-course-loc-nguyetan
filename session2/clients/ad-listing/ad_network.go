package ad_listing

import (
	"log"
	"net/http"
	"os"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)

type ClientOption func(*client)

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func WithRetryTimes(retryTimes int) ClientOption {
	return func(c *client) {
		c.retryTimes = retryTimes
	}
}

func WithLogger(logger *log.Logger) ClientOption {
	return func(c *client) {
		c.logger = logger
	}
}

func WithLoggerToFile(filename string) ClientOption {
	return func(c *client) {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Failed to open log file:", err)
		}
		c.logger = log.New(file, "", log.LstdFlags)
	}
}

func NewClient(baseUrl string, options ...ClientOption) *client {
	c := &client{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		retryTimes: 3,
		logger:     nil,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}
