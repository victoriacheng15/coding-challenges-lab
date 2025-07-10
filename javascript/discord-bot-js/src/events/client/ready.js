const { Events, ActivityType } = require("discord.js");

module.exports = {
	name: Events.ClientReady,
	once: true,
	async execute(client) {
		console.log(`The bot ${client.user.tag} has logged in and online`);

		const activities = [
			"the weather api",
			"the weather somewhere on the internet",
			"weather forecasting",
		];

		setInterval(() => {
			const status = activities[Math.floor(Math.random() * activities.length)];
			client.user.setPresence({
				activities: [
					{
						name: `${status}`,
						type: ActivityType.Watching,
					},
				],
				status: "dnd",
			});
		}, 10000);
	},
};
