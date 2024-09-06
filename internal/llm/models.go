package llm

// "model": "llama3.1",
//   "prompt": "Who am I",
//   "stream": false,
// 	 "context":[128009, ...]
// }

type LLMRequest struct {
	Model   string  `json:"model"`
	Prompt  string  `json:"prompt"`
	Context []int64 `json:"context,omitempty"`
	Stream  bool    `json:"stream"`
}

// Reference:
// https://github.com/ollama/ollama/blob/main/docs/api.md

type LLMResponse struct {
	Model              string  `json:"model"`
	CreatedAt          string  `json:"created_at"`
	Response           string  `json:"response"`
	Done               bool    `json:"done"`
	DoneReason         string  `json:"done_reason"`
	Context            []int64 `json:"context"`
	TotalDuration      int64   `json:"total_duration"`
	LoadDuration       int64   `json:"load_duration"`
	PromptEvalCount    int     `json:"prompt_eval_count"`
	PromptEvalDuration int64   `json:"prompt_eval_duration"`
	EvalCount          int     `json:"eval_count"`
	EvalDuration       int64   `json:"eval_duration"`
}

// {
// 	"model": "llama3.1",
// 	"created_at": "2024-09-05T15:52:03.79126Z",
// 	"response": "You told me earlier that your name is ...",
// 	"done": true,
// 	"done_reason": "stop",
// 	"context": [
// 	  128009,
// 	  128006,
//    ...
// 	  17219,
// 	  30
// 	],
// 	"total_duration": 7215912208,
// 	"load_duration": 5860136958,
// 	"prompt_eval_count": 50,
// 	"prompt_eval_duration": 267205000,
// 	"eval_count": 30,
// 	"eval_duration": 1079129000
//   }
