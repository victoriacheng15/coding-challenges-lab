const { REST, Routes } = require("discord.js");
const fs = require("fs");
const path = require("path");
const { BOT_TOKEN, CLIENT_ID, GUILD_ID } = require("../../utils/env");

module.exports = (client) => {
	client.handleCommands = async () => {
		const commandPath = path.join(__dirname, "../../", "commands");
		const commandFolders = fs.readdirSync(commandPath);

		for (const folder of commandFolders) {
			const commandFiles = fs
				.readdirSync(path.join(commandPath, folder))
				.filter((file) => file.endsWith(".js"));

			for (const file of commandFiles) {
				const filePath = path.join(commandPath, folder, file);
				const command = require(filePath);
				client.commands.set(command.data.name, command);
				client.commandArray.push(command.data.toJSON());
			}
		}

		const rest = new REST().setToken(BOT_TOKEN);
		try {
			console.log(`Started refreshing commands.`);
			// Routes.applicationCommands(CLIENT_ID, GUILD_ID) -> this will work for specific guild(server)
			// if only CLIENT_ID => work for multi guilds(servers)
			const data = await rest.put(
				Routes.applicationGuildCommands(CLIENT_ID, GUILD_ID),
				{
					body: client.commandArray,
				},
			);
			console.log(`Refreshed ${data.length} commands.`);
		} catch (error) {
			console.error(`Error refreshing commands: ${error}`);
		}
	};
};
