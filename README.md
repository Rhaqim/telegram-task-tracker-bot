# Telegram Tracking Bot

A simple Telegram bot written in Go for tracking items or events. This bot allows users to interact with commands to track items, check tracking status, list tracked items, and more. The bot also supports a multi-step interaction to add entries to a tracked item.

## Features

- **Track an item or event:** Users can initiate tracking, provide a title, and add multiple entries.
- **Check tracking status:** Users can check the status of their tracking requests.
- **List tracked items:** Users can view a list of items they have tracked.
- **Handle multiple commands:** The bot provides commands like `/start`, `/track`, `/status`, and `/list`.

## Commands

- `/start`: Start interacting with the bot.
- `/track`: Initiate tracking of an item or event.
- `/status`: Check the current status of your tracking requests.
- `/list`: List all tracked items.
- `/done`: Finish adding entries to the tracked item.

## Getting Started

### Prerequisites

- Go 1.18 or later
- A Telegram bot token from [BotFather](https://core.telegram.org/bots#botfather)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Rhaqim/telegram-task-tracker-bot.git
   cd telegram-task-tracker-bot
   ```

2. **Set up your environment:**

   - Create a `.env` file in the root directory of your project.
   - Add your Telegram bot token:

     ```env
     TELEGRAM_BOT_TOKEN=your-telegram-bot-token
     ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the bot:**

   ```bash
   go run main.go
   ```

## Usage

1. **Add the bot to a Telegram group or chat.**
2. **Start interacting with the bot:**
   - Type `/start` to initialize the bot.
   - Use `/track` to start tracking an item. The bot will guide you through providing a title and multiple entries.
   - Use `/done` to finish tracking.
   - Use `/list` to see your tracked items.

## Code Structure

- **`main.go`:** Entry point of the bot, where the bot is initialized and commands are processed.
- **`repo/commands.go`:** Contains command definitions and handlers for processing different bot commands.
- **`pkg/logger/logger.go`:** Custom logger to handle error logging.

## Contribution

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) - A Golang bindings for the Telegram Bot API.
- [BotFather](https://core.telegram.org/bots#botfather) - Tool for creating Telegram bots.

---

This README provides a comprehensive overview of the bot, including setup instructions, usage, and an outline of the code structure. You can customize the links and information as per your specific repository.
