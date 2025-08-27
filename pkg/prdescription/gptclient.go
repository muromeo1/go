package prdescription

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Url   string
	Token string
	Model string
}

type Response struct {
	Id     string   `json:"id"`
	Error  string   `json:"error"`
	Output []Output `json:"output"`
}

type Output struct {
	Content []Content `json:"content"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewClient(url, token, model string) *Client {
	return &Client{
		Url:   url,
		Token: token,
		Model: model,
	}
}

func (client *Client) Responses(gitLog string) string {
	input := fmt.Sprintf(input(), gitLog)

	payload := map[string]string{
		"model":        client.Model,
		"input":        input,
		"instructions": instructions(),
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + client.Token,
	}

	http := NewHttpClient(client.Url, headers)

	log.Println("Fetching response from OpenAI...")
	resp := http.Post(payload)

	response := &Response{}

	err := json.Unmarshal(resp, response)
	if err != nil {
		log.Fatal("Unmarshal failed: ", err)
	}

	return response.Output[0].Content[0].Text
}

func instructions() string {
	return `You are an expert at creating git PR descriptions based on file changes.
	You will be provided with a git log and you will create a PR description using the context of the project.
	You should be succinct using markdown and returning only the description in the following structure:
	
	## PR Description
	
	[A brief description of the changes made in the PR]
	
	- [Bullet point 1]
	- [Bullet point 2]
	- [Bullet point 3]
	
	Ensure that the description is clear, concise, and provides enough context for reviewers to understand the purpose and scope of the changes.`
}

func input() string {
	return `Based on these file changes: %s
	create a git PR description using the context of the project.
	You should be succinct using markdown and returning only the description in the following structure
	
	## PR Description
	
	description goes here
	
	possible bullet points`
}
