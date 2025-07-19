# API Endpoints

This document outlines all REST API endpoints for the swutrack platform.

## Base URL
- **Development**: `http://localhost:8080`
- **Production**: `https://api.swutrack.com`

## Common Response Formats

### Success Response
```json
{
  "data": {},
  "message": "Success message"
}
```

### Error Response
```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message",
    "details": {}
  }
}
```

### Pagination Response
```json
{
  "data": [],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "total_pages": 5
  }
}
```

## Authentication

### OAuth Login
**POST** `/auth/login`

Initiates OAuth flow with supported providers.

**Request Body:**
```json
{
  "provider": "google" | "discord"
}
```

**Response:**
```json
{
  "data": {
    "redirect_url": "https://oauth.provider.com/authorize?..."
  }
}
```

### OAuth Callback
**GET** `/auth/callback`

Handles OAuth provider callback and returns user session.

**Query Parameters:**
- `code` (string): Authorization code from OAuth provider
- `state` (string): CSRF protection state parameter

**Response:**
```json
{
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "username": "username",
      "oauth_provider": "google",
      "created_at": "2024-01-01T00:00:00Z"
    },
    "session_token": "session_token_here"
  }
}
```

### Logout
**POST** `/auth/logout`

Invalidates the current session.

**Headers:** `Authorization: Bearer <session_token>`

## User Management

### Get Current User
**GET** `/api/v1/users/me`

Returns current authenticated user's profile.

**Headers:** `Authorization: Bearer <session_token>`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "username": "username",
    "oauth_provider": "google",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Update User Profile
**PUT** `/api/v1/users/me`

Updates current user's profile information.

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:**
```json
{
  "username": "new_username"
}
```

## Deck Management

### List User Decks
**GET** `/api/v1/decks`

Returns paginated list of user's decks with current version info.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `page` (int, default: 1): Page number
- `limit` (int, default: 20): Items per page
- `active_only` (bool, default: false): Show only active decks
- `search` (string): Search deck names

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "name": "Deck Name",
      "leader": "card_id_luke_skywalker",
      "base": "card_id_echo_base",
      "description": "Deck description",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "current_version": {
        "id": "uuid",
        "version": "2.0",
        "created_at": "2024-01-02T00:00:00Z"
      },
      "total_versions": 3,
      "game_count": 15,
      "win_rate": 0.67
    }
  ],
  "pagination": {...}
}
```

### Get Deck Details
**GET** `/api/v1/decks/{deck_id}`

Returns detailed information about a specific deck with current version decklist.

**Headers:** `Authorization: Bearer <session_token>`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "name": "Deck Name",
    "leader": "card_id_luke_skywalker",
    "base": "card_id_echo_base",
    "description": "Deck description",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "current_version": {
      "id": "uuid",
      "version": "2.0",
      "decklist": {
        "cards": [
          {
            "id": "card_id_xwing",
            "quantity": 3,
            "type": "unit",
            "cost": 4
          }
        ],
        "sideboard": [
          {
            "id": "card_id_sabotage",
            "quantity": 2,
            "type": "event",
            "cost": 1
          }
        ]
      },
      "created_at": "2024-01-02T00:00:00Z"
    }
  }
}
```

### Create Deck
**POST** `/api/v1/decks`

Creates a new deck with initial version (automatically set to "1.0").

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:**
```json
{
  "name": "Deck Name",
  "leader": "card_id_luke_skywalker",
  "base": "card_id_echo_base",
  "description": "Optional description",
  "decklist": {
    "cards": [
      {
        "id": "card_id_xwing",
        "quantity": 3,
        "type": "unit",
        "cost": 4
      }
    ],
    "sideboard": [
      {
        "id": "card_id_sabotage",
        "quantity": 2,
        "type": "event",
        "cost": 1
      }
    ]
  }
}
```

### Update Deck
**PUT** `/api/v1/decks/{deck_id}`

Updates deck metadata (name and description only). Leader and base cannot be changed after deck creation.

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:**
```json
{
  "name": "Updated Deck Name",
  "description": "Updated description"
}
```

### Delete Deck
**DELETE** `/api/v1/decks/{deck_id}`

Soft deletes a deck (sets is_active to false).

**Headers:** `Authorization: Bearer <session_token>`

## Deck Version Management

### List Deck Versions
**GET** `/api/v1/decks/{deck_id}/versions`

Returns all versions of a specific deck ordered by creation date (newest first).

