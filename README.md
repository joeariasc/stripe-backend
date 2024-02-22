# Go Backend for Flutter Stripe SDK Example

This backend service is designed to work with the Flutter app demonstrating Stripe SDK integration. It handles payment processing securely and efficiently. The service is written in Go and can be exposed via ngrok for development and testing purposes.

## Prerequisites

- **Go**: Ensure you have Go installed on your machine. If you haven't installed Go, you can download it from [the official Go website](https://golang.org/dl/). Follow the installation instructions for your operating system.
- **ngrok**: ngrok is used to expose your local server to the internet. If you haven't installed ngrok, download it from [ngrok's official website](https://ngrok.com/download) and follow the setup instructions.

## Setting Up the Project

### Clone the Repository

Clone the backend repository to your local machine using:

```bash
git clone git@github.com:joeariasc/stripe-backend.git
```

#### Navigate into the project directory:

```bash
cd stripe-backend
```

#### Download Dependencies

```bash
go mod tidy
```

#### Running the Backend
In the project directory, start the backend server by running:

```bash
go run main.go
```

Replace main.go with the appropriate entry file if your project uses a different name. This command compiles and runs the Go program, starting your backend server on a default port (usually http://localhost:8080 unless specified otherwise in your code).

### Exposing the Backend with ngrok
To expose your local backend server to the internet, use ngrok. Open a new terminal window and run:

```bash
ngrok http 8080
```

Replace 8080 with the port number your Go backend server is running on if it's different. ngrok will provide you with a URL that forwards to your local server. Use this URL in your Flutter app's .env file or wherever the backend URL is configured, to connect the Flutter app to your local backend server through ngrok.

### Health Check

```
https://<random-id>.ngrok.io/api/test
```