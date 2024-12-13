# NETOPIA Payment SDK - Sample Application

This is a sample Go application demonstrating the integration with the NETOPIA Payment SDK using the [Gin Web Framework](https://gin-gonic.com/).

## Features

- **Start Payment**: Initiates a payment request with necessary configurations.
- **Check Payment Status**: Queries the payment status based on transaction and order IDs.
- **IPN Verification**: Processes and validates IPN (Instant Payment Notification) messages from NETOPIA.

## Requirements

- Go version 1.18 or higher
- [NETOPIA Payment SDK](https://github.com/netopiapayments/go-sdk)
- `.env` file for configuration

## Installation

1. Clone the repository:
   ```bash
    git clone https://github.com/netopiapayments/sample-app-go.git
    cd sample-app-go
   ```

2. Install dependencies:
   ```bash
    go mod tidy
   ```

3. Create a `.env` file in the root directory 
   ```bash
    cp .env.example .env
   ```

4. Run the application:
   ```bash
   go run ./cmd/main.go
   ```

## Endpoints

### 1. `/start-payment` (POST)

**Description:** Initiates a payment request.  

**Response:**
- On success, returns the payment URL to complete the transaction.

---

### 2. `/ipn` (POST)

**Description:** Processes Instant Payment Notifications (IPN) from NETOPIA.  

**Response:**
- Validates the IPN and returns a confirmation response.

---

### 3. `/status` (POST)

**Description:** Retrieves the status of a payment.  

**Payload:**
```json
{
  "ntpID": "NTPTXN123",
  "orderID": "ORDER-12345"
}
```

**Response:**
- Returns the current status of the payment.

---

## Project Structure

```
.
├── collection           # Basic POSTMAN collection for calling the endpoints of the sample app
│   └── NETOPIA GO SDK Sample.postman_collection.json        
├── cmd
│   └── main.go          # Entry point of the application
├── config
│   └── config.go        # Configuration loading
├── handlers
│   ├── payment.go       # Handlers for payment routes
│   └── status.go        # Handler for status route
│   └── ipn.go           # Handler for IPN route
├── services
│   └── netopia.go       # Business logic for payment processing
├── go.mod               # Module dependencies
├── go.sum               # Checksums for dependencies
└── .env                 # Environment variables
```

## Usage Notes

- The application is configured to run in **sandbox mode** by default (`IS_LIVE=false`). Switch to live mode by setting `IS_LIVE=true` in the `.env` file.
- The `NOTIFY_URL` and `REDIRECT_URL` must be publicly accessible for production use.
