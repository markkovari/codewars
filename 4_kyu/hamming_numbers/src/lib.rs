fn hamming(n: usize) -> u64 {
    let mut hamming_numbers = vec![1];
    let mut i2 = 0;

    let mut i3 = 0;
    let mut i5 = 0;

    while hamming_numbers.len() < n {
        let m2 = hamming_numbers[i2] * 2;
        let m3 = hamming_numbers[i3] * 3;
        let m5 = hamming_numbers[i5] * 5;

        let next = m2.min(m3.min(m5));

        if next == m2 {
            i2 += 1;
        }
        if next == m3 {
            i3 += 1;
        }
        if next == m5 {
            i5 += 1;
        }

        hamming_numbers.push(next);
    }

    hamming_numbers[n - 1]
}
// Add your tests here.
// See https://doc.rust-lang.org/stable/rust-by-example/testing/unit_testing.html

#[cfg(test)]
mod tests {
    use super::hamming;

    #[test]
    fn sample_tests() {
        assert_eq!(hamming(1), 1);
        assert_eq!(hamming(2), 2);
        assert_eq!(hamming(3), 3);
        assert_eq!(hamming(4), 4);
        assert_eq!(hamming(5), 5);
        assert_eq!(hamming(6), 6);
        assert_eq!(hamming(7), 8);
        assert_eq!(hamming(8), 9);
        assert_eq!(hamming(9), 10);
        assert_eq!(hamming(10), 12);
        assert_eq!(hamming(11), 15);
        assert_eq!(hamming(12), 16);
        assert_eq!(hamming(13), 18);
        assert_eq!(hamming(14), 20);
        assert_eq!(hamming(15), 24);
        assert_eq!(hamming(16), 25);
        assert_eq!(hamming(17), 27);
        assert_eq!(hamming(18), 30);
        assert_eq!(hamming(19), 32);
    }
}
