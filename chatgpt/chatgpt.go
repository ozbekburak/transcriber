package chatgpt

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
	"github.com/rakyll/openai-go/completion"
	"github.com/rakyll/openai-go/whisper"
)

func TranscribeAudio(audioFile string) (string, error) {
	sesh := openai.NewSession(os.Getenv("OPENAI_API_KEY"))
	wc := whisper.NewClient(sesh, "")

	if audioFile == "" {
		log.Println("You must specify an audio file to transcribe")
	}

	log.Printf("Audio file to transcribe: %s", audioFile)
	log.Println("Transcribing audio...")
	f, err := os.Open(audioFile)
	if err != nil {
		log.Printf("Failed to open audio file error: %v", err)
	}
	defer f.Close()
	resp, err := wc.Transcribe(context.TODO(), &whisper.CreateCompletionParams{
		Language:    "en",
		Audio:       f,
		AudioFormat: "mp3",
	})
	if err != nil {
		log.Fatalf("Failed to transcribe audio error: %v", err)
	}

	log.Println("Transcibe completed...")
	return resp.Text, nil
}

func AskChatGPT(prompt string) ([]string, error) {
	client := chat.NewClient(openai.NewSession(os.Getenv("OPENAI_API_KEY")), "gpt-3.5-turbo")
	resp, err := client.CreateCompletion(context.Background(), &chat.CreateCompletionParams{
		Messages: []*chat.Message{
			{
				Role:    "user",
				Content: prompt},
		},
	})
	if err != nil {
		// If we exceeded the limit, try to ask davinci model
		log.Printf("Failed to create openai gpt-3.5-turbo completion error: %v", err)
		if strings.Contains(err.Error(), "status_code=429") {
			log.Println("Exceeded the limit, trying to ask davinci model...")
			answers, err := askDavinci([]string{prompt})
			if err != nil {
				log.Printf("Failed to ask chat gpt with model davinci-text-003 error: %v", err)
				return nil, err
			}
			var keywordList []string
			for _, answer := range answers {
				keywordList = append(keywordList, answer.Text)
			}

			return keywordList, nil
		}
		return nil, err
	}

	var contents []string
	for _, choice := range resp.Choices {
		contents = append(contents, choice.Message.Content)
	}

	return contents, nil
}

func askDavinci(prompt []string) ([]*completion.Choice, error) {
	client := completion.NewClient(openai.NewSession(os.Getenv("OPENAI_API_KEY")), "text-davinci-003")
	resp, err := client.Create(context.Background(), &completion.CreateParams{
		N:         1,
		MaxTokens: 200,
		Prompt:    prompt})
	if err != nil {
		log.Printf("Failed to create openai completion error with model of davinci-text-003: %v", err)
		return nil, err
	}

	return resp.Choices, nil
}
