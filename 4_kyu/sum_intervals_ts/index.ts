export function sumOfIntervals(intervals: [number, number][]): number {
    if (intervals.length == 0) return 0;
    if (intervals.length == 1) return intervals[0][1] - intervals[0][0];
    let result = 0;
    intervals.sort(([a], [b]) => a - b);
    const [first, ...rest] = intervals;
    let [start, end] = first;
    for (const [currentStart, currentEnd] of rest) {
        if (currentStart > end) {
            result += end - start;
            [start, end] = [currentStart, currentEnd];
        } else {
            end = Math.max(end, currentEnd);
        }
    }
    result += end - start;
    return result;
}