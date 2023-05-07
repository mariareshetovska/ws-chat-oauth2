# Websocket Chat App with Authentication and SSL Certificate

This is a simple chat application built in Golang that uses WebSockets for real-time communication and Google and Linkedin authentication for user authentication. It also includes an SSL certificate for secure communication.

## How to Use
1. Clone this repository

2. Install Golang if it is not already installed.

3. Create a .env file in the root directory of the project and add the following variables:
```env
GOOGLE_CLIENT_ID=<your-google-client-id>
GOOGLE_CLIENT_SECRET=<your-google-client-secret>
GOOGLE_REDIRECT_URL=<your-google-redirect-url>

LINKEDIN_CLIENT_ID=<your-linkedin-client-id>
LINKEDIN_CLIENT_SECRET=<your-linkedin-client-secret>
LINKEDIN_REDIRECT_URL=<your-linkedin-redirect-url>
```
5. Generate an SSL certificate by running the following command in the root directory of the project:
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout go-server.key -out go-server.crt
```


Also, to download and install all dependencies required by the application, simply run the deploy.sh file inside the root directory:
```
sh deploy.sh
```
And then go to .env file and add variables. After that to build and start app:
```
go build -o wsapp ./cmd
./wsapp
```

5. Navigate to https://your_url in your web browser to see the chat application.

6. Click on "Google Log In" or "LinkedIn Log In" to authenticate using OAuth2.

7. After successful authentication, you will be redirected to the chat page where you can start chatting
