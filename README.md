# “Talentgrowth: A Musical Course Academy Powered by Golang, Docker, and MongoDB”

Combines the power of Golang, Docker, and MongoDB to deliver an exceptional learning experience. Explore our RESTful API, designed for seamless integration and scalability. This repository contains code for learning purposes, and you’re welcome to use it freely 🎶✨

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

        JWT_SECRET="SecreteJwt"
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

Before running the tests, ensure that the MongoDB server is up as the tests may interact with MongoDB collections. To set an environment variable temporarily for the duration of the test command, use the following syntax:

   - On Unix-like systems (Linux, macOS):

      You can set the environment variable inline just before the command without affecting the global environment:
      ```
      MONGO_URI="mongodb://your_uri_here" go test -v ./...
      ```

   - On Windows:

      You can use the set command before running go test:
      ```
      set MONGO_URI=mongodb://your_uri_here
      go test -v ./...
      ```

   - Alternatively, for a one-liner in Windows PowerShell:
      ```
      $env:MONGO_URI="mongodb://your_uri_here"; go test -v ./...
      ```

These commands set the environment variable temporarily for the duration of the test command. If you need the variable to persist, you would set it in your shell's configuration file or Windows environment settings.


### Running the application

Once the Docker containers are up and running, the backend server should be accessible via http://localhost:7890. You can check the API documentation at:
``` 
http://localhost:7890/api/v1/docs/index.html 
```

To generate updated API documentation after making changes, run the following command in your terminal:
```
swag init
```

### Data Migration Documentation

#### Overview
Data migration is a critical process that involves transferring data between storage types, formats, or systems. In the context of our application, we often need to migrate user data within our MongoDB database to ensure that all documents conform to updated schema requirements.

#### Migration Process
To perform a data migration, follow these steps:

1. **Prepare Migration Script**: Before updating the migration script, confirm that the MongoDB server is up. Ensure that the migration script includes all necessary fields and default values. The script `data_migration.go` is isolated from the main application and can be found in the `migrations` directory.

2. **Set Environment Variables**: Before running the migration, make sure that the MongoDB URI and database name are correctly set in the `.env` file or as environment variables.

3. **Run the Migration**: Execute the migration script by running:
   ```bash
   go run migrations/data_migration.go
   ```
   This script connects to the MongoDB database, updates the user documents, and logs the outcome.

   To run the data migration script with an inline environment variable setting in Windows PowerShell, use the following command:
   ```
   $env:MONGO_URI="mongodb://your_uri_here"; go run migrations/data_migration.go
   ```

#### Error Handling
If the migration fails, the script will log an error with a specific failure message. Check the logs to identify any issues with the migration process.

#### Verification
Post-migration, verify the updates by querying the MongoDB database to ensure that all documents reflect the new schema changes.

#### Rollback Strategy
In case of a migration failure or if the new data model needs to be reverted, prepare a rollback script that can restore the previous state of the database documents.

### Additional Notes
- Always backup the database before running a migration to prevent data loss.
- Test the migration process in a development or staging environment before applying changes to the production database.

