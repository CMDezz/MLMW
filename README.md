# MLMW - Music Library Management Website

## Overview

MLMW is a web-based application for managing a music library, allowing users to organize playlists, tracks, and other music-related data.

## Tech Stack

- **Frontend (FE)**:
  - ReactJS
  - Ant Design (Antd)
  - Tailwind CSS
- **Backend (BE)**:
  - Golang
  - Gin Web Framework
- **Database**:
  - PostgreSQL
- **Containerization**:
  - Docker (for Backend and Database)

## Features

- Playlist and track management
- Authentication
- Search music

## Setup and Installation (for detail pls view readme.md within each folder FE/BE)

### Prerequisites

- Node.js 18+ (for frontend)
- Docker (for backend and database)
- Golang 1.22.3+
- Postgres

### Steps

1. Clone the repository:

   ```bash
   https://github.com/CMDezz/MLMW.git
   ```

2. To run the entire app using Docker Compose, from the root directory:

```bash
docker compose -f docker-compose-dev.yaml up -d --build
```

3. To shutdown :

```bash
docker compose -f docker-compose-dev.yaml down
```

4. To run the app from local, please read more detail at each folder FE/BE
