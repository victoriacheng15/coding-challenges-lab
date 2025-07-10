const { SlashCommandBuilder } = require("discord.js");

module.exports = {
	data: new SlashCommandBuilder()
		.setName("ping")
		.setDescription("Replies with Pong! and this is a test command"),
	async execute(interaction) {
		// deferReply will make the bot think for a bit before reply
		await interaction.reply("Pong! Calculating ping...");
		const ping = Date.now() - interaction.createdTimestamp;
		try {
			await interaction.editReply(`Pong! Ping time is ${ping}ms.`);
		} catch (error) {
			await interaction.followUp(`Pong! Ping time is ${ping}ms.`);
		}
	},
};
