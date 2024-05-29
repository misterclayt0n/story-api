package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateStory(prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDwjYUIncBk-lG_52-vNm1Acw2Z4LQ75yQ"))
	if err != nil {
		return "", err
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	res, err := cs.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	return extractContent(res), nil
}

func extractContent(resp *genai.GenerateContentResponse) string {
	var content string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				content += fmt.Sprint(part)
			}
		}
	}
	return content
}
