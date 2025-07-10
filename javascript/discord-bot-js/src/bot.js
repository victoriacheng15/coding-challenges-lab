const fs = require("fs");
const path = require("path");
const { Client, GatewayIntentBits, Collection } = require("discord.js");
const { autoMessage } = require("./utils/autoMessage");
const { BOT_TOKEN } = require("./utils/env");

const client = new Client({
	intents: [
		GatewayIntentBits.Guilds,
		GatewayIntentBits.GuildMessages,
		GatewayIntentBits.MessageContent,
	],
});
client.commands = new Collection();
client.commandArray = [];

const functionFolders = fs.readdirSync(path.join(__dirname, "functions"));
for (const folder of functionFolders) {
	const functionFiles = fs
		.readdirSync(path.join(__dirname, "functions", folder))
		.filter((file) => file.endsWith(".js"));
	for (const file of functionFiles) {
		require(path.join(__dirname, "functions", folder, file))(client);
	}
}

autoMessage(client);

client.handleCommands();
client.handleEvents();
client.login(BOT_TOKEN);
