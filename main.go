package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gosimple/slug"
)

func main() {
	valueWithoutFriendlyName := promptString("Enter the value for convert to Friendly Url")
	valueWithFriendlyName := getFriendlyUrl(valueWithoutFriendlyName)
	printMessageAndValue("Friendly Url", valueWithFriendlyName)
	fmt.Println("----------------")
	fmt.Println("Some examples")
	printTests()
}

func promptString(messageToShow string) string {
	fmt.Println(messageToShow)
	return scanner()
}

func getFriendlyUrl(value string) string {
	text := slug.Make(value)
	return text
}

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func printMessageAndValue(message string, value string) {
	fmt.Println(message + " : " + value)
}

func printTests() {
	text := slug.Make("Hellö Wörld хелло ворлд")
	fmt.Println(text) // Will print: "hello-world-khello-vorld"

	someText := slug.Make("影師")
	fmt.Println(someText) // Will print: "ying-shi"

	enText := slug.MakeLang("This & that", "en")
	fmt.Println(enText) // Will print: "this-and-that"

	deText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deText) // Will print: "diese-und-dass"

	slug.Lowercase = false // Keep uppercase characters
	deUppercaseText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deUppercaseText) // Will print: "Diese-und-Dass"

	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := slug.Make("water is hot")
	fmt.Println(textSub) // Will print: "sand-is-hot"
}
