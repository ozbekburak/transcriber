package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ozbekburak/transcriber/chatgpt"
)

// TODO: Make audio file data type

var audioDir = "audio"

func main() {
	// Iterate through audio dir to transcribe
	audioFiles, err := os.ReadDir(audioDir)
	if err != nil {
		log.Printf("Failed to get audio files in directory of '%s' error: %v", audioDir, err)
		return
	}

	for _, audioFile := range audioFiles {
		path := filepath.Join(audioDir, audioFile.Name())

		text, err := startTranscription(path)
		if err != nil {
			log.Printf("Failed to transcribe audio file '%s' error: %v", path, err)
			continue
		}
		log.Println("Transcription completed successfully!")

		if err = startTranslation(text); err != nil {
			log.Printf("Failed to translate the transcribed text error: %v", err)
			continue
		}
		log.Println("Translation completed successfully!")
		fmt.Println("-----------------------------------")
	}
}

func startTranscription(audioFile string) (string, error) {
	transcribedText, err := chatgpt.TranscribeAudio(audioFile)
	if err != nil {
		log.Printf("Failed to transcribe audio error: %v", err)
	}
	log.Println("We got the transcribed text, preparing to save the transcript...")

	// Save the transcript
	transcriptFile, err := save(transcribedText, "transcripts", "en", "transcript")
	if err != nil {
		log.Printf("Failed to save transcript file error: %v", err)
		return "", err
	}
	log.Printf("Transcript saved to %s", transcriptFile.Name())

	return transcribedText, nil
}

func startTranslation(text string) error {
	prompt := "Please translate the following text into "
	scanner := bufio.NewScanner(os.Stdin)
	// Ask the user to which language to translate
	fmt.Print("Which language you want to translate the audio: ")
	scanner.Scan()
	translationLang := scanner.Text()

	prompt += translationLang + ":\n" + text

	answers, err := chatgpt.AskChatGPT(prompt)
	if err != nil {
		log.Printf("Failed to translate the transcribed text error: %v", err)
		return err
	}
	log.Println("We got the translated text, preparing to save the translation...")

	// Save the translation
	translationFile, err := save(answers[0], "translations", translationLang, "translation")
	if err != nil {
		log.Printf("Failed to save translation file error: %v", err)
		return err
	}
	log.Printf("Translation saved to %s", translationFile.Name())

	return nil

}

func save(text, dir, lang, dataType string) (*os.File, error) {
	// fileName: transcript-english-20210101120000.txt
	timestamp := time.Now().Format("20060102150405")
	fileName := filepath.Join(dir, fmt.Sprintf("%s-%s-%s.txt", dataType, lang, timestamp))

	// Open a file for writing, create it if it doesn't exist, and append to it
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open file error: %v", err)
		return nil, err
	}
	defer file.Close()

	// Create a bufio.Writer for the file
	writer := bufio.NewWriter(file)

	words := strings.Split(text, " ")
	lineLength := 150
	line := ""
	for _, word := range words {
		if len(line)+len(word)+1 > lineLength {
			// Write the current line to the file and start a new one
			fmt.Fprintln(writer, line)
			line = word
		} else {
			// Add the word to the current line
			if line == "" {
				line = word
			} else {
				line += " " + word
			}
		}
	}
	// Write the last line to the file
	fmt.Fprintln(writer, line)

	// Flush the buffer to ensure that all data is written to the file
	writer.Flush()

	// if _, err = file.WriteString(data); err != nil {
	// 	log.Printf("Failed to write to file error: %v", err)
	// 	return nil, err
	// }

	return file, nil
}
