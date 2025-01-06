package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func NewOpenAiClient() *openai.Client {
	token := os.Getenv("QWEN_API_KEY")
	baseUrl := os.Getenv("QWEN_BASE_URL")

	config := openai.DefaultConfig(token)
	config.BaseURL = baseUrl

	return openai.NewClientWithConfig(config)
}

// Chat 对话
func Chat(message []openai.ChatCompletionMessage) openai.ChatCompletionMessage {
	client := NewOpenAiClient()
	response, err := client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model:    ModelQwenPlus,
		Messages: message,
	})

	if err != nil {
		log.Println(err)
		return openai.ChatCompletionMessage{}
	}

	return response.Choices[0].Message
}

func ToolChat(message []openai.ChatCompletionMessage, tools []openai.Tool) openai.ChatCompletionMessage {
	client := NewOpenAiClient()
	response, err := client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model:      ModelQwenTurbo,
		Messages:   message,
		Tools:      tools,
		ToolChoice: ToolChoiceAuto,
	})

	if err != nil {
		log.Println(err)
		return openai.ChatCompletionMessage{}
	}
	return response.Choices[0].Message
}