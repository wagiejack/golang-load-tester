# Go HTTP Load Tester

A lightweight, concurrent HTTP load testing tool built in Go. This project was completed as part of the [Coding Challenges - Build Your Own Load Tester](https://codingchallenges.fyi/challenges/challenge-load-tester/).

## Prerequisites
1. Install Go (version 1.16 or higher)
   ```bash
   # For Ubuntu/Debian
   sudo apt-get install golang-go

   # For MacOS
   brew install go

   # For Windows
   # Download from https://go.dev/dl/
   ```

2. Clone the repository
   ```bash
   git clone [your-repository-url]
   cd [repository-name]
   ```

## Examples of Testing

### Basic Testing
Test a single endpoint:
```bash
go run . -u https://api.example.com
```

### Load Testing
Send 100 concurrent POST requests with a JSON payload:
```bash
go run . -u https://api.example.com/users \
         -m POST \
         -bd '{"name": "test"}' \
         -c 100
```

### Batch Testing
Test multiple endpoints defined in a file with 10 concurrent requests:
```bash
go run . -f urls.txt -c 10 -n 5
```

## Command Line Arguments
```
-u  URL to test (required if not using -f)
-m  HTTP method (GET, POST, PUT, PATCH, DELETE)
-bd Request body for POST/PUT/PATCH methods
-n  Number of requests per URL (default: 1)
-c  Number of concurrent requests (default: 1)
-f  File containing URLs to test
```

## URL File Format
```
GET https://api.example.com
POST https://api.example.com/users {"name": "test"}
PUT https://api.example.com/users/1 {"name": "updated"}
DELETE https://api.example.com/users/1
```

## Features

### Core Functionality
- Concurrent request execution using Go routines
- Support for all standard HTTP methods (GET, POST, PUT, PATCH, DELETE)
- URL validation and sanitization
- Batch processing from file input
- Configurable request concurrency and count

### Performance Metrics
- Request timing analysis
  - Total request time
  - Time to First Byte (TTFB)
  - Time to Last Byte (TTLB)
- Response statistics
  - Success/failure rate
  - Status code distribution
  - Request per second calculation
- Performance summary
  - Min/Max/Mean timing calculations
  - Concurrent request handling stats
  - Error rate analysis

### Error Handling
- Network error detection
- Invalid URL detection
- Malformed request handling
- Connection timeout management
