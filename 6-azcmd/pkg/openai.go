package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CallChatCompletion sends a request to the OpenAI chat completion API and returns the response
func ChatCompletion(prompt string) (*Commands, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	messages := []Message{
		{
			Role:    "system",
			Content: GetSettings().SystemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	request := OpenAIRequest{
		Messages:       messages,
		Model:          GetSettings().Model,
		ResponseFormat: FormatType{Type: "json_object"}, // this is important
		Temperature:    0.1,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, GetSettings().Endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GetSettings().APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}
	res := response.Choices[0].Message.Content
	var commands Commands
	if err := json.Unmarshal([]byte(res), &commands); err != nil {
		return nil, fmt.Errorf("error unmarshaling commands: %w", err)
	}
	if len(commands.Commands) == 0 {
		return nil, fmt.Errorf("no commands in response")
	}
	//fmt.Printf("Response: %s\n", res)
	//fmt.Printf("Commands: %v\n", commands.Commands)
	return &commands, nil

}
