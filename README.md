# Task Manager Go

A simple task manager application built using Go and Clean Architecture principles. This project demonstrates how to structure a Go project in a modular and scalable way using clean architecture.

## Features

- User registration
- Task creation
- Task management (create, view)
- Data is stored in memory (in future versions, you can add database support)

## Project Structure
    .
    ├── README.md
    ├── api
    ├── cmd
    │   └── main.go
    ├── docs
    ├── go.mod
    ├── internal
    │   ├── entities
    │   │   ├── task.go
    │   │   └── user.go
    │   ├── infrastructure
    │   │   └── storage.go
    │   ├── interfaces
    │   │   └── cli.go
    │   └── usecases
    │       ├── login.go
    │       ├── register.go
    │       └── task.go
    ├── pkg
    └── tests

## Setup & Installation

### Prerequisites
- Go 1.18 or later

### Steps to run the project:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/task-manager-go.git
   cd task-manager-go

2. **Install dependencies (optional if you use Go modules):**
    ```bash
    go mod tidy

3. **Run the application**
    ```
    go run cmd/main.go


## Contributing

To contribute to this repository, follow these steps:

1. **Fork the repository**
   
   Click on the "Fork" button at the top right of the repository to create your own copy.

2. **Clone your forked repository and create a new branch**
   
   Clone the repository to your local machine and create a new branch:
   ```bash
   git clone https://github.com/hassanpdn/task-manager-go.git
   cd task-manager-go
   git checkout -b feature/your-feature
3. **Make your changes**
Edit the code to add new features or fix bugs. Ensure your changes are tested and working properly.

4. **Commit your changes**
Once your changes are complete, commit them with a clear, concise message describing what was changed:

    git commit -am 'Add a new feature or fix a bug'


5. **Create a Pull Request**

Go to the original repository on GitHub, and click "New Pull Request." Describe your changes and submit the pull request.
