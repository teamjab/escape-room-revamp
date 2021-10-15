const sum = require("./apitest.js");

// basic starter test to link up test suite
test("adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(3);
});
