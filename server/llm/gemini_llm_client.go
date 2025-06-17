package llm

import (
	"context"

	"github.com/hansbala/myncer/core"
)

func NewGeminiLlmClient() core.LlmClient {
	return &geminiLlmClientImpl{}
}

type geminiLlmClientImpl struct{}

var _ core.LlmClient = (*geminiLlmClientImpl)(nil)

func (o *geminiLlmClientImpl) GetResponse(
	ctx context.Context,
	systemPrompt string,
	userPrompt string,
) (string, error) {
	return "", core.NewError("gemini llm client not implemented")
}
