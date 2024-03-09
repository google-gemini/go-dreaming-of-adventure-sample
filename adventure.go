// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/status"
)

const systemInstructionsFile = "system-instructions.md"

var sleepTime = struct {
	character time.Duration
	sentence  time.Duration
}{
	character: time.Millisecond * 30,
	sentence:  time.Millisecond * 300,
}

// Streaming output column position.
var col = 0

// getBytes returns the file contents as bytes.
func getBytes(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file bytes %v: %v\n", path, err)
	}
	return bytes
}

// newClient creates a new API client using API_KEY environment variable.
func newClient(ctx context.Context) *genai.Client {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatalf("Environment variable API_KEY is not set.\n" +
			"To obtain an API key, visit https://aistudio.google.com/, select 'Get API key'.\n")
	}

	// New client, using API key authorization.
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}
	return client
}

func main() {
	ctx := context.Background()
	client := newClient(ctx)
	defer client.Close()

	// Configure desired model.
	model := client.GenerativeModel("gemini-1.5-pro-latest")

	// Initialize new chat session.
	session := model.StartChat()

	// Set system instructions.
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(getBytes(systemInstructionsFile))},
		Role:  "system",
	}

	dreamQuestion := "What do you want to dream about?"

	// Establish chat history.
	session.History = []*genai.Content{{
		Role:  "model",
		Parts: []genai.Part{genai.Text(dreamQuestion)},
	}}

	printRuneFormatted('\n')
	topic := askUser(dreamQuestion)
	sendAndPrintResponse(ctx, session, topic)

	chat(ctx, session)
}

// chat is a simple chat loop.
func chat(ctx context.Context, session *genai.ChatSession) {
	for {
		fmt.Println()
		action := askUser(">>")
		resp := fmt.Sprintf("The user wrote: %v\n\nWrite the next short paragraph.", action)
		sendAndPrintResponse(ctx, session, resp)
	}
}

// askUser prompts the user for input.
func askUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		printStringFormatted(fmt.Sprintf("%v ", prompt))
		action, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v\n", err)
		}
		action = strings.TrimSpace(action)
		if (len(action)) == 0 {
			continue
		}
		return action
	}
}

// sendAndPrintResponse sends a message to the model and prints the response.
func sendAndPrintResponse(ctx context.Context, session *genai.ChatSession, text string) {
	it := session.SendMessageStream(ctx, genai.Text(text))
	printRuneFormatted('\n')
	printRuneFormatted('\n')

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			printStringFormatted("\n\nYou feel a jolt of electricity as you realize you're being unplugged from the matrix.\n\n")
			log.Printf("Error sending message: err=%v\n", err)

			var ae *apierror.APIError
			if errors.As(err, &ae) {
				log.Printf("ae.Reason(): %v\n", ae.Reason())
				log.Printf("ae.Details().Help.GetLinks(): %v\n", ae.Details().Help.GetLinks())
			}

			if s, ok := status.FromError(err); ok {
				log.Printf("s.Message: %v\n", s.Message())
				for _, d := range s.Proto().Details {
					log.Printf("- Details: %v\n", d)
				}
			}
			os.Exit(1)
		}
		for _, cand := range resp.Candidates {
			streamPartialResponse(cand.Content.Parts)
		}
	}
	printRuneFormatted('\n')
}

// streamPartialResponse prints the partial response.
func streamPartialResponse(parts []genai.Part) {
	for _, part := range parts {
		printStringFormatted(fmt.Sprintf("%v", part))
	}
}

// printStringFormatted prints the string and formats it, with delays for effect.
func printStringFormatted(text string) {
	for _, c := range text {
		printRuneFormatted(c)
	}
}

// printRuneFormatted prints the rune and formats it, with delays for effect.
func printRuneFormatted(c rune) {
	switch c {
	case '.':
		fmt.Print(string(c))
		col++
		time.Sleep(sleepTime.sentence)
	case '\n':
		fmt.Print(string(c))
		col = 0
	case ' ':
		if col == 0 {
			// Do nothing.
		} else if col > 80 {
			fmt.Print("\n")
			col = 0
		} else {
			fmt.Print(string(c))
			col++
		}
	default:
		fmt.Print(string(c))
		col++
	}
	time.Sleep(sleepTime.character)
}
