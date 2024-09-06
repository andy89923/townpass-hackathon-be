package llm_test

import (
	"encoding/json"
	"go-cleanarch/internal/llm"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLLMRequest(t *testing.T) {

	req := llm.LLMRequest{
		Model:  "ollama3.1",
		Prompt: "Why is the sky blue?",
		Stream: false,
	}

	// var data []byte
	data, err := json.Marshal(req)
	require.NoError(t, err)

	t.Log(data)

	var actReq llm.LLMRequest
	err = json.Unmarshal(data, &actReq)
	require.NoError(t, err)

	t.Log(actReq)
}
