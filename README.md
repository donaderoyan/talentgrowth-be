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

### Running the application

Once the Docker containers are up and running, the backend server should be accessible via:
