import {describe, expect, it} from "vitest";
import { WC } from "../src/wc.js";


describe("checking custom wc tool", () => {
  const wc = new WC();

  it("should return logged value", () => {
    const log = console.log;
    let loggedOutput = "";
    console.log = (...args) => { loggedOutput += args.join(" ") };
  
    wc.logger("ccwc -c test.txt");
  
    console.log = log;
  
    expect(loggedOutput).toContain("ccwc -l test.txt");
  })
})