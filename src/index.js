#!/usr/bin/env node

import {
	getStats,
	getStatsFromData,
	createOptions,
	checkFlag,
	validateFile,
} from "./utils/index.js";

const args = process.argv.slice(2);
const isInteractive = process.stdin.isTTY;

function processFile(file) {
	const { bytes, lines, words, characters } = getStats(file);
	const options = createOptions(bytes, lines, words, characters, file);

	const hasFlag = args.some((arg) => options.hasOwnProperty(arg));

	if (!hasFlag) {
		console.log(`${lines} ${words} ${bytes} ${file}`);
	}

	checkFlag(args, options);
}

function processInputData(data) {
	const { byteCount, lineCount, wordCount, charCount } = getStatsFromData(data);
	const options = createOptions(byteCount, lineCount, wordCount, charCount);

	checkFlag(args, options);
}

if (isInteractive) {
	const file = validateFile(args);
	processFile(file);
} else {
	let inputData = "";
	process.stdin.setEncoding("utf8");
	process.stdin.on("data", (data) => (inputData += data));
	process.stdin.on("end", () => processInputData(inputData));
}
