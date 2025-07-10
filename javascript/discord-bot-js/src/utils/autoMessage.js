const schedule = require("node-schedule");
const { getWeatherInfo } = require("./fetchWeather");
const { CHANNEL_ID, CITY, COUNTRY_CODE } = require("./env");

// https://crontab.guru/ - use this link to set the time you would like
const time = "*/1 * * * *";

async function autoMessage(client) {
	const { temp, min_temp, max_temp, humidity } = await getWeatherInfo(
		CITY,
		COUNTRY_CODE,
	);

	console.log(`###### Test Auto Message - ${new Date().toISOString().split("T")[1]} ######`)

	const weatherMsg = `Feeling curious about the weather in ${CITY} today? ️ It's currently ${temp}c, with a low of ${min_temp}c and a high of ${max_temp}c. Humidity's at ${humidity}%, so dress accordingly! `;

	schedule.scheduleJob(time, async () => {
		const channel = client.channels.cache.get(CHANNEL_ID);
		await channel.send(`${weatherMsg}`);

		console.log((`###### The Auto Message has sent at - ${new Date().toISOString().split("T")[1]} ######`))
	});
}

module.exports = { autoMessage };
