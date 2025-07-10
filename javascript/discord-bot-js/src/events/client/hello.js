const { Events } = require("discord.js");

module.exports = {
	name: Events.MessageCreate,
	async execute(message) {
		if (message.content === "hello") {
			message.reply("Hey there! with reply");
		}

		if (message.content === "hey") {
			await message.channel.send("hey there! without reply");
		}
	},
};
