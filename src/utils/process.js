import { WC } from "../wc.js";

const wc = new WC();

export function getStats(file) {
	return {
		bytes: wc.countBytes(file),
		lines: wc.countLines(file),
		words: wc.countWords(file),
		characters: wc.countCharacters(file),
	};
}

export function getStatsFromData(data) {
	return {
		byteCount: Buffer.byteLength(data, "utf8"),
		lineCount: (data.match(/\n/g) || []).length,
		wordCount: data.split(/\s+/).filter((word) => !!word).length,
		charCount: data.length,
	};
}

export function createOptions(bytes, lines, words, chars, file) {
	const options = {
		"-c": file ? `${bytes} ${file}` : `${bytes}`,
		"-l": file ? `${lines} ${file}` : `${lines}`,
		"-w": file ? `${words} ${file}` : `${words}`,
		"-m": file ? `${chars} ${file}` : `${chars}`,
	};

	return options;
}

export function validateFile(args) {
	const file = args.at(-1);
	if (!file) {
		console.error(
			"Error: No file provided. Please provide a file path or piped input.",
		);
		process.exit(1);
	}
	return file;
}
