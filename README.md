# Music Microservices Application (AP2-Final)

A scalable microservices-based music platform developed for the Advanced Programming 2 course at Astana IT University (AITU). The project demonstrates modern backend engineering practices using Go, gRPC, Protocol Buffers, and PostgreSQL.

Instead of using a monolithic architecture, the system is separated into independent microservices communicating through high-performance gRPC (HTTP/2) calls.

##  Project Goals
This project was designed to demonstrate:
* Microservices architecture principles
* Service-to-service communication with gRPC
* API Gateway pattern implementation
* Clean backend architecture using the Repository Pattern
* REST ↔ gRPC request translation
* PostgreSQL integration
* Scalable distributed backend design

---

##  System Architecture

The application consists of four independent backend services and one shared Proto module.

```text
                +----------------------+
                |      Frontend        |
                |   Postman / Client   |
                +----------+-----------+
                           |
                           | HTTP/JSON
                           v
                +----------------------+
                |     API Gateway      |
                |       :8080          |
                |   Gin + REST API     |
                +----------+-----------+
         +-----------------+-----------------+
         |                 |                 |
         v                 v                 v
 +--------------+ +--------------+ +--------------+
 | User Service | | CatalogServ. | | PlaylistServ.|
 |    :50051    | |    :50052    | |    :50053    |
 +------+-------+ +------+-------+ +------+-------+
        |                |                |
        v                v                v
   PostgreSQL       PostgreSQL       PostgreSQL
 Services Overview1. API Gateway (:8080)The API Gateway acts as the single public entry point for all client requests.Responsibilities: Accept REST HTTP requests, parse JSON payloads, route requests to proper gRPC services, convert REST ↔ Protobuf messages, and return JSON responses to clients.Technologies: Gin-Gonic, gRPC Client, REST API2. User Service (:50051)Handles all user-related operations.Features: User registration, authentication, password handling, and database interaction.Database: PostgreSQLArchitecture: Repository Pattern, gRPC Server3. Catalog Service (:50052)Responsible for music catalog management.Features: Add tracks, retrieve track metadata, search tracks, and manage catalog storage.Track Metadata: Title, Artist, Duration4. Playlist Service (:50053)Manages playlists and track associations.Features: Create playlists, attach tracks to playlists, maintain user-playlist relationships, and manage playlist retrieval.5. Proto Module (/proto)Contains all shared .proto contract files.Responsibilities: Define gRPC services, define Protobuf message structures, and generate Go code for both gRPC clients and servers.💻 Tech StackCategoryTechnologyLanguageGo (Golang)API FrameworkGin-GonicCommunicationgRPCSerializationProtocol Buffers v3DatabasePostgreSQLArchitectureMicroservicesPatternsAPI Gateway, Repository PatternProtocolHTTP/2📂 Project StructurePlaintextap2-final/
│
├── api-gateway/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── user-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── catalog-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── playlist-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── proto/
│   ├── user.proto
│   ├── catalog.proto
│   ├── playlist.proto
│   └── generated/
│
├── docker-compose.yml
├── go.work
├── go.mod
└── README.md
 Running the Project LocallyBefore running the project, ensure that Go is installed, PostgreSQL is running, and ports 8080, 50051, 50052, and 50053 are available.Open 4 separate terminals in the root directory and execute the following:Terminal 1: User ServiceBashcd user-service
go run cmd/main.go
Terminal 2: Catalog ServiceBashcd catalog-service
go run cmd/main.go
Terminal 3: Playlist ServiceBashcd playlist-service
go run cmd/main.go
Terminal 4: API GatewayBashcd api-gateway
go run cmd/main.go
🛠️ API SpecificationAll client requests must go through the API Gateway.Base URL: http://localhost:80801. Register a New UserEndpoint: POST /registerRequest Body:JSON{
  "username": "bakdaulet",
  "email": "test@aitu.kz",
  "password": "superpassword123"
}
Example Response:JSON{
  "message": "user registered successfully"
}
2. Add Track to CatalogEndpoint: POST /tracksRequest Body:JSON{
  "title": "New Song",
  "artist": "Famous Artist",
  "duration": 180
}
Example Response:JSON{
  "id": "track_1",
  "message": "track created successfully"
}
3. Get Track by IDEndpoint: GET /tracks/:idExample Request: GET /tracks/track_1Example Response:JSON{
  "id": "track_1",
  "title": "New Song",
  "artist": "Famous Artist",
  "duration": 180
}
4. Create PlaylistEndpoint: POST /playlistsRequest Body:JSON{
  "user_id": "1",
  "title": "My Favorite Tracks",
  "track_ids": [
    "track_1",
    "track_2",
    "track_3"
  ]
}
Example Response:JSON{
  "playlist_id": "playlist_1",
  "message": "playlist created successfully"
}
 Testing the ApplicationYou can test the system using Postman, cURL, Insomnia, or frontend client applications.Example cURL RequestBashcurl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
  "username":"bakdaulet",
  "email":"test@aitu.kz",
  "password":"superpassword123"
}'
 Why gRPC?This project leverages gRPC for internal service-to-service communication because it provides:High performance natively running via HTTP/2.Compact binary serialization (Protocol Buffers) reducing payload sizes.Strongly typed contracts preventing schema mismatches between services.Faster processing speeds compared to traditional REST JSON APIs.Automatic code generation for microservice communication layers.🛠️ Future ImprovementsPlanned enhancements for future iterations:Security: Implement JWT Authentication & API Rate Limiting.Deployment: Add Docker containerization and Kubernetes deployment.Performance: Introduce Redis caching and Message Brokers (Kafka/RabbitMQ).Observability: Set up unified monitoring with Prometheus + Grafana.
