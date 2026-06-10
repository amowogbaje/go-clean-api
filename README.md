# Go Clean Architecture API

A lightweight, robust REST API built in Go using the **Clean Architecture** pattern. This project features containerized environments using Docker and Docker Compose, the Gin Gonic web framework, a PostgreSQL database, and local volume mounting for persistent media uploads.

## 🚀 Features

* **Clean Architecture:** Independent layers (Domain, Usecase, Repository, Delivery) ensuring separation of concerns and high testability.
* **Zero Local Dependencies:** The entire Go toolchain and database run inside Docker containers—no local Go or Postgres installation required.
* **Persistent Local Uploads:** Uploaded media files bypass the container and are saved directly into your local machine's project directory using Docker volumes.
* **WSL2/Windows Optimized:** Configured to run on port `8888` to avoid common local port conflicts.

---

## 📂 Project Structure

```text
go_clean_api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point & dependency wiring
├── internal/
│   ├── domain/
│   │   └── media.go                # Domain models, Entities, and Interfaces
│   ├── repository/
│   │   └── postgres_media.go       # Infrastructure layer (PostgreSQL implementation)
│   ├── usecase/
│   │   └── media_usecase.go        # Core Business Logic implementation
│   └── delivery/
│       └── http/
│           └── media_handler.go    # HTTP Routing & Request Handlers (Gin Gonic)
├── uploads/                        # Locally bound folder containing uploaded files
├── .env.example                    # Template for environment configurations
├── Dockerfile                      # Multi-stage optimized Docker build configuration
└── docker-compose.yml              # Multi-container service orchestration

```

---

## 🛠️ Getting Started

### Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.

### Installation & Setup

1. **Clone or navigate to your project directory:**
```bash
cd go_clean_api

```


2. **Configure Environment Variables:**
Copy the example environment file to create your local configurations:
```bash
cp .env.example .env

```


Open the `.env` file and configure your credentials. The defaults are ready to go for local development:
```env
DB_HOST=postgres
DB_PORT=5432
POSTGRES_USER=admin
POSTGRES_PASSWORD=secretpassword
POSTGRES_DB=clean_db

```


3. **Spin up the Containers:**
Launch the database and application services using Docker Compose:
```bash
docker compose up --build

```


*Note: On successful launch, the PostgreSQL database will automatically create the required database tables, and the server will start listening for traffic at `http://127.0.0.1:8888/`.*

---

## 📋 API Documentation

### 1. Welcome Route

Returns a simple text check confirming the server is online and accessible.

* **URL:** `/api/`
* **Method:** `GET`
* **Headers:** None
* **Success Response (200 OK):**
```json
{
  "message": "welcome to using Go"
}

```



### 2. Upload Media

Uploads an image or a video file. The system dynamically parses the extension to save it as either an `image` or `video` type, maps the URL, and stores the record in Postgres.

* **URL:** `/api/upload/`
* **Method:** `POST`
* **Body Content-Type:** `multipart/form-data`
* **Payload Params:**
* `file`: *(Binary)* The actual image or video file.


* **Success Response (200 OK):**
```json
{
  "message": "File uploaded successfully",
  "data": {
    "link": "/uploads/my_awesome_photo.jpg"
  }
}

```



### 3. Fetch All Media Records

Retrieves a list of all media metadata processed and saved into the PostgreSQL database.

* **URL:** `/api/medias`
* **Method:** `GET`
* **Success Response (200 OK):**
```json
[
  {
    "id": 1,
    "urls": "/uploads/my_awesome_photo.jpg",
    "type": "image"
  },
  {
    "id": 2,
    "urls": "/uploads/holiday_video.mp4",
    "type": "video"
  }
]

```



---

## 🧪 Testing with cURL

You can quickly verify endpoints via your terminal using these pre-made cURL commands.

**Test Welcome Route:**

```bash
curl [http://127.0.0.1:8888/api/](http://127.0.0.1:8888/api/)

```

**Test Media Upload:**
*(Make sure to replace `/path/to/your/file.jpg` with a valid file path on your local computer)*

```bash
curl -X POST [http://127.0.0.1:8888/api/upload/](http://127.0.0.1:8888/api/upload/) \
  -F "file=@/path/to/your/file.jpg"

```

**Test Fetching All Medias:**

```bash
curl [http://127.0.0.1:8888/api/medias](http://127.0.0.1:8888/api/medias)

```
