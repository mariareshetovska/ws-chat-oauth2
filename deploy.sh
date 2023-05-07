#!/bin/bash

# Install Go
wget https://golang.org/dl/go1.19.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
rm go1.19.linux-amd64.tar.gz

# Set environment variables
cat > .env << EOF
GOOGLE_CLIENT_ID=<your-google-client-id>
GOOGLE_CLIENT_SECRET=<your-google-client-secret>
GOOGLE_REDIRECT_URL=<your-google-redirect-url>

LINKEDIN_CLIENT_ID=<your-linkedin-client-id>
LINKEDIN_CLIENT_SECRET=<your-linkedin-client-secret>
LINKEDIN_REDIRECT_URL=<your-linkedin-redirect-url>
EOF

# Generate an SSL certificate by running the following command in the root directory of the project
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout go-server.key -out go-server.crt

# Exit the SSH session
exit

EOF

# Exit the script
exit
