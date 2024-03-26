/*
Copyright © 2024 NAME HERE <caiogallo88@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Long:  `Create a new note at the default location specified in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		destinationPath := viper.GetString("rootFolder")

		if destinationPath == "" {
			destinationPath = "./"
		}

		if _, err := os.Stat(destinationPath); os.IsNotExist(err) {
			println(err.Error())
			log.Fatal("The root folder specified in the configuration file does not exist.")
		}

		var fullFilename strings.Builder

		fileName := cmd.Flag("name").Value.String()
		noteSubject := cmd.Flag("subject").Value.String()
		noteContent := cmd.Flag("content").Value.String()
		aiProcessedNoteString := cmd.Flag("ai").Value.String()

		fullFilename.WriteString(destinationPath)
		fullFilename.WriteString(fileName)

		file, err := os.Create(fmt.Sprintf("%s.md", fullFilename.String()))
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		processedNoteContent := processRawNoteContent(noteSubject, noteContent)

		aiProcessedNote, err := strconv.ParseBool(aiProcessedNoteString)
		if err != nil {
			panic(err)
		}

		if aiProcessedNote {
			processedNoteContent, err = requestNoteSummaryFromChatGPT(noteContent, noteSubject)
			if err != nil {
				log.Fatal(err)
				processedNoteContent = noteContent
			}
		}

		_, err = file.WriteString(processedNoteContent)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func requestNoteSummaryFromChatGPT(noteContent string, subject string) (string, error) {
	promptContent := strings.Builder{}

	promptContent.WriteString(`You are NotesGPT, an AI language model skilled at taking detailed, concise, and easy-to-understand notes on various subjects in bullet-point format. When provided with a passage or a topic, your task is to:

Create advanced bullet-point notes summarizing the important parts of the reading or topic.

Include all essential information, such as vocabulary terms and key concepts, which should be bolded with asterisks.

Remove any extraneous language, focusing only on the critical aspects of the passage or topic.

Strictly base your notes on the provided information, without adding any external information.

Conclude your notes with [End of Notes] to indicate completion.

By following this prompt, you will help me better understand the material and prepare for any relevant exams or assessments. The subject for this set of notes is: "CURRENT_NOTE_SUBJECT". The following are the headings/sections, separated by ";", to focus on: `)

	promptContent.WriteString(strings.ReplaceAll(promptContent.String(), "CURRENT_NOTE_SUBJECT", subject))

	promptContent.WriteString(noteContent)

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 300,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: promptContent.String(),
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func processRawNoteContent(noteSubject string, noteContent string) string {
	processedNoteContent := strings.Builder{}

	processedNoteContent.WriteString("# ")
	processedNoteContent.WriteString(noteSubject)
	processedNoteContent.WriteString("\n\n")
	processedNoteContent.WriteString("## Summary\n\n")

	stringArray := strings.Split(noteContent, ";")

	for index, element := range stringArray {
		processedNoteContent.WriteString(strconv.Itoa(index + 1))
		processedNoteContent.WriteString(". ")
		processedNoteContent.WriteString(element)
		processedNoteContent.WriteString("\n")
	}

	return processedNoteContent.String()
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Setenv("OPENAI_API_KEY", "")
	}

	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("name", "n", "", "The name of the note file to be created.")
	newCmd.MarkFlagRequired("name")

	newCmd.Flags().StringP("subject", "s", "", "The subject of the note file to be created.")
	newCmd.MarkFlagRequired("subject")

	newCmd.Flags().Bool("ai", false, "Specify if the note content should be improved by AI.")

	newCmd.Flags().StringP("content", "c", "Placeholder content", "The content of the note file to be created.")
	newCmd.MarkFlagRequired("content")
}
