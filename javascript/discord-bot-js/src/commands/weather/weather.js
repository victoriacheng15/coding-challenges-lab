const { SlashCommandBuilder, EmbedBuilder } = require("discord.js");
const {
	getWeatherInfo,
	getTodayDate,
	formatTime,
	formatLocation,
} = require("../../utils/fetchWeather");

module.exports = {
	data: new SlashCommandBuilder()
		.setName("weather")
		.setDescription("Get the weather of a city")
		.addStringOption((option) =>
			option
				.setName("city")
				.setDescription("Enter the name of the city")
				.setRequired(true),
		)
		.addStringOption((option) =>
			option
				.setName("country-code")
				.setDescription("Enter the country code of the city")
				.setRequired(true),
		),

	async execute(interaction) {
		const city = interaction.options.getString("city");
		const countryCode = interaction.options.getString("country-code");

		const { temp, min_temp, max_temp, humidity, sunrise, sunset } =
			await getWeatherInfo(city, countryCode);
		const today = getTodayDate();
		const location = formatLocation({ city, countryCode });
		const weatherDesc = `Display weather information for ${location} on ${today}`;

		const weatherEmbed = new EmbedBuilder()
			.setColor(0x0099ff)
			.setTitle("Weather Info")
			.setDescription(weatherDesc)
			.addFields(
				{
					name: "Temperature",
					value: `${temp}c`,
					inline: false,
				},
				{
					name: "Min Temperature",
					value: `${min_temp}c`,
					inline: true,
				},
				{
					name: "Max Temperture",
					value: `${max_temp}c`,
					inline: true,
				},
				{
					name: "Humidity",
					value: `${humidity}%`,
					inline: false,
				},
				{
					name: "Sunrise",
					value: `${formatTime(sunrise)}`,
					inline: true,
				},
				{
					name: "Sunset",
					value: `${formatTime(sunset)}`,
					inline: false,
				},
			);

		await interaction.reply({ embeds: [weatherEmbed], ephemeral: false });
	},
};
