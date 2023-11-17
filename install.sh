#!/bin/bash

# Build the Go application
go build -o nidify

# Move the executable to the appropriate directory
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
    # Unix-like system
    sudo mv nidify /usr/local/bin
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    # Windows system
    echo "Please add the directory containing the 'nidify' executable to your system's PATH environment variable."
else
    echo "Unknown operating system. Please move the 'nidify' executable to a directory in your system's PATH."
fi