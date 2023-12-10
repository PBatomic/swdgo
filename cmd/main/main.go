package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	projectName   string
	projectPath   string
	useTailwind   bool
	generateGoMod bool
)

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

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	initialize()
}

func prepareProjectFolder(projectPath string) error {
	mode := os.FileMode(0755)
	os.MkdirAll(projectPath, mode)
	foldersToCreate := []string{"/cmd/main", "/web/template", "/web/public/css", "/web/public/assets", "/internal/server", "/internal/routes"}
	filesToCreate := []string{"cmd/main/main.go", "web/template/index.html", "web/public/css/style.css", "web/main.css", "internal/server/server.go", "internal/routes/route.go"}
}
