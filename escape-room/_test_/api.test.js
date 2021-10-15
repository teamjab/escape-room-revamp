const sum = require("./apitest.js");
const getRequest = require("./apitest.js");

const url = "https://catfact.ninja/breeds";
// basic starter test to link up test suite
describe("api tests", () => {
  test("adds 1 + 2 to equal 3", () => {
    expect(sum(1, 2).toBe(3));
  });
  test("checks if returns API response", () => {
    expect(getRequest(url)).toBe(undefined);
  });
});
