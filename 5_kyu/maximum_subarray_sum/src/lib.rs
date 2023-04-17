fn max_sequence(seq: &[i32]) -> i32 {
    let mut max = 0;
    let mut sum = 0;
    for i in seq {
        sum += i;
        if sum < 0 {
            sum = 0;
        }
        if sum > max {
            max = sum;
        }
    }
    max
}

#[cfg(test)]
mod tests {
    use super::max_sequence;

    #[test]
    fn sample_tests() {
        assert_eq!(max_sequence(&[-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
        assert_eq!(max_sequence(&[11]), 11);
        assert_eq!(max_sequence(&[-32]), 0);
    }
}
