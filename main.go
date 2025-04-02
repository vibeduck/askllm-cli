package main

import (
	"fmt"
	"os"

	"github.com/kaji/askllm-cli/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	openaiAPIKey string
	rootCmd      = &cobra.Command{
		Use:   "askllm [question]",
		Short: "Ask questions to an OpenAI-compatible API from the command line",
		Long: `askllm is a CLI tool that allows you to ask questions to an OpenAI-compatible API
directly from the command line. You can provide your API key either through the
OPENAI_API_KEY environment variable or using the --openai-api-key flag.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			question := args[0]
			return askQuestion(question)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&openaiAPIKey, "openai-api-key", "", "OpenAI API key")
	viper.BindEnv("openai-api-key", "OPENAI_API_KEY")
	viper.BindPFlag("openai-api-key", rootCmd.PersistentFlags().Lookup("openai-api-key"))
}

func askQuestion(question string) error {
	apiKey := viper.GetString("openai-api-key")
	if apiKey == "" {
		return fmt.Errorf("API key not provided. Please set OPENAI_API_KEY environment variable or use --openai-api-key flag")
	}

	client := api.NewClient(apiKey)
	response, err := client.Ask(question)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	fmt.Println(response)
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
