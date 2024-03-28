fn main() {
    println!("Hello, world!");
}

// generate docs for this method
/// # Sum of Intervals
/// Write a function called `sum_intervals` that accepts an array of intervals, and returns the sum of all the interval lengths. Overlapping intervals should only be counted once.
/// Intervals
/// Intervals are represented by a pair of integers in the form of an array. The first value of the interval will always be less than the second value. Interval example: [1, 5] is an interval from 1 to 5. The length of this interval is 4.
/// Overlapping Intervals
/// List containing overlapping intervals:
/// [
/// \[1,4\],
/// \[7, 10\],
/// \[3, 5\]
/// ]
/// The sum of the lengths of these intervals is 7. Since [1, 4] and [3, 5] overlap, we can treat the interval as [1, 5], which has a length of 4.
/// Examples
/// ```
/// use sum_intervals::sum_intervals;
/// assert_eq!(sum_intervals(&[(1, 5)]), 4);
///     
/// assert_eq!(sum_intervals(&[(1, 5), (6, 10)]), 8);
///     
/// assert_eq!(sum_intervals(&[(1, 4), (7, 10), (3, 5)]), 7);
///     
/// assert_eq!(sum_intervals(&[(-1_000_000_000, 1_000_000_000)]), 2_000_000_000);
///     
/// assert_eq!(sum_intervals(&[(0, 20), (-100_000_000, 10), (30, 40)]), 100_000_030);
/// ```
/// # Remarks
/// The input array may contain multiple intervals that overlap with each other.
/// The array may contain duplicate intervals.
/// Do not modify the input.
/// # Performance
/// The test cases have large intervals, so the solution should be efficient.
/// # Source
/// <https://www.codewars.com/kata/52b7ed099cdc285c300001cd/train/rust>

#[allow(dead_code)]
fn sum_intervals(intervals: &[(i32, i32)]) -> i32 {
    if intervals.is_empty() {
        return 0;
    }
    if intervals.len() == 1 {
        return intervals[0].1 - intervals[0].0;
    }
    let mut intervals = intervals.to_vec();

    intervals.sort_by(|a, b| a.0.cmp(&b.0));

    let mut sum = 0;
    let mut longest_range = intervals.iter().take(1).last().unwrap();
    for curr in intervals.iter() {
        sum += if curr.0 > longest_range.1 {
            longest_range.1 - longest_range.0
        } else if curr.1 > longest_range.1 {
            curr.0 - longest_range.0
        } else {
            0
        };
        longest_range = &curr;
    }

    sum += longest_range.1 - longest_range.0;

    return sum;
}

#[cfg(test)]
mod sample_tests {
    use super::*;
    const ERR_MSG: &str = "\nYour result (left) did not match expected output (right).";

    #[test]
    fn non_overlapping_intervals() {
        assert_eq!(sum_intervals(&[(1, 5)]), 4, "{}", ERR_MSG);
        assert_eq!(sum_intervals(&[(1, 5), (6, 10)]), 8, "{}", ERR_MSG);
    }

    #[test]
    fn overlapping_intervals() {
        // assert_eq!(sum_intervals(&[(1, 5), (1, 5)]), 4, "{}", ERR_MSG);
        assert_eq!(sum_intervals(&[(1, 4), (7, 10), (3, 5)]), 7, "{}", ERR_MSG);
    }

    #[test]
    fn large_intervals() {
        assert_eq!(
            sum_intervals(&[(-1_000_000_000, 1_000_000_000)]),
            2_000_000_000,
            "{}",
            ERR_MSG
        );
        assert_eq!(
            sum_intervals(&[(0, 20), (-100_000_000, 10), (30, 40)]),
            100_000_030,
            "{}",
            ERR_MSG
        );
    }
}
