package llm

import (
	"context"

	genai "google.golang.org/genai"

	"github.com/hansbala/myncer/core"
)

func NewGeminiLlmClient() core.LlmClient {
	return &geminiLlmClientImpl{}
}

type geminiLlmClientImpl struct{}

var _ core.LlmClient = (*geminiLlmClientImpl)(nil)

func (g *geminiLlmClientImpl) GetResponse(
	ctx context.Context,
	systemPrompt string,
	userPrompt string,
) (string, error) {
	client, err := g.getClient(ctx)
	if err != nil {
		return "", core.WrappedError(err, "failed to get gemini client")
	}
	model, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-pro",
		[]*genai.Content{
			{
				Parts: []*genai.Part{
					{
						Text: systemPrompt,
					},
					{
						Text: userPrompt,
					},
				},
				Role: "user",
			},
		},
		nil, /*config*/
	)
	if err != nil {
		return "", core.WrappedError(err, "failed to get response from gemini")
	}
	return model.Text(), nil
}

func (g *geminiLlmClientImpl) getClient(ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  core.ToMyncerCtx(ctx).Config.GetLlmConfig().GetGeminiConfig().GetApiKey(),
		Backend: genai.BackendGeminiAPI,
	})
}
