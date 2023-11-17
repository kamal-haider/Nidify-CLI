package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createDirectory(dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Printf("Directory '%s' created successfully.\n", dirPath)
}

func createFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	fmt.Printf("File '%s' created successfully.\n", filePath)
}

func processEntry(currentPath string, entry map[string]interface{}, featureName string) {
	for key, value := range entry {
		targetPath := filepath.Join(currentPath, key)
		createDirectory(targetPath)

		switch v := value.(type) {
		case []interface{}:
			for _, element := range v {
				switch e := element.(type) {
				case string:
					filePath := filepath.Join(targetPath, strings.ReplaceAll(e, "${feature}", featureName))
					createFile(filePath)
				case map[string]interface{}:
					createDirectory(targetPath)
					processEntry(targetPath, e, featureName)
				}
			}
		case map[string]interface{}:
			processEntry(targetPath, v, featureName)
		}
	}
}

func createStructureFromJsonAndName(featureName string, jsonBlob []byte) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(jsonBlob, &jsonData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	createDirectory(featureName)
	processEntry(featureName, jsonData, featureName)
	fmt.Printf("Structure in '%s' created successfully.\n", featureName)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the feature name and template file location as command-line arguments.")
		return
	}

	featureName := os.Args[1]
	templateFilePath := os.Args[2]

	jsonBlob, err := os.ReadFile(templateFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	createStructureFromJsonAndName(featureName, jsonBlob)
}
