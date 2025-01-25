package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var allDogecoinTransactions []map[string]interface{}

// fetchTransactions fetches transactions from the API for a given offset
func fetchDogecoinTransactions(offset int) ([]map[string]interface{}, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.blockchair.com/dogecoin/transactions?s=time(desc)&limit=100&offset=%d", offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", "https://blockchair.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Host", "api.blockchair.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15")
	req.Header.Set("Referer", "https://blockchair.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response using a map
	var apiResp map[string]interface{}
	if err := json.Unmarshal(bodyText, &apiResp); err != nil {
		return nil, err
	}

	// Extract data array from the response
	data, ok := apiResp["data"].([]interface{})
	if !ok {
		fmt.Println(apiResp["data"])
		return nil, fmt.Errorf("failed to parse data array")
	}

	// Convert to slice of maps
	transactions := make([]map[string]interface{}, len(data))
	for i, item := range data {
		transactions[i] = item.(map[string]interface{})
	}
	return transactions, nil
}

func RunDogecoinEvent() {
	for offset := 0; offset <= 3000; offset += 100 {
		transactions, err := fetchDogecoinTransactions(offset)
		if err != nil {
			log.Fatal(err)
		}
		allDogecoinTransactions = append(allDogecoinTransactions, transactions...)

		time.Sleep(3 * time.Second)
	}

	// // Filter transactions based on output_total_usd and time
	// for _, transaction := range allDogecoinTransactions {
	// 	outputTotalUSD, ok := transaction["output_total_usd"].(float64)
	// 	if ok && outputTotalUSD > 1_000_000 {
	// 		timeStr, ok := transaction["time"].(string)
	// 		if ok && isWithinLastDay(timeStr) {
	// 			fmt.Printf("Transaction with output_total_usd > 1,000,000 within last day: %s at %s\n", formatNumber(outputTotalUSD), timeStr)
	// 		}
	// 	}
	// }
}
