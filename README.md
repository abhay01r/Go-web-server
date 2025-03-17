# A-simple-Go-web-server
This is a basic web server built using Go's net/http package. It listens on port 8080 and serves a simple welcome message when accessed via a web browser or API request.

**How It Works**

The server runs on localhost:8080.

When a user visits http://localhost:8080/, it calls the home function.

The home function responds with the message: "Welcome to my server".

**Prerequisites**

Install Go 

Create a go.mod file using:

**Installation & Running**

go mod init mywebserver

Run the server:

go run main.go

Open your browser and go to:

http://localhost:8080/

You should see:

Welcome to my server

