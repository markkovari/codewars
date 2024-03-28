import { assertEquals } from "https://deno.land/std@0.221.0/assert/mod.ts";
import { sumOfIntervals } from './index.ts'

const testCases: { input: [number, number][], expected: number }[] = [
    { input: [[1, 5]], expected: 4 },
    { input: [[1, 5], [6, 10]], expected: 8 },
    { input: [[1, 5], [1, 5]], expected: 4 },
    { input: [[1, 4], [7, 10], [3, 5]], expected: 7 },
    { input: [[-1e9, 1e9]], expected: 2e9 },
    { input: [[0, 20], [-1e8, 10], [30, 40]], expected: 1e8 + 30 },
];

Deno.test("Test intervals", async (test) => {
    for (const testCase of testCases) {
        await test.step(`Sum of intervals: ${testCase.input} => ${testCase.expected}`, () => {
            assertEquals(sumOfIntervals(testCase.input), testCase.expected, `Failed for ${testCase.input}, should be ${testCase.expected}`);
        });
    }
});