**Headers:** `Authorization: Bearer <session_token>`

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "version": "2.0",
      "is_current": true,
      "created_at": "2024-01-02T00:00:00Z",
      "game_count": 8,
      "win_rate": 0.75
    },
    {
      "id": "uuid",
      "version": "1.0",
      "is_current": false,
      "created_at": "2024-01-01T00:00:00Z",
      "game_count": 7,
      "win_rate": 0.57
    }
  ]
}
```

### Get Deck Version Details
**GET** `/api/v1/decks/{deck_id}/versions/{version_id}`

Returns detailed information about a specific deck version including decklist.

**Headers:** `Authorization: Bearer <session_token>`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "deck_id": "uuid",
    "version": "2.0",
    "is_current": true,
    "decklist": {
      "cards": [
        {
          "id": "card_id_xwing",
          "quantity": 3,
          "type": "unit",
          "cost": 4
        }
      ],
      "sideboard": [
        {
          "id": "card_id_sabotage",
          "quantity": 2,
          "type": "event",
          "cost": 1
        }
      ]
    },
    "created_at": "2024-01-02T00:00:00Z"
  }
}
```

### Create Deck Version
**POST** `/api/v1/decks/{deck_id}/versions`

Creates a new version of an existing deck and sets it as current. Version number is automatically determined by the system.

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:**
```json
{
  "decklist": {
    "cards": [
      {
        "id": "card_id_xwing",
        "quantity": 3,
        "type": "unit",
        "cost": 4
      }
    ],
    "sideboard": [
      {
        "id": "card_id_sabotage",
        "quantity": 2,
        "type": "event",
        "cost": 1
      }
    ]
  }
}
```

### Set Current Version
**PUT** `/api/v1/decks/{deck_id}/versions/{version_id}/current`

Sets a specific version as the current version for the deck.

**Headers:** `Authorization: Bearer <session_token>`

### Delete Deck Version
**DELETE** `/api/v1/decks/{deck_id}/versions/{version_id}`

Deletes a specific deck version. Cannot delete if games reference this version.

**Headers:** `Authorization: Bearer <session_token>`

## Game Tracking

### List Games
**GET** `/api/v1/games`

Returns paginated list of user's games.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `page` (int, default: 1): Page number
- `limit` (int, default: 20): Items per page
- `deck_id` (uuid): Filter by deck
- `deck_version_id` (uuid): Filter by specific deck version
- `result` (string): Filter by result (win/loss/draw)
- `opponent_leader` (string): Filter by opponent leader card ID
- `date_from` (date): Filter games from date (YYYY-MM-DD)
- `date_to` (date): Filter games to date (YYYY-MM-DD)

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "deck_version": {
        "id": "uuid",
        "deck_id": "uuid",
        "deck_name": "Deck Name",
        "version": "2.0"
      },
      "opponent_leader": "card_id_darth_vader",
      "opponent_base": "card_id_death_star",
      "result": "win",
      "notes": "Close game, won by 1 point",
      "game_date": "2024-01-01T15:30:00Z",
      "created_at": "2024-01-01T15:35:00Z"
    }
  ],
  "pagination": {...}
}
```

### Get Game Details
**GET** `/api/v1/games/{game_id}`

Returns detailed information about a specific game.

**Headers:** `Authorization: Bearer <session_token>`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "deck_version": {
      "id": "uuid",
      "deck_id": "uuid",
      "deck_name": "Deck Name",
      "deck_leader": "card_id_luke_skywalker",
      "deck_base": "card_id_echo_base",
      "version": "2.0"
    },
    "opponent_leader": "card_id_darth_vader",
    "opponent_base": "card_id_death_star",
    "result": "win",
    "notes": "Close game, won by 1 point",
    "game_date": "2024-01-01T15:30:00Z",
    "created_at": "2024-01-01T15:35:00Z"
  }
}
```

### Record Game
**POST** `/api/v1/games`

Records a new game result.

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:**
```json
{
  "deck_version_id": "uuid",
  "opponent_leader": "card_id_darth_vader",
  "opponent_base": "card_id_death_star",
  "result": "win" | "loss" | "draw",
  "notes": "Optional game notes",
  "game_date": "2024-01-01T15:30:00Z"
}
```

### Update Game
**PUT** `/api/v1/games/{game_id}`

Updates an existing game record.

**Headers:** `Authorization: Bearer <session_token>`

**Request Body:** Same as Record Game

### Delete Game
**DELETE** `/api/v1/games/{game_id}`

Deletes a game record.

