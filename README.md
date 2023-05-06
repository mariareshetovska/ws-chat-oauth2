# Websocket Chat App with Authentication

This is a simple chat application built in Golang that uses WebSockets for real-time communication and Google and Linkedin authentication for user authentication.

## How to Use
1. Clone this repository to your local machine.

2. Install Golang if it is not already installed.

3. Create a .env file in the root directory of the project and add the following variables:
```env
GOOGLE_CLIENT_ID=<your-google-client-id>
GOOGLE_CLIENT_SECRET=<your-google-client-secret>

LINKEDIN_CLIENT_ID=<your-linkedin-client-id>
LINKEDIN_CLIENT_SECRET=<your-linkedin-client-secret>
```
4. Run the application using the following command in the root directory of the project:
```go
go run cmd/main.go
```
5. Navigate to http://localhost:8080 in your web browser to see the chat application.

6. Click on "Google Log In" or "LinkedIn Log In" to authenticate using OAuth2.

7. After successful authentication, you will be redirected to the chat page where you can start chatting
