package core

import "context"

// LlmClient is a generic interface for interacting with Large Language Model(s).
type LlmClient interface {
	GetResponse(ctx context.Context, systemPrompt string, userPrompt string) (string, error)
}
