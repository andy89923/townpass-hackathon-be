package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultModel string = "ollama3.1"
	LLMBaseURI   string = "http:%+v//localhost:%+v11434"
	GENERATEPATH string = "/api/generate"
	PrePrompt    string = "Give you a json type and with a paragraphs. " +
		"Try parsing at most attributes as you can, and return json type only\n" +
		"Json Definition:%+v" +
		"Paragraphs:%+v"
)

func GetLLMTargetJson(plainText string, targetJson interface{}) error {

	llmReq := LLMRequest{
		Model:  DefaultModel,
		Prompt: PrePrompt + plainText,
		Stream: false,
	}

	jsonPayload, err := json.Marshal(llmReq)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON:%+v", err)
	}
	req, err := http.NewRequest("POST", LLMBaseURI+GENERATEPATH, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("Error creating request:%+v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request:%+v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body:%+v", err)
	}

	var llmResp LLMResponse
	err = json.Unmarshal(body, &llmResp)
	if err != nil {
		return fmt.Errorf("Unmarshal LLMResponse error:%+v", err)
	}

	err = json.Unmarshal(body, &targetJson)
	if err != nil {
		return fmt.Errorf("Unmarshal LLMResponse error:%+v", err)
	}
	return nil
}
