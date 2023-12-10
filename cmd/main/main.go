package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Flags
var (
	projectName   string
	projectPath   string
	useTailwind   bool
	generateGoMod bool
)

func main() {
	initialize()
	err := prepareProjectFolder(projectPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initialize initializes the cmdline flags
func initialize() {
	flag.StringVar(&projectName, "projectName", "", "Specify project name")
	flag.BoolVar(&useTailwind, "useTailwind", false, "Use tailwind css")
	flag.StringVar(&projectPath, "projectPath", "", "Project file path")
	flag.BoolVar(&generateGoMod, "generateGoMod", false, "Generate go mod file based on project name or github repo")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: myprogram [options]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if projectName == "" {
		fmt.Println("Project name is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if projectPath == "" {
		fmt.Println("Project path is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

// getGoRuntimeVersion returns the current go runtime version
func getGoRuntimeVersion() string {
	return strings.Split(runtime.Version(), "go")[1]
}

// prepareProjectFolder creates the project folder and subfolders
func prepareProjectFolder(projectPath string) error {
	// Main folder creation
	mode := os.FileMode(0755)
	os.MkdirAll(projectPath, mode)

	// Project subfolders
	foldersToCreate := []string{"/cmd/main", "/web/template", "/web/public/css", "/web/public/assets", "/internal/server", "/internal/routes"}
	filesToCreate := []string{"cmd/main/main.go", "web/template/index.html", "web/public/css/style.css", "web/main.css", "internal/server/server.go", "internal/routes/route.go"}

	// Create folders
	fmt.Println("Creating folders...")
	for _, folder := range foldersToCreate {
		fmt.Printf("Creating folder %s\n", folder)
		err := os.MkdirAll(projectPath+folder, mode)
		if err != nil {
			return err
		}
	}

	// Create files
	fmt.Println("Creating files...")
	for _, file := range filesToCreate {
		fmt.Printf("Creating file %s\n", file)
		_, err := os.Create(projectPath + "/" + file)
		if err != nil {
			return err
		}
	}

	// Create go.mod file if needed
	if generateGoMod {
		fmt.Println("Creating go.mod file...")
		_, err := os.Create(projectPath + "/go.mod")
		if err != nil {
			return err
		}
		// Write to go.mod file
		f, err := os.OpenFile(projectPath+"/go.mod", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.WriteString("module " + projectName + "\n\ngo " + getGoRuntimeVersion()); err != nil {
			return err
		}

	}
	// Generate tailwindcss config file
	if useTailwind {
		fmt.Println("Creating tailwind.config.js file...")
		_, err := os.Create(projectPath + "/tailwind.config.js")
		if err != nil {
			return err
		}
		// Write to tailwind.config.js file
		f, err := os.OpenFile(projectPath+"/tailwind.config.js", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.WriteString("module.exports = {\n  purge: [],\n  darkMode: false, // or 'media' or 'class'\n  theme: {\n    extend: {},\n  },\n  variants: {\n    extend: {},\n  },\n  plugins: [],\n}"); err != nil {
			return err
		}
	}

	return nil
}
