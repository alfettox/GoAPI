###### Giovanni De Franceschi

# Bank Account Management Application

This is a simple Go application for managing bank accounts. It provides basic CRUD (Create, Read, Update, Delete) operations for bank accounts using a RESTful API.

## Features

- Create a new bank account
- Retrieve a list of all bank accounts
- Retrieve details of a specific bank account
- Update the information of a bank account
- Delete a bank account

## Technologies Used

- Go (Golang) - Programming language
- Gorilla Mux - Router and dispatcher for HTTP servers
- JSON - Data format for API requests and responses

## Getting Started

### Prerequisites

- Go should be installed on your system.

### Installation

1. Clone this repository to your local machine.

```bash
git clone <repository-url>
```

Navigate to the project directory.
```bash
cd <project-directory>
```
Run the application.
```bash
go run main.go
```
The application should start, and you can access it by opening a web browser or using an API client.

### Usage
##### API Endpoints
- GET /accounts: Retrieve a list of all bank accounts.
- GET /account/{number}: Retrieve details of a specific bank account.
- POST /account: Create a new bank account. Send account details in the request body as JSON.
- PUT /account/{number}: Update the information of a bank account. Send updated details in the request body as JSON.
- DELETE /account/{number}: Delete a bank account.

##### Example Requests
Create a new bank account:
json
```
POST /account

{
    "AccountNumber": "C0004",
    "Balance": "5000.0",
    "AccountDescription": "Savings Account",
    "Name": "John Doe"
}
```
Update a bank account:
```
PUT /account/C0004

{
    "AccountNumber": "C0004",
    "Balance": "5500.0",
    "AccountDescription": "Updated Savings Account",
    "Name": "John Doe"
}
```

Delete a bank account:
```
DELETE /account/C0004

```