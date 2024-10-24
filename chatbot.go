package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var patternsResponses = map[*regexp.Regexp]func(string) string{
	regexp.MustCompile(`(?i)hi|hello|hey|halo|hai`): func(_ string) string {
		return "Hello, how can I assist you with learning programming today?"
	},
	regexp.MustCompile(`(?i)(.*)?(feeling|kabar)`): func(_ string) string {
		return "I am doing great! How about you?"
	},
	regexp.MustCompile(`(?i)(.*)?(name|nama)`): func(_ string) string {
		return "I am a rule-based chatbot created to help you learn programming, from basic concepts to advanced topics."
	},
	regexp.MustCompile(`(?i)(.*)?(do|help|lakukan)`): func(_ string) string {
		return "I can help you provide YouTube Video. Just type \"$i want to learn <keyword>\" to get started!"
	},
	regexp.MustCompile(`(?i)(.*)?(pertanyaan|question)`): func(_ string) string {
		return "Sure, feel free to ask your programming-related question!"
	},
	regexp.MustCompile(`(?i)(.*)?(about|tentang)(.*)`): func(keyword string) string {
		return fmt.Sprintf("Do you find %s interesting?", keyword)
	},
	regexp.MustCompile(`(?i)(.*)?(terjebak|stuck)(.*)`): func(_ string) string {
		return "Feeling stuck is normal in programming. Take a deep breath and let me help you figure it out!"
	},
	regexp.MustCompile(`(?i)(.*)?(menyenangkan|fun)(.*)`): func(_ string) string {
		return "Programming can be very rewarding! What aspect of it are you enjoying the most?"
	},
	regexp.MustCompile(`(?i)(.*)?error(.*)`): func(_ string) string {
		return "Encountering an error is part of the learning process."
	},
	regexp.MustCompile(`(?i)(.*)?(sulit|susah|hard|difficult)(.*)`): func(_ string) string {
		return "Programming can be challenging at times."
	},
	regexp.MustCompile(`(?i)(.*)?menurut (saya|aku),(.*)`): func(_ string) string {
		return "Would you like to explore this topic further?"
	},
	regexp.MustCompile(`(?i)\$i want to learn (.+)`): func(keyword string) string {
		return fmt.Sprintf("Searching for '%s' YouTube video...", keyword)
	},
	regexp.MustCompile(`(?i)thank you|thanks|terimakasih|makasih`): func(_ string) string {
		return "You are welcome! If you have more questions, feel free to ask."
	},
	regexp.MustCompile(`(?i)bye|goodbye`): func(_ string) string {
		return "Goodbye! Have a great day!"
	},
}

const defaultResponse = "I'm sorry, I don't understand that. Can you please rephrase?"

// getResponse matches user input to a pattern and returns the corresponding response
func getResponse(userInput string) string {
	for pattern, responseFunc := range patternsResponses {
		if match := pattern.FindStringSubmatch(userInput); match != nil {
			// Call the function with the first capture group (or empty if none)
			if len(match) > 1 {
				return responseFunc(match[1])
			}
			return responseFunc("")
		}
	}
	return defaultResponse
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Chatbot started! Type your message below:")

	// Infinite loop to continuously take user input
	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')  // Read user input
		userInput = strings.TrimSpace(userInput) // Remove any extra newlines/spaces

		// Exit the program if the user types "exit" or "quit"
		if strings.EqualFold(userInput, "bye") || strings.EqualFold(userInput, "goodbye") {
			fmt.Println("Bot: Goodbye! Have a great day!")
			break
		}

		// Get the bot's response
		response := getResponse(userInput)
		fmt.Println("Bot:", response)
	}
}
