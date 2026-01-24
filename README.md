# Gator-CLI

A command-line RSS feed aggregator written in Go. Gator allows you to manage multiple RSS feeds, follow feeds, and browse posts all from the terminal.

## Features

- **User Management**: Register and manage multiple user accounts
- **Feed Management**: Add and manage RSS feeds
- **Feed Following**: Follow specific feeds to aggregate their posts
- **Feed Aggregation**: Automatically fetch and aggregate posts from followed feeds at specified intervals
- **Browse Posts**: View posts from followed feeds with pagination

## Prerequisites

- **Go** 1.25.4 or higher
- **PostgreSQL** database

## Installation

### Option 1: Using `go install` (Recommended)

If you have Go installed, you can install Gator directly:

```bash
go install github.com/deimerin/gator-cli@latest
```

This will build and install the `gator` binary to your `$GOPATH/bin` directory. Make sure it's in your `$PATH` so you can run `gator` from anywhere.

### Option 2: Manual Build

1. Clone the repository:
```bash
git clone https://github.com/deimerin/gator-cli.git
cd gator-cli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the binary:
```bash
go build -o gator
```

Then run with `./gator` or move the binary to your `$PATH`.

### Setup

1. Set up your database:
   - Create a PostgreSQL database
   - Run migrations from the `sql/schema/` directory to set up tables

2. Configure the tool:
   - Create a file named `.gatorconfig.json` in your home directory (`~/.gatorconfig.json`)
   - Add your PostgreSQL database URL to the config
   - Example configuration:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator_db"
}
```

## Usage

Once installed, you can run Gator commands with:

```bash
gator [command] [arguments]
```

### Available Commands

#### User Management

**Register a new user:**
```bash
gator register <username>
```

**Login to an existing user:**
```bash
gator login <username>
```

**List all users:**
```bash
gator users
```

**Reset all data:**
```bash
gator reset
```

#### Feed Management

**Add a new feed:**
```bash
gator addfeed <feed_name> <feed_url>
```

**List all feeds:**
```bash
gator feeds
```

**Follow a feed:**
```bash
gator follow <feed_url>
```

**View your followed feeds:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow <feed_url>
```

#### Feed Aggregation

**Start aggregating feeds:**
```bash
gator agg <time_interval>
```

Where `<time_interval>` is a duration string like:
- `30s` - 30 seconds
- `1m` - 1 minute
- `5m` - 5 minutes
- `1h` - 1 hour

Example:
```bash
gator agg 1m
```

#### Browse Posts

**Browse posts from followed feeds:**
```bash
gator browse <number_of_posts>
```

Example:
```bash
gator browse 10
```

This will display the 10 most recent posts from your followed feeds.

## Project Structure

```
.
├── commands.go           # Command registration and routing
├── handler_*.go          # Command handlers
├── main.go              # Entry point
├── rss_feed.go          # RSS feed parsing
├── internal/
│   ├── config/          # Configuration management
│   └── database/        # Database queries and models
└── sql/
    ├── queries/         # SQL query definitions
    └── schema/          # Database migrations
```

## Example Workflow

```bash
# Register a new user
gator register alice

# Login to user
gator login alice

# Add an RSS feed
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"

# Follow the feed
gator follow "https://news.ycombinator.com/rss"

# Start aggregating (fetches feeds every minute)
gator agg 1m

# In another terminal, browse posts
gator browse 20
```

## License

This project is open source and available under the MIT License.
