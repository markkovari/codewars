fn exp_sum(n: u64) -> u64 {
    if n == 0 {
        return 1;
    }
    let mut dp = vec![0; (n + 1) as usize];
    dp[0] = 1;
    for i in 1..n {
        for j in i..=n {
            dp[j as usize] += dp[(j - i) as usize];
        }
    }
    dp[n as usize] + 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic_sample_tests() {
        assert_eq!(exp_sum(1), 1);
        assert_eq!(exp_sum(2), 2);
        assert_eq!(exp_sum(3), 3);
        assert_eq!(exp_sum(4), 5);
        assert_eq!(exp_sum(5), 7);
        assert_eq!(exp_sum(10), 42);
    }
}
