# Health Check App

## What is the Health Check App?
The Health Check App is a Go-based utility that monitors the health of a set of HTTP endpoints. It tests the availability of endpoints, calculates their availability percentage, and logs the results. This tool is designed to help you keep track of the health and performance of your web services.

## How It's Developed (Technical Overview)
The Health Check App is developed in Go and follows best practices, including Go modules for dependency management. It uses YAML configuration files to define endpoints, performs health checks every 15 seconds, and logs the availability percentage of each domain. It utilizes Go's standard libraries for HTTP requests and concurrent execution for efficiency.

## How to Run the Health Check App
1. Clone this repository to your local machine.
2. Make sure you have Go installed.
3. Navigate to the project's root directory in your terminal.
4. Run the app with the command: `go run main.go path/to/config.yaml`, where `path/to/config.yaml` is the path to your configuration file in YAML format.
5. The app will start testing the defined endpoints and log their availability percentages. Press `Ctrl+C` to stop the app.
