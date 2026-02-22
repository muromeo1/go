package copilotusage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"
)

type Result struct {
	Message        string         `json:"message"`
	Status         string         `json:"status"`
	QuotaSnapshots QuotaSnapshots `json:"quota_snapshots"`
}

type QuotaSnapshots struct {
	PremiumInteractions PremiumInteractions `json:"premium_interactions"`
}

type PremiumInteractions struct {
	PercentRemaining float64 `json:"percent_remaining"`
}

func FetchUsage(url, token string) string {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest:", err.Error())
	}

	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("Editor-Version", "vscode/1.107.0")
	req.Header.Add("Editor-Plugin-Version", "copilot-chat/0.35.0")
	req.Header.Add("User-Agent", "GitHubCopilotChat/0.35.0")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Do:", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ReadAll:", err.Error())
	}

	result := &Result{}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Fatalln("Unmarshal:", err.Error())
	}

	if result.Status != "" {
		message := "error"

		if result.Message != "" {
			message = result.Message
		}

		return fmt.Sprintf("%s: %s", result.Status, message)
	}

	remaining := round(100 - result.QuotaSnapshots.PremiumInteractions.PercentRemaining)

	return fmt.Sprintf("%.1f%%", remaining)
}

func round(value float64) float64 {
	multiplier := math.Pow(10, float64(1))
	return math.Round(value*multiplier) / multiplier
}
