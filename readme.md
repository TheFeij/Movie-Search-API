# Movie Search API

## Introduction
The Movie Search project provide an API for searching movies. It offers a functionality to search for movies and retrieve detailed information about them.

## Components

- **API**: Handles HTTP requests and interacts with other services.
- **Elasticsearch**: Contains information about top 1000 imdb movies.
- **Redis**: Caches requested data.
- **RapidAPI**: Integrates with rapidapi's movie search api.

## Prerequisites
Ensure you have the following dependencies installed:

- Docker
- Kubernetes (for deployment)

## Setup
1. Clone the repository:

    ```
    git clone https://github.com/TheFeij/Movie-Search-API
    ```

2. Navigate to the project directory:

    ```
    cd Movie_Search_API
    ```

3. Start Docker containers:

    ```
    docker-compose up
    ```

4. Deploy Kubernetes resources:

    ```
    kubectl apply -f kubernetes/
    ```

## Usage
- Access the API at `http://localhost:8080` when using docker compose.
- When deployed with Kubernetes, access the API using the Kubernetes cluster IP and port 30000.
- Use the `/search` endpoint to search for movies.

### Example Queries:
1. Search your queries:
    - Endpoint: `/search?query=avatar`
    - This query will return movies and series related to "avatar"

## Development
- To run tests, use `make test`.
- Check individual component directories for specific test instructions.
- To generate and view the Go documentation, run `make godoc` and access it at `http://localhost:6060`.

## Contributing
Contributions are welcome! Feel free to open issues or pull requests.
