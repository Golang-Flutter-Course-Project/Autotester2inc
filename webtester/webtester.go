package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type InputData struct {
	URL   string   `json:"url"`
	Tests []string `json:"tests"`
}

type Result struct {
	Test   string `json:"test"`
	Result bool   `json:"result"`
}

type OpenAIRequest struct {
	Model    string      `json:"model"`
	Messages []ChatInput `json:"messages"`
}

type ChatInput struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message ChatInput `json:"message"`
	} `json:"choices"`
}

func main() {
	http.HandleFunc("/run", runHandler)
	log.Println("Server started at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var input InputData
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	results, err := checkPage(input.URL, input.Tests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем результаты в Go-сервис
	err = sendResultsToGoAPI(results)
	if err != nil {
		http.Error(w, "Failed to send results: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respBytes, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func sendResultsToGoAPI(results []Result) error {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(results).
		Post("http://go-api:8081/api/results")

	if err != nil {
		return err
	}
	if resp.StatusCode() >= 400 {
		return fmt.Errorf("go API responded with status %d: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

func checkPage(url string, prompts []string) ([]Result, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}

	results := []Result{}
	for _, criterion := range prompts {
		result := checkWithRules(doc, criterion)
		if result == nil {
			context := getRelevantText(doc)
			answer, err := askAI(criterion, context)
			if err != nil {
				res := false
				result = &res
			} else {
				res := strings.HasPrefix(strings.ToLower(answer), "yes") || strings.Contains(strings.ToLower(answer), "yes")
				result = &res
			}
		}
		results = append(results, Result{Test: criterion, Result: *result})
	}

	return results, nil
}

func checkWithRules(doc *goquery.Document, criterion string) *bool {
	c := strings.ToLower(criterion)
	if strings.Contains(c, "login") {
		found := doc.Find("input[type='password']").Length() > 0
		return &found
	}
	if strings.Contains(c, "submit") || strings.Contains(c, "button") {
		found := doc.Find("button").Length() > 0 || doc.Find("input[type='submit']").Length() > 0
		return &found
	}
	if strings.Contains(c, "header") && strings.Contains(c, "welcome") {
		found := false
		doc.Find("h1,h2").Each(func(i int, s *goquery.Selection) {
			if strings.Contains(strings.ToLower(s.Text()), "welcome") {
				found = true
			}
		})
		return &found
	}
	return nil
}

func getRelevantText(doc *goquery.Document) string {
	var bestText string
	doc.Find("main, div").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if len(text) > len(bestText) {
			bestText = text
		}
	})
	if len(bestText) > 1000 {
		return bestText[:1000]
	}
	return bestText
}

func askAI(criterion, context string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		apiKey = "sk-0M12tbkrnKubF86nHyKPKidkqoqfNzei"
	}

	client := resty.New()
	client.SetTimeout(10 * time.Second)
	req := OpenAIRequest{
		Model: "gpt-4.1-2025-04-14",
		Messages: []ChatInput{
			{Role: "user", Content: fmt.Sprintf("Yes or No: %s? Context: %s", criterion, context)},
		},
	}

	var resp OpenAIResponse
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&resp).
		Post("https://api.proxyapi.ru/openai/v1/chat/completions")

	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}
	return resp.Choices[0].Message.Content, nil
}
