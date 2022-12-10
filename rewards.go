package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main2() {
	client := &http.Client{}

	data, _ := json.Marshal([]map[string]any{{
		"operationName": "RedeemCustomReward",
		"extensions": map[string]any{"persistedQuery": map[string]any{
			"version": 1, "sha256Hash": "d56249a7adb4978898ea3412e196688d4ac3cea1c0c2dfd65561d229ea5dcc42",
		}},
		"variables": map[string]any{"input": map[string]any{
			"channelID": "42412771",
			"cost":      100,
			"prompt":    "Я вспоминаю про правильную осанку. Спасибо!",
			"rewardID":  "0b2ae804-0c15-4901-9817-55ff6dcf53a1",
			"title":     "Ровнее спинку!",
			// channelID":     "156632065",
			// "cost":          1,
			// "prompt":        nil,
			// "rewardID":      "cdaf337e-7eaf-4adb-a4ef-08e476d4d649",
			// "textInput":     time.Now().Format(time.ANSIC),
			// "title":         "test",
			"transactionID": uuid.New().String(),
		}},
	}})
	req, err := http.NewRequest("POST", "https://gql.twitch.tv/gql#origin=twilight", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://www.twitch.tv/")
	req.Header.Set("Authorization", "OAuth xdd")
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Set("Origin", "https://www.twitch.tv")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
