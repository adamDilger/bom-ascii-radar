package bom

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func FetchImageNames(productCode string) ([]string, error) {
	url := fmt.Sprintf("http://www.bom.gov.au/products/%s.loop.shtml", productCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for url [%s]: %v", url, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch immage for url [%s]: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image names for product code [%s]: %v", productCode, resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for url [%s]: %v", url, err)
	}

	bodyString := string(bodyBytes)

	r, _ := regexp.Compile("theImageNames\\[\\d+\\] = \"(.*?)\"")

	matches := []string{}
	for _, m := range r.FindAllStringSubmatch(bodyString, 10) {
		matches = append(matches, m[1])
	}

	return matches, nil
}
