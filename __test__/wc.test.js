import { beforeEach, describe, expect, it, vi } from "vitest";
import { WC } from "../src/wc.js";
import { mockedContents } from "./example.js";
import fs from "fs";

describe("WC class", () => {
	let wc;

	beforeEach(() => {
		wc = new WC();
	});

	describe("logger method", () => {
		it("should log the passed message", () => {
			const logSpy = vi.spyOn(console, "log");

			wc.logger("ccwc -l example.txt");

			expect(logSpy).toHaveBeenCalledWith("ccwc -l example.txt");
		});
	});

	describe("fileReader method", () => {
		it("should return the file contents", () => {
			vi.spyOn(fs, "readFileSync").mockReturnValue(mockedContents);

			const result = wc.fileReader("file.txt");

			expect(result).toBe(mockedContents);
		});
	});

	describe("countBytes method", () => {
		it("should return the file size in bytes", () => {
			vi.spyOn(fs, "statSync").mockReturnValue({ size: 100 });

			const result = wc.countBytes("file.txt");

			expect(result).toBe(100);
		});
	});

	describe("countLines method", () => {
		it("returns correct line count for a string", () => {
			vi.spyOn(wc, "fileReader").mockReturnValue(mockedContents);

			const result = wc.countLines("");

			expect(result).toBe(6);
		});
	});

	describe("countWords method", () => {
		it("returns correct word count for a string", () => {
			vi.spyOn(wc, "fileReader").mockReturnValue(mockedContents);

			const result = wc.countWords("");

			expect(result).toBe(53);
		});
	});

	describe("countCharacters method", () => {
		it("returns correct character count for a string", () => {
			vi.spyOn(wc, "fileReader").mockReturnValue(mockedContents);

			const result = wc.countCharacters("");

			expect(result).toBe(346);
		});
	});
});
