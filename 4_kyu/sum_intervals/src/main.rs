use std::ops::Range;

fn main() {
    println!("Hello, world!");
}
// https://www.codewars.com/kata/52b7ed099cdc285c300001cd/train/rust

fn sum_intervals(intervals: &[(i32, i32)]) -> i32 {
    let as_range_internals: Vec<Range<i32>> = intervals
        .iter()
        .map(|(a, b)| *a..*b)
        .collect::<Vec<Range<i32>>>();
    let first = 1..10;
    let second = 6..12;

    todo!();
}

fn intersection(first: Range<i32>, second: Range<i32>) -> Option<usize> {
    let (first_max, second_max) = (first.clone().max(), second.clone().max());
    let (first_min, second_min) = (first.clone().min(), second.clone().min());
    if first_max < second_min || first_min > second_max {
        return None;
    }

    Some(1)
}

// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

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
        assert_eq!(sum_intervals(&[(1, 5), (1, 5)]), 4, "{}", ERR_MSG);
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
