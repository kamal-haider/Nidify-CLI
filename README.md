# Nidify

## Installation

To install this application, you need to have Go installed on your machine. If you don't have it installed, you can download it from [here](https://golang.org/dl/). After installing Go, clone this repository to your local machine.

Once you have cloned the repository, navigate to the project directory. There you will find a shell script named `install.sh`. This script is used to build the Go application and move the executable to the appropriate directory. Run this script by navigating into the nidify directory and typing the following command in your terminal:
```bash
bash install.sh
```

## Running the Application

To run the application, ensure that the `nidify` executable is in your system's PATH. Then, in your terminal, run the following command: 
```bash
nidify test template.json
```

## Configure the template

The JSON template should be structured as a nested object where each key represents a directory and each value can either be an array of filenames or another nested object. Filenames can include a `${feature}` tag that will be replaced with the feature name provided when running the application. Here is an example of a template:
```json
{
    "directory_name": [
        "file_name.js",
        "${feature}_feature_name.js"
    ]
}
```
