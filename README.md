# Gator - Blog Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL. Gator allows you to manage multiple RSS feeds, follow feeds from other users, and browse aggregated blog posts.

## Prerequisites

Before running Gator, ensure you have the following installed:

- **Go** (version 1.25.3 or later) - [Download Go](https://golang.org/dl/)
- **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the Gator CLI using `go install`:

```bash
go install github.com/iTsLhaj/gator@latest
```

This will compile and install the `gator` command to your `$GOPATH/bin` directory.

## Configuration

Before using Gator, you need to create a configuration file:

1. Create a file named `.gatorconfig.json` in your home directory (`~/.gatorconfig.json`)
2. Add your PostgreSQL connection string and initial username:

```json
{
  "db_url": "postgres://user:password@localhost/gator_db",
  "current_user_name": "your_username"
}
```

Replace:
- `user` with your PostgreSQL username
- `password` with your PostgreSQL password
- `gator_db` with your desired database name

## Running the Program

Once configured, you can run Gator commands:

```bash
gator <command> [arguments]
```

## Available Commands

### User Management

- **register** - Create a new user account
  ```bash
  gator register john_doe
  ```

- **login** - Switch to an existing user
  ```bash
  gator login john_doe
  ```

- **users** - List all users (current user marked with `(current)`)
  ```bash
  gator users
  ```

### Feed Management

- **addfeed** - Add a new RSS feed and follow it
  ```bash
  gator addfeed "Feed Name" "https://example.com/feed.xml"
  ```

- **feeds** - Display all available feeds in a formatted table
  ```bash
  gator feeds
  ```

- **follow** - Follow an existing feed
  ```bash
  gator follow "https://example.com/feed.xml"
  ```

- **following** - Show feeds you're currently following
  ```bash
  gator following
  ```

- **unfollow** - Stop following a feed
  ```bash
  gator unfollow "https://example.com/feed.xml"
  ```

### Browsing

- **browse** - Display recent posts from followed feeds (defaults to 2 posts)
  ```bash
  gator browse
  gator browse 10  # Show 10 posts
  ```

### Feed Aggregation

- **agg** - Continuously fetch and aggregate feeds at specified intervals
  ```bash
  gator agg 30s    # Fetch feeds every 30 seconds
  gator agg 5m     # Fetch feeds every 5 minutes
  ```

### Database Management

- **reset** - Clear all users, feeds, and follows (use with caution!)
  ```bash
  gator reset
  ```

## Project Structure

- `internal/config/` - Configuration management
- `internal/database/` - Database query generation (sqlc)
- `sql/schema/` - Database schema migrations
- `sql/queries/` - SQL query definitions
- `main.go` - Application entry point
- `commands.go` - Command handlers
- `core.go` - Core functionality and utilities
- `rss.go` - RSS feed parsing
- `middleware.go` - Authentication middleware

## Development

To regenerate database code after modifying SQL files:

```bash
sqlc generate
```

This reads from `sqlc.yaml` and generates Go code from your SQL queries.

## License

MIT
