// project.go
package main

import (
	"fmt"
	"os"
	"strings"
)

// CreateProjectStructure generates directories, files, and minimal configuration.
// Called by main.go after collecting user inputs.
func CreateProjectStructure(projectName, useTypeScript, installMotion, wrapperChoice string) error {
	// 1. Create the main project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return err
	}

	// 2. Change working directory to the newly created project
	if err := os.Chdir(projectName); err != nil {
		return err
	}

	// 3. Create additional sub-directories (src, public, etc.)
	subDirs := []string{"src", "src/wrappers", "public"}
	for _, dir := range subDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// 4. Create a package.json
	if err := createPackageJSON(useTypeScript, installMotion); err != nil {
		return err
	}

	// 5. Create a tsconfig.json if TypeScript was chosen
	if useTypeScript == "yes" {
		if err := createTSConfig(); err != nil {
			return err
		}
	}

	// 6. Download the correct wrapper files (React or Vue)
	if err := SetupWrapper(wrapperChoice); err != nil {
		return err
	}

	// 7. Create a basic README
	if err := createReadme(projectName); err != nil {
		return err
	}

	// 8. Print instructions for next steps
	fmt.Println(ColorGreen + "Project directories created." + ColorReset)
	fmt.Println(ColorCyan + "Next steps:" + ColorReset)
	fmt.Println("1) cd " + projectName)
	fmt.Println("2) Install dependencies: npm install (or yarn install)")
	fmt.Println("3) Start building your fancy Motion project!")

	return nil
}

// createPackageJSON creates a minimal package.json file
func createPackageJSON(useTypeScript, installMotion string) error {
	fileContent := `{
  "name": "fancy-motion-app",
  "version": "1.0.0",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "start": "vite preview"
  },
  "dependencies": {
    "react": "^18.0.0",
    "react-dom": "^18.0.0"
  },
  "devDependencies": {
    "vite": "^4.0.0"
  }
}`

	// If the user wants framer-motion, add it to dependencies
	if installMotion == "yes" {
		fileContent = strings.Replace(fileContent, `"react-dom": "^18.0.0"`, `"react-dom": "^18.0.0",
    "framer-motion": "^10.0.0"`, 1)
	}

	// If user wants TypeScript, add TypeScript and type definitions
	if useTypeScript == "yes" {
		fileContent = strings.Replace(fileContent, `"vite": "^4.0.0"`, `"vite": "^4.0.0",
    "typescript": "^4.8.4",
    "@types/react": "^18.0.0",
    "@types/react-dom": "^18.0.0"`, 1)
	}

	return writeFile("package.json", fileContent)
}

// createTSConfig writes out a minimal tsconfig.json
func createTSConfig() error {
	tsConfig := `{
  "compilerOptions": {
    "target": "ESNext",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "strict": true,
    "jsx": "react-jsx",
    "esModuleInterop": true,
    "skipLibCheck": true,
    "allowSyntheticDefaultImports": true
  },
  "include": ["src"]
}`
	return writeFile("tsconfig.json", tsConfig)
}

// createReadme creates a minimal README.md
func createReadme(projectName string) error {
	readmeContent := fmt.Sprintf(`# %s

This is a fancy new Motion project built with React (or Vue), optionally including framer-motion!

## Getting Started

1. \`npm install\`
2. \`npm run dev\`

Enjoy!
`, projectName)
	return writeFile("README.md", readmeContent)
}

// createReadme creates a minimal README.md without causing syntax errors.
func createReadme(projectName string) error {
    // We’ll build the README content in a normal string
    readmeContent := fmt.Sprintf("# %s\n\n", projectName) +
        "This is a fancy new Motion project built with React (or Vue), optionally including framer-motion!\n\n" +
        "## Getting Started\n\n" +
        "1. `npm install`\n" +
        "2. `npm run dev`\n\n" +
        "Enjoy!\n"

    // Then write it out to README.md
    return writeFile("README.md", readmeContent)
}
