package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ANSI color codes for fancy output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

// ASCII art banner
const banner = `
  __  __       _   _
 |  \/  |     | | (_)
 | \  / | ___ | |_ _  ___ ___
 | |\/| |/ _ \| __| |/ __/ _ \
 | |  | | (_) | |_| | (_|  __/
 |_|  |_|\___/ \__|_|\___\___|

   ~ Fancy Motion Project Setup ~
`

func main() {
	printBanner()

	reader := bufio.NewReader(os.Stdin)

	// Ask for the project name
	projectName := Prompt(reader, ColorYellow+"Enter your project name:"+ColorReset, "my-motion-app")

	// Ask if user wants TypeScript
	useTypeScript := Prompt(reader, ColorYellow+"Use TypeScript? (yes/no):"+ColorReset, "yes")

	// Ask if user wants to install framer-motion
	installMotion := Prompt(reader, ColorYellow+"Install 'framer-motion'? (yes/no):"+ColorReset, "yes")

	// Ask which wrapper user wants: React or Vue
	wrapperChoice := Prompt(reader, ColorYellow+"Which wrapper do you want to use? (react/vue):"+ColorReset, "react")

	fmt.Println()
	fmt.Println(ColorCyan + "Creating project structure..." + ColorReset)
	err := CreateProjectStructure(projectName, useTypeScript, installMotion, wrapperChoice)
	if err != nil {
		log.Fatal(ColorRed + "Error creating project structure: " + err.Error() + ColorReset)
	}

	fmt.Println(ColorGreen + "Done! Happy coding!" + ColorReset)
}

// printBanner prints a fancy ASCII banner
func printBanner() {
	fmt.Println(ColorCyan + banner + ColorReset)
}
