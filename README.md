# Number Classifier API

A RESTful API that analyzes numbers and returns their mathematical properties along with interesting facts.

## Features

- Determines if a number is prime
- Checks if a number is perfect
- Identifies Armstrong numbers
- Calculates digit sum
- Returns even/odd property
- Provides interesting mathematical facts via Numbers API
- Redis caching for improved performance

## API Documentation

### Classify Number

Returns mathematical properties and facts about a given number.

**Endpoint:** `GET /api/classify-number`

**Query Parameters:**
- `number` (required): Integer to analyze

**Success Response (200 OK):**
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

**Error Response (400 Bad Request):**
```json
{
    "number": "invalid",
    "error": true
}
```

**Property Combinations:**
- `["armstrong", "odd"]` - Armstrong number that is odd
- `["armstrong", "even"]` - Armstrong number that is even
- `["odd"]` - Odd number
- `["even"]` - Even number

### Health Check

Endpoint to verify API health status.

**Endpoint:** `GET /health`

**Success Response (200 OK):**
```json
{
    "status": "healthy"
}
```

## Setup and Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/number-classifier.git
cd number-classifier
```

2. Run with Docker Compose:
```bash
docker-compose up -d
```

3. Or build and run locally:
```bash
go mod download
go run cmd/server/main.go
```

## Environment Variables

- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin framework mode (default: release)
- `NUMBERS_API_BASE_URL`: Numbers API base URL (default: http://numbersapi.com)
- `REDIS_URL`: Redis connection URL (optional)

## Examples

1. Get properties of number 42:
```bash
curl "http://localhost:8080/api/classify-number?number=42"
```

2. Check if 371 is an Armstrong number:
```bash
curl "http://localhost:8080/api/classify-number?number=371"
```

## Development

Run unit tests:
```bash
go test ./...
```

Build Docker image:
```bash
docker build -t number-classifier .
```

## Testing the API

The project includes a test script that verifies various API endpoints and responses.

1. Run the tests (defaults to localhost):
```bash
./scripts/test_api.sh
```

2. Test against a specific API URL:
```bash
./scripts/test_api.sh "http://your-api-url:8080"
```

The script tests:
- Health endpoint
- Armstrong number (371)
- Perfect number (6)
- Prime number (17)
- Regular number (42)
- Invalid input handling

Requirements:
- curl
- jq (for JSON formatting)

Install requirements on Ubuntu:
```bash
sudo apt-get update && sudo apt-get install -y curl jq
```

## License

MIT License
