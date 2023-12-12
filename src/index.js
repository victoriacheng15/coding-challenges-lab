#!/usr/bin/env node

import { WC } from "./wc.js";

const wc = new WC();
const args = process.argv.slice(2);
let inputData = "";

function checkFlag(obj) {
	const selectedFlag = args.find(flag => obj.hasOwnProperty(flag));
	if (selectedFlag) {
		wc.logger(obj[selectedFlag]);
	}
}

function processFile(file) {
	if (!args.length) {
		wc.logger("Error: No arguments provided. Please provide arguments.");
		return;
	}

	const [bytes, lines, words, chars] = [
		wc.countBytes(file),
		wc.countLines(file),
		wc.countWords(file),
		wc.countCharacters(file),
	];

	const options = {
		"-c": `${bytes} ${file}`,
		"-l": `${lines} ${file}`,
		"-w": `${words} ${file}`,
		"-m": `${chars} ${file}`,
	};

	const hasFlag = args.some((arg => options.hasOwnProperty(arg)));

	if (!hasFlag) {
		wc.logger(`${lines} ${words} ${bytes} ${file}`);
		return;
	}

	checkFlag(options);
}

function processInputData(data) {
  const byteCount = Buffer.byteLength(data, "utf8");
  const lineCount = (data.match(/\n/g) || []).length;
  const wordCount = data.split(/\s+/).filter((word) => word !== "").length;
  const charCount = data.length;

  const options = {
    "-c": `${byteCount}`,
    "-l": `${lineCount}`,
    "-w": `${wordCount}`,
    "-m": `${charCount}`,
  };

  checkFlag(options);
}

if (process.stdin.isTTY) {
  // No piped input, assume file path as last argument
  const file = args.at(-1);
  if (!file) {
    wc.logger("Error: No file provided. Please provide a file path or piped input.");
    process.exit(1);
  }
  processFile(file);
} else {
  // Piped input detected, accumulate data
  process.stdin.setEncoding("utf8");
  process.stdin.on("data", (data) => {
    inputData += data;
  });

  process.stdin.on("end", () => {
    // Combine file and piped data for processing
    if (inputData === "") {
      processFile(args.at(-1));
    } else {
      processInputData(inputData);
    }
  });
}
