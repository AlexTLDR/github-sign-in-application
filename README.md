# Go Web App with GitHub Authentication

This web application, written in Go, enables users to sign in using their GitHub accounts. The app leverages OAuth2 for authentication with GitHub, offering a seamless login experience and access to user data such as repositories and profile information.

## Features

- **GitHub OAuth2 Authentication**: Enables users to sign in securely using GitHub.
- **Scopes Configuration**: Requests access to both user profile and repository information (`repo`, `user` scopes).
- **Environment-Based Configuration**: Uses environment variables for managing secrets and configurations.
- **Simple Web Server**: Utilizes Go's built-in `net/http` package to serve HTTP requests.

## Prerequisites

To run the application, ensure you have the following:

- [Go](https://golang.org/dl/) 1.19 or higher installed.
- A [GitHub OAuth App](https://github.com/settings/developers) configured with:
    - **Client ID**
    - **Client Secret**
    - **Redirect URL** set to `http://localhost:8080/github/callback`

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/AlexTLDR/github-sign-in-application.git
    ```
2. Structure of the .env file:
    ```
    GITHUB_CLIENT_ID=your_github_client_id
    GITHUB_CLIENT_SECRET=your_github_client_secret
   ```
