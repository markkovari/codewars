use std::ops::Range;

fn main() {
    println!("Hello, world!");
}
// https://www.codewars.com/kata/52b7ed099cdc285c300001cd/train/rust
fn sum_intervals(intervals: &[(i32, i32)]) -> i32 {
    let mut intervals: Vec<Range<i32>> = intervals.iter().map(|(a, b)| *a..*b).collect();

    intervals.sort_by(|a, b| a.start.cmp(&b.start));

    let mut sum = 0;
    let mut longest_range = intervals.iter().take(1).last().unwrap();
    for curr in intervals.iter() {
        if curr.start > longest_range.end {
            sum += longest_range.end - longest_range.start;
            longest_range = curr;
        } else if curr.end > longest_range.end {
            sum += curr.start - longest_range.start;
            longest_range = &curr;
        }
    }
    sum += longest_range.end - longest_range.start;

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