**Headers:** `Authorization: Bearer <session_token>`

## Statistics & Analytics

### Dashboard Overview
**GET** `/api/v1/stats/overview`

Returns high-level statistics for dashboard.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `period` (string, default: "all"): Time period (30d, 90d, 1y, all)

**Response:**
```json
{
  "data": {
    "total_games": 150,
    "total_wins": 95,
    "total_losses": 52,
    "total_draws": 3,
    "overall_win_rate": 0.633,
    "active_decks": 5,
    "total_deck_versions": 12,
    "games_this_week": 8,
    "favorite_deck": {
      "id": "uuid",
      "name": "Luke Aggro",
      "current_version": "2.0",
      "game_count": 35
    },
    "recent_performance": [
      {
        "date": "2024-01-01",
        "wins": 3,
        "losses": 1
      }
    ]
  }
}
```

### Deck Performance
**GET** `/api/v1/stats/decks`

Returns performance statistics by deck and version.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `period` (string, default: "all"): Time period
- `deck_id` (uuid): Specific deck stats
- `by_version` (bool, default: false): Group stats by deck version

**Response:**
```json
{
  "data": [
    {
      "deck_id": "uuid",
      "deck_name": "Luke Aggro",
      "total_games": 25,
      "wins": 18,
      "losses": 6,
      "draws": 1,
      "win_rate": 0.72,
      "versions": [
        {
          "version_id": "uuid",
          "version": "2.0",
          "total_games": 15,
          "wins": 12,
          "losses": 3,
          "win_rate": 0.80
        },
        {
          "version_id": "uuid",
          "version": "1.0",
          "total_games": 10,
          "wins": 6,
          "losses": 3,
          "draws": 1,
          "win_rate": 0.60
        }
      ]
    }
  ]
}
```

### Matchup Analysis
**GET** `/api/v1/stats/matchups`

Returns win/loss statistics against different leader/base combinations.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `period` (string, default: "all"): Time period
- `deck_id` (uuid): Filter by specific deck
- `deck_version_id` (uuid): Filter by specific deck version

**Response:**
```json
{
  "data": [
    {
      "opponent_leader": "card_id_darth_vader",
      "opponent_base": "card_id_death_star",
      "total_games": 15,
      "wins": 9,
      "losses": 6,
      "win_rate": 0.60,
      "most_used_deck": "Luke Aggro",
      "most_used_version": "2.0"
    }
  ]
}
```

### Performance Trends
**GET** `/api/v1/stats/trends`

Returns performance data over time for charts.

**Headers:** `Authorization: Bearer <session_token>`

**Query Parameters:**
- `period` (string, default: "90d"): Time period
- `granularity` (string, default: "week"): Data granularity (day, week, month)
- `deck_id` (uuid): Filter by specific deck
- `deck_version_id` (uuid): Filter by specific deck version

**Response:**
```json
{
  "data": {
    "win_rate_trend": [
      {
        "period": "2024-W01",
        "win_rate": 0.67,
        "total_games": 6
      }
    ],
    "games_per_period": [
      {
        "period": "2024-W01",
        "game_count": 6
      }
    ],
    "deck_usage": [
      {
        "deck_name": "Luke Aggro",
        "version": "2.0",
        "percentage": 0.45
      }
    ]
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| `INVALID_REQUEST` | Request validation failed |
| `UNAUTHORIZED` | Authentication required |
| `FORBIDDEN` | Access denied |
| `NOT_FOUND` | Resource not found |
| `CONFLICT` | Resource conflict (duplicate name/version, etc.) |
| `RATE_LIMITED` | Too many requests |
| `SERVER_ERROR` | Internal server error |
| `VALIDATION_ERROR` | Input validation failed |
| `DECK_NOT_FOUND` | Specified deck does not exist |
| `DECK_VERSION_NOT_FOUND` | Specified deck version does not exist |
| `GAME_NOT_FOUND` | Specified game does not exist |
| `INVALID_DECK_DATA` | Deck data is invalid or incomplete |
| `VERSION_HAS_GAMES` | Cannot delete version with associated games |

## Rate Limiting

- **General API**: 1000 requests per hour per user
- **Authentication**: 10 requests per minute per IP
- **Statistics**: 100 requests per hour per user

Rate limit headers are included in all responses:
- `X-RateLimit-Limit`: Request limit per window
- `X-RateLimit-Remaining`: Requests remaining in current window
- `X-RateLimit-Reset`: Unix timestamp when the window resets