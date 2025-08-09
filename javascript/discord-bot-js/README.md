# Coding Challenges - Write Your Own Discord Bot

Discord is a platform where everyone in the server can communicate through voice/video calls and text messages.

## Getting started

1. Installation

```bash
git clone git@github.com:victoriacheng15/coding-challenges-lab.git

cd javascript/discord-bot-js
```

2. configaturation

```bash
mv .env.exmaple .env
```

- BOT_TOKEN=""
- CLIENT_ID=""
- CLIENT_SECRET=""

For the values above, please refer to the Discord guide on "Building your first Discord app," which can be accessed [here](https://discord.com/developers/docs/getting-started). The guide provides step-by-step instructions on obtaining the required tokens and IDs for your bot.

- GUILD_ID=""
- CHANNEL_ID=""

For the values above:
- `GUILD_ID` represents the ID of the server where the bot is currently located. You can copy the server ID by right-clicking on the server name.
- `CHANNEL_ID `represents the ID of the specific channel within the server. The process of obtaining the channel ID is the same: right-click on the desired channel and copy its ID.

- API_KEY=""
- API_HOST=""

These values are used to fetch weather data from an API. You can obtain an API key and API host by registering on [Weather by API-NINJA](https://rapidapi.com/apininjas/api/weather-by-api-ninjas/) or any other weather API of your preference.

- CITY=""
- COUNTRY_CODE=""

These settings are for the auto message feature, which sends a weather update to you at a specific time. If you choose not to use this feature, you can ignore these configurations. Additionally, you can safely delete the `autoMessage.js` file located under the `src/utils` folder.

3. Run the bot!

Ensure that Docker and Docker Desktop are installed on your machine by following the guide available [here](https://docs.docker.com/desktop/)

```bash
docker run -d --name=discord-bot .
```

Alternatively, you can use the command `docker compose up` in the terminal, or right-click on the `docker-compose.yml` file and select `compose up`.

## features

### Commands

- /ping
   -  The `/ping` command replies with "pong!" and calculates the time it takes to process the command. This allows you to observe the delay in message editing, demonstrating the bot's ability to send and edit messages after a specified interval.
- /weather
   - The `/weather` command provides weather information for a chosen location. Enter the city name and country code (e.g., ca or us), and the bot will display temperature, minimum temperature, maximum temperature, humidity, sunrise time, and sunset time.

### Messages

- hello
  - The `hello` message is a user-triggered command that prompts the bot to reply directly to the user.
- hey
  - The `hey` message is designed for users who want the bot to send their message to the entire channel, creating a broadcast effect.

### Auto Message

The auto message utilize `node-schedule` that allows you to schedule the time with cron to send weather message at a specific time of your choice.

## Links

- [Write Your Own Discord Bot](https://codingchallenges.fyi/challenges/challenge-discord)
- [Coding Challenges Website](https://codingchallenges.fyi)