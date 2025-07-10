const fs = require("fs");
const path = require("path");

module.exports = (client) => {
	client.handleEvents = async () => {
		const eventPath = path.join(__dirname, "../../", "events");
		const eventFolders = fs.readdirSync(eventPath);

		for (const folder of eventFolders) {
			const eventFiles = fs
				.readdirSync(path.join(eventPath, folder))
				.filter((file) => file.endsWith(".js"));

			switch (folder) {
				case "client":
					for (const file of eventFiles) {
						const event = require(path.join(eventPath, folder, file));
						if (event.once) {
							client.once(event.name, (...args) => event.execute(...args));
						} else {
							client.on(event.name, (...args) => event.execute(...args));
						}
					}
			}
		}
	};
};
