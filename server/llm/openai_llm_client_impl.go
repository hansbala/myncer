package llm

import (
	"context"

	"github.com/hansbala/myncer/core"
)

func NewOpenAILlmClient() core.LlmClient {
	return &openAILlmClientImpl{}
}

type openAILlmClientImpl struct{}

var _ core.LlmClient = (*openAILlmClientImpl)(nil)

func (o *openAILlmClientImpl) GetResponse(
	ctx context.Context,
	systemPrompt string,
	userPrompt string,
) (string, error) {
	return "", core.NewError("open ai llm client not implemented")
}
