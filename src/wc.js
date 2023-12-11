import fs from "fs";

export class WC {
	constructor() {}

	logger(...args) {
		console.log(...args);
	}

	fileReader(file) {
		return fs.readFileSync(file, "utf8");
	}

	countBytes(file) {
		const stats = fs.statSync(file);
		return stats.size;
	}

	countLines(file) {
		const contents = this.fileReader(file);
		const lines = contents.split("\n");
		let lineCounts = lines.length;
		if (lines[lines.length - 1] === "") {
			lineCounts -= 1;
		}

		return lineCounts;
	}

	countWords(file) {
		const contents = this.fileReader(file);
		const words = contents.split(/\s+/).filter((word) => word !== "");
		return words.length;
	}

	countCharacters(file) {
		const contents = this.fileReader(file);
		return contents.length;
	}
}
