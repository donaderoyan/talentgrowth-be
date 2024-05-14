# talentgrowth-be
Backend side

## Getting Started with talentgrowth-be

This guide will help you get the backend side of the `talentgrowth-be` up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:
- Docker
- Docker Compose
- Go (Golang)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourgithub/talentgrowth-be.git
   cd talentgrowth-be
   ```

2. Set up the environment variables:
   - Copy the `.envxample` file to `.env`.
   - The dev `.env` file should look like this (modify the variables as needed):
     ```
        GO_PORT=7890
        GO_ENV=test

        MONGO_URI=mongodb://root:password123@mongodb:27017/?authSource=admin
        MONGO_DB=talentgrowth
     ```

3. Build and run the Docker containers:
   ```bash
   docker compose up --build
   ```

### Running the tests

To run tests for the application, use the following command:
```bash
go test -v ./...
```

### Setting Environment Variables for Testing

Before running the tests, ensure that the MongoDB server is up as the tests may interact with MongoDB collections.

#### Setting Environment Variables Temporarily

To set an environment variable temporarily for the duration of the test command, use the following syntax:

**On Unix-like systems (Linux, macOS):**
You can set the environment variable inline just before the command without affecting the global environment:
```
MONGO_URI="mongodb://your_uri_here" go test -v ./...
```
**On Windows:**
You can use the set command before running go test:
```
set MONGO_URI=mongodb://your_uri_here
go test -v ./...
```

Alternatively, for a one-liner in Windows PowerShell:
```
$env:MONGO_URI="mongodb://your_uri_here"; go test -v ./...
```

These commands set the environment variable temporarily for the duration of the test command. If you need the variable to persist, you would set it in your shell's configuration file or Windows environment settings.


### Running the application

Once the Docker containers are up and running, the backend server should be accessible via:
