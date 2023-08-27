package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run nidify <FeatureName> <ModelName>")
		return
	}

	featureName := os.Args[1]
	featureFileName := pascalToSnake(featureName)
	modelName := os.Args[2]
	modelFileName := pascalToSnake(modelName)
	presentationLayer := "presentation"
	widgetSublayer := "widget"
	stateSublayer := "state"
	controllerSublayer := "controller"
	applicationLayer := "application"
	serviceSublayer := "service"
	domainLayer := "domain"
	modelSublayer := "model"
	dataLayer := "data"
	repositorySublayer := "repository"
	dtoSublayer := "dto"
	datasourceSublayer := "data_source"

	// Create the feature directory
	err := os.Mkdir(featureFileName, 0755)
	if err != nil {
		fmt.Println("Error creating parent directory:", err)
		return
	}

	// Create the presentation layer inside the feature directory
	presentationPath := featureFileName + "/" + presentationLayer
	err = os.Mkdir(presentationPath, 0755)
	if err != nil {
		fmt.Println("Error creating presentation directory:", err)
		return
	}

	// Create the widget sublayer inside the presentation layer
	widgetPath := presentationPath + "/" + widgetSublayer
	err = os.Mkdir(widgetPath, 0755)
	if err != nil {
		fmt.Println("Error creating widget directory:", err)
		return
	}

	// Create the widget file inside the widget sublayer
	createFile(widgetPath + "/" + featureFileName + "_page.dart")

	// Create the state sublayer inside the presentation layer
	statePath := presentationPath + "/" + stateSublayer
	err = os.Mkdir(statePath, 0755)
	if err != nil {
		fmt.Println("Error creating state directory:", err)
		return
	}

	// Create the state files inside the state sublayer
	createFile(statePath + "/" + featureFileName + "_data_state.dart")
	createFile(statePath + "/" + featureFileName + "_loading_state.dart")
	createFile(statePath + "/" + featureFileName + "_error_state.dart")

	// Create the controller sublayer inside the presentation layer
	controllerPath := presentationPath + "/" + controllerSublayer
	err = os.Mkdir(controllerPath, 0755)
	if err != nil {
		fmt.Println("Error creating controller directory:", err)
		return
	}

	// Create the controller file inside the controller sublayer
	createFile(controllerPath + "/" + featureFileName + "_controller.dart")

	// Create the application layer inside the feature directory
	applicationPath := featureFileName + "/" + applicationLayer
	err = os.Mkdir(applicationPath, 0755)
	if err != nil {
		fmt.Println("Error creating application directory:", err)
		return
	}

	// Create the service sublayer inside the application layer
	servicePath := applicationPath + "/" + serviceSublayer
	err = os.Mkdir(servicePath, 0755)
	if err != nil {
		fmt.Println("Error creating service directory:", err)
		return
	}

	// Create the service file inside the service sublayer
	createFile(servicePath + "/" + featureFileName + "_service.dart")

	// Create the domain layer inside the feature directory
	domainPath := featureFileName + "/" + domainLayer
	err = os.Mkdir(domainPath, 0755)
	if err != nil {
		fmt.Println("Error creating domain directory:", err)
		return
	}

	// Create the model sublayer inside the domain layer
	modelPath := domainPath + "/" + modelSublayer
	err = os.Mkdir(modelPath, 0755)
	if err != nil {
		fmt.Println("Error creating model directory:", err)
		return
	}

	// Create the model file inside the model sublayer
	createFile(modelPath + "/" + modelFileName + ".dart")

	// Create the data layer inside the feature directory
	dataPath := featureFileName + "/" + dataLayer
	err = os.Mkdir(dataPath, 0755)
	if err != nil {
		fmt.Println("Error creating data directory:", err)
		return
	}

	// Create the repository sublayer inside the data layer
	repositoryPath := dataPath + "/" + repositorySublayer
	err = os.Mkdir(repositoryPath, 0755)
	if err != nil {
		fmt.Println("Error creating repository directory:", err)
		return
	}

	// Create the repository file inside the repository sublayer
	createFile(repositoryPath + "/" + featureFileName + "_repository.dart")

	// Create the dto sublayer inside the repository sublayer
	dtoPath := dataPath + "/" + dtoSublayer
	err = os.Mkdir(dtoPath, 0755)
	if err != nil {
		fmt.Println("Error creating dto directory:", err)
		return
	}

	// Create the dto file inside the dto sublayer
	createFile(dtoPath + "/" + featureFileName + "_dto.dart")

	// Create the datasource sublayer inside the repository sublayer
	datasourcePath := dataPath + "/" + datasourceSublayer
	err = os.Mkdir(datasourcePath, 0755)
	if err != nil {
		fmt.Println("Error creating datasource directory:", err)
		return
	}

	// Create the data source file inside the data source sublayer
	createFile(datasourcePath + "/" + featureFileName + "_data_source.dart")

	fmt.Printf("Directory '%s' created successfully.\n", featureFileName)
}

func createPresentationLayer(path string) {
	// Create the directory
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Printf("Directory '%s' created successfully.\n", path)
}

func createDirectory(path string) {
	// Create the directory
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Printf("Directory '%s' created successfully.\n", path)
}

func createFile(path string) {
	// Create the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' created successfully.\n", path)
}

func pascalToSnake(input string) string {
	var result strings.Builder

	for i, char := range input {
		if unicode.IsUpper(char) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
