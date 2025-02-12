# Zdrapywacz
A Go-based web scraper that monitors junior developer job listings and sends Discord notifications for new postings.

## Features

- 🔄 Automatically scrapes JustJoin.IT job board every 10 minutes
- 💾 Stores job offers in MariaDB database
- 🤖 Discord bot integration with:
  - Embedded messages containing job details
  - Company logos in embeds
  - Technology-specific role mentions for:
    - Vue
    - TypeScript
    - Golang
    - Svelte
    - Node.js
    - JavaScript
    - Nuxt
    - Laravel
- 🔍 Filters for junior-level positions
- 🔄 Automatic duplicate detection

## Prerequisites

- Go 1.20+
- MariaDB 11.8
- Discord Bot Token
- Discord Server with configured roles

## Setup

1. Clone the repository
```sh
git clone https://github.com/yourusername/Zdrapywacz.git
cd Zdrapywacz
```

2. Create a .env file in the root directory:
```env
DB_USER=your_mysql_user
DB_PASSWORD=your_mysql_password
DB_PORT=3306
DB_NAME=your_database_name
BOT_ID=your_discord_bot_token
```

3. Create MySQL database:
```sql
CREATE DATABASE your_database_name;
```

4. Build and run using Make:
```sh
make run
```

Or manually:
```sh
go build -o bin/output.exe
./bin/output.exe
```

## Discord Setup

1. Create a Discord bot at [Discord Developer Portal](https://discord.com/developers/applications)
2. Enable necessary intents for your bot
3. Create roles in your Discord server for technology mentions
4. Update the role IDs in `discordBot/discordbot.go`:
```go
techs := map[string]string{
    "vue":        "your_vue_role_id",
    "typescript": "your_typescript_role_id",
    // ...
}
```

## Project Structure

```
Zdrapywacz/
├── main.go           # Main application and scraper logic
├── discordBot/       # Discord bot implementation
├── databaseconf/     # Database models and configuration
├── .env             # Environment configuration
└── Makefile         # Build and run commands
```

## Environment Variables

- `DB_USER`: MySQL username
- `DB_PASSWORD`: MySQL password
- `DB_PORT`: MySQL port (default: 3306)
- `DB_NAME`: Database name
- `BOT_ID`: Discord bot token

## License

This project is open-source and available under the MIT License.
