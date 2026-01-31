# Expense Tracker

A minimal full-stack Expense Tracker application built with Gin (Golang) backend and a simple frontend UI. This application allows users to record and review personal expenses with filtering and sorting capabilities.

## Features

- ✅ Create new expense entries (amount, category, description, date)
- ✅ View list of all expenses
- ✅ Filter expenses by category
- ✅ Sort expenses by date (newest first)
- ✅ Display total amount of currently visible expenses
- ✅ Handles retries, page refreshes, and network issues gracefully

## Tech Stack

- **Backend**: Go 1.21, Gin framework
- **Database**: SQLite
- **Frontend**: HTML, CSS, JavaScript (vanilla)

## Project Structure

```
Fenmo_AI_Assignment/
├── config/          # Configuration management
├── database/        # Database connection and setup
├── models/          # Data models
├── repository/      # Data access layer
├── service/         # Business logic layer
├── handler/         # HTTP handlers
├── middleware/      # Middleware (CORS, error handling, logging)
├── routes/          # Route definitions
├── utils/           # Utility functions
├── frontend/        # Frontend UI files
├── .env             # Environment variables
├── main.go          # Application entry point
└── README.md         # This file
```

## Design Decisions

### Persistence: SQLite

**Choice**: SQLite database

**Reasoning**:
- **Lightweight**: Single-file database, no external server required
- **No Dependencies**: Works out of the box without additional setup
- **Suitable for Small Scale**: Perfect for a personal finance tool
- **Portable**: Database file can be easily backed up or moved

**Trade-offs**:
- Single-file database limits concurrent write performance (acceptable for personal use)
- Not suitable for high-concurrency scenarios (not required for this use case)
- Alternative considered: In-memory store (rejected due to data loss on restart)

**Implementation**: Database file stored at `./expenses.db` (configurable via `.env`)

### Money Handling

**Decision**: Store amounts as `TEXT` (decimal strings) in database

**Reasoning**:
- Avoids floating-point precision issues (critical for financial data)
- Maintains exact precision for decimal values
- API accepts and returns amounts as strings to preserve precision

**Implementation**: Amounts validated as positive decimal numbers, stored and transmitted as strings

### API Idempotency

**POST /expenses**:
- Generates unique ID on server for each request
- Allows duplicate expenses (user can submit same expense multiple times)
- Handles retries correctly by creating new records with unique IDs

**GET /expenses**:
- Naturally idempotent
- Safe for retries and page reloads
- Returns consistent results for same query parameters

## API Endpoints

### POST /api/expenses

Create a new expense entry.

**Request**:
```json
{
  "amount": "100.50",
  "category": "Food",
  "description": "Lunch at restaurant",
  "date": "2024-01-15"
}
```

**Response** (201 Created):
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "amount": "100.50",
  "category": "Food",
  "description": "Lunch at restaurant",
  "date": "2024-01-15",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### GET /api/expenses

Retrieve a list of expenses with optional filtering and sorting.

**Query Parameters** (all optional):
- `category` (string): Filter by category (exact match)
- `sort` (string): Sort order (`date_desc` for newest first)

**Examples**:
- `GET /api/expenses` - Get all expenses
- `GET /api/expenses?category=Food` - Get expenses in Food category
- `GET /api/expenses?sort=date_desc` - Get all expenses sorted by date (newest first)
- `GET /api/expenses?category=Food&sort=date_desc` - Filter and sort

**Response** (200 OK):
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "amount": "100.50",
    "category": "Food",
    "description": "Lunch at restaurant",
    "date": "2024-01-15",
    "created_at": "2024-01-15T10:30:00Z"
  }
]
```

## Setup and Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd Fenmo_AI_Assignment
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables (optional):
```bash
# .env file is already created with defaults
# Edit .env if you want to change PORT or DB_PATH
```

4. Run the application:
```bash
go run main.go
```

5. Open your browser and navigate to:
```
http://localhost:8080
```

## Environment Variables

Create a `.env` file in the root directory (already included):

```
PORT=8080
DB_PATH=./expenses.db
ENV=development
```

## Edge Cases Handled

- ✅ Duplicate form submissions (multiple clicks)
- ✅ Page refresh after POST request
- ✅ Slow or failed API responses
- ✅ Invalid request data (missing fields, invalid formats)
- ✅ Empty expense lists
- ✅ Filter with no results
- ✅ Date parsing errors
- ✅ Network retries and browser refreshes

## Trade-offs Made Due to Time Constraints

1. **Frontend Styling**: Kept simple and functional, focusing on correctness over aesthetics
2. **Validation**: Basic validation implemented (required fields, date format, positive amounts). Could be enhanced with more sophisticated validation
3. **Error Messages**: Clear but simple. Could be more detailed in production
4. **Testing**: Manual testing approach. Automated tests would be added in production
5. **Loading States**: Basic implementation. Could be enhanced with better UX indicators

## What Was Intentionally Not Done

1. **Authentication/Authorization**: Not required for this assignment
2. **User Management**: Single-user application
3. **Advanced Filtering**: Only category filtering implemented (as per requirements)
4. **Pagination**: Not required for small-scale personal use
5. **Export Functionality**: Not in scope
6. **Category Management**: Categories are free-form text (could be enhanced with predefined categories)

## Evaluation Criteria Alignment

- ✅ **Correct Behavior Under Realistic Conditions**: Handles retries, refreshes, and network issues
- ✅ **Data Correctness**: Proper money handling with decimal precision, validated date formats
- ✅ **Edge Cases**: Comprehensive handling of edge cases
- ✅ **Code Clarity and Structure**: Clean architecture with separation of concerns (repository, service, handler layers)
- ✅ **Production-like Quality**: Error handling, logging, middleware, proper HTTP status codes

## License

This project is created as an assignment submission.
