# A-simple-Go-web-server
This is a basic web server built using Go's net/http package. It listens on port 8080 and serves a simple welcome message when accessed via a web browser or API request.

**How It Works**

1. The server runs on localhost:8080.

2. When a user visits http://localhost:8080/, it calls the home function.

3. The home function responds with the message: "Welcome to my server".

**Prerequisites**

1. Install Go 

2. Create a go.mod file using:

**Installation & Running**

1. go mod init mywebserver

2. Run the server:

   go run main.go

3. Open your browser and go to:

http://localhost:8080/

4. You should see:

Welcome to my server

