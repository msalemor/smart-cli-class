package pkg

type Command struct {
	Command     string   `json:"command"`
	Args        []string `json:"args"`
	Explanation string   `json:"explanation"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type FormatType struct {
	Type string `json:"type"`
}

type OpenAIRequest struct {
	Messages       []Message  `json:"messages"`
	Model          string     `json:"model"`
	Temperature    float64    `json:"temperature"`
	ResponseFormat FormatType `json:"response_format"`
}

type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Message      Message `json:"message"`
		Index        int     `json:"index"`
		Logprobs     string  `json:"logprobs"`
		FinishReason string  `json:"finish_reason"`
		Delta        Message `json:"delta"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}
