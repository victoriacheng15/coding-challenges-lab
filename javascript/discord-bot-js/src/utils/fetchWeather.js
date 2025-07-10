const { API_KEY, API_HOST } = require("./env");

class WeatherUtils {
	async getWeatherInfo(city, country) {
		const API_URL = "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather";
		const URL = `${API_URL}?city=${city}&country=${country}`;
		const options = {
			method: "GET",
			headers: {
				"X-RapidAPI-Key": API_KEY,
				"X-RapidAPI-Host": API_HOST,
			},
		};

		try {
			const response = await fetch(URL, options);
			const data = await response.json();
			return data;
		} catch (error) {
			console.error(error);
		}
	}

	formatTime(date) {
		return new Date(date * 1000).toLocaleString().split(",").at(-1);
	}

	getTodayDate() {
		const today = new Date();
		const year = today.getFullYear();
		const month = (today.getMonth() + 1).toString().padStart(2, "0");
		const day = today.getDate().toString().padStart(2, "0");

		return `${year}-${month}-${day}`;
	}

	formatLocation({ city, countryCode }) {
		const formats = {
			city: city[0].toUpperCase() + city.slice(1),
			countryCode: countryCode.toUpperCase(),
		};

		return `${formats.city}, ${formats.countryCode}`;
	}
}

module.exports = new WeatherUtils();
