package ad_listing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

func (c *client) GetAdByCate(ctx context.Context, cate string) (*AdsResponse, error) {
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request - Cate %v, Duration: %v", cate, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, cate)
	fmt.Printf("URL request: %v\n", url)
	resp, err := c.httpClient.Get(url)
	for attempt := 0; attempt <= c.retryTimes; attempt++ {
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			// Retry for 5xx status codes
			c.logger.Printf("Retrying request - Cate: %v, Attempt: %d, StatusCode: %d", cate, attempt+1, resp.StatusCode)
			time.Sleep(1 * time.Second) // Add a delay before retrying

			// Close the previous response body
			resp.Body.Close()
			continue
		}

		break // Exit the loop if the request is successful or the status code is not 5xx
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nResponse %v\n", string(b))

	var adResp AdsResponse
	error := json.Unmarshal(b, &adResp)
	if error != nil {
		fmt.Println("Error unmarshaling JSON:", error)
	}

	return &adResp, err
}
