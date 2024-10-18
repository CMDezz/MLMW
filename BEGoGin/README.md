# Backend - Music Library Management Website

## Overview

This folder contains the backend code for the Music Library Management Website (MLMW), built with Golang using the Gin web framework and PostgreSQL.

## Tech Stack

- **Golang** (version: 1.22.3)
- **Gin** (web framework)
- **PostgreSQL** (using `postgres-alpine:16`)

## Installation and Setup

### Prerequisites

Make sure you have Docker installed on your machine.

### Running the Project

To run the backend and database using Docker, execute the following command:

```bash
docker-compose up -d --build
```

### Troubleshooting

If `docker-compose up` doesn't work due to issues with `start.sh` or `wait-for.sh`, follow these steps:

1. Copy the entire content of the script files into a text editor (e.g., Notepad).
2. Copy the content back to the respective script files.

### Using Make

You can also use `make` to run the app with the following commands step by step:

```bash
make postgres
make createdb
make migrateInitUp
make start
```
